package hcl

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/sirupsen/logrus"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// configFileSchema is the schema for the top-level of a config file. We use
// the low-level HCL API for this level so we can easily deal with each
// block type separately with its own decoding logic.
var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "terraform"},
		{Type: "provider", LabelNames: []string{"name"}},
		{Type: "variable", LabelNames: []string{"name"}},
		{Type: "locals"},
		{Type: "output", LabelNames: []string{"name"}},
		{Type: "module", LabelNames: []string{"name"}},
		{Type: "resource", LabelNames: []string{"id", "name"}},
		{Type: "data", LabelNames: []string{"type", "name"}},
	},
}

type configuration struct {
	locals        []*Local
	outputs       []*Output           `hcl:"output,block"`
	providerLocks []*LockfileProvider `hcl:"provider,block"`
	resources     []*Resource         `hcl:"resource,block"`
	variables     []*Variable         `hcl:"variable,block"`
	// Terraform
	// Modules
	// Datas
}

type Variable struct {
	ID          string      `hcl:"id,label"`
	Type        interface{} `hcl:"type,attr"`
	Description string      `hcl:"description,optional"`
	Default     interface{} `hcl:"default,optional"`
	Remain      hcl.Body    `hcl:",remain"`
}

type Resource struct {
	Type   string   `hcl:"type,label"`
	Name   string   `hcl:"name,label"`
	Remain hcl.Body `hcl:",remain"`
}

type Output struct {
	ID          string `hcl:"id,label"`
	Description string `hcl:"description,optional"`
	Value       string `hcl:"value"`
}

type Local struct {
	Name      string
	attribute *hcl.Attribute
}

func newConfiguration() *configuration {
	return &configuration{
		locals:        make([]*Local, 0),
		outputs:       make([]*Output, 0),
		providerLocks: make([]*LockfileProvider, 0),
		resources:     make([]*Resource, 0),
		variables:     make([]*Variable, 0),
	}
}

// terraformBlockSchema is the schema for a top-level "terraform" block in
// a configuration file.
var terraformBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{Name: "required_version"},
		{Name: "experiments"},
		{Name: "language"},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "backend", LabelNames: []string{"type"}},
		{Type: "cloud"},
		{Type: "required_providers"},
		{Type: "provider_meta", LabelNames: []string{"provider"}},
	},
}

// ModuleParser is something that can parse a terraform module (configuration
// directory)
type ModuleParser interface {
	ParseModuleDirectory(string) error
	ParseTerraformFile(string) error
	Parser() *hclparse.Parser
	Module() *tfconfig.Module
}

type moduleParser struct {
	cfg         *configuration
	fundamental *hclparse.Parser
	module      *tfconfig.Module
	*logrus.Logger
}

// NewModuleParser builds an object that conforms to the ModuleParser interface
func NewModuleParser(logger *logrus.Logger) ModuleParser {
	return &moduleParser{
		cfg:         newConfiguration(),
		fundamental: hclparse.NewParser(),
		Logger:      logger,
	}
}

func (mp *moduleParser) ParseModuleDirectory(dirname string) error {
	module, diags := tfconfig.LoadModule(dirname)
	if diags.HasErrors() {
		return fmt.Errorf(diags.Error())
	}
	mp.module = module

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	for _, f := range files {
		logrus.Debugf("examining file %s", f.Name())
		if strings.HasSuffix(f.Name(), ".tf") || strings.HasSuffix(f.Name(), ".hcl") {
			fullname := path.Join(dirname, f.Name())
			logrus.Debug("parsing!")
			if err := mp.ParseTerraformFile(fullname); err != nil {
				return err
			}
		}
	}
	return nil
}

// ParseTerraformFile parses a single file
func (mp *moduleParser) ParseTerraformFile(filename string) error {
	file, diags := mp.fundamental.ParseHCLFile(filename)
	if err := handleDiags(mp.fundamental, diags); err != nil {
		return err
	}

	// the _ drops the part of the body that does not match the schema
	content, _, diags := file.Body.PartialContent(configFileSchema)
	if err := handleDiags(mp.fundamental, diags); err != nil {
		return err
	}

	// TODO: maybe we don't need to fully decode all the blocks. Maybe it would
	// be enough just to loop through the content.Blocks and register their Type,
	// Labels, and Body.Attributes[].Expr for scope traversals...
	// Decoding them lets me learn more, but if they're all type hcl.Body, I
	// don't see why I have to decode them into multiple struct types too.
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{},
		Functions: map[string]function.Function{},
	}
	for _, block := range content.Blocks {
		var diags hcl.Diagnostics
		switch block.Type {
		case "provider":
			provider := &LockfileProvider{}
			diags = gohcl.DecodeBody(block.Body, ctx, provider)
			if err := handleDiags(mp.fundamental, diags); err != nil {
				break
			}
			provider.ID = block.Labels[0]
			mp.cfg.providerLocks = append(mp.cfg.providerLocks, provider)
		case "variable":
			variable := &Variable{}
			diags = gohcl.DecodeBody(block.Body, ctx, variable)
			if err := handleDiags(mp.fundamental, diags); err != nil {
				break
			}
			variable.ID = block.Labels[0]
			mp.cfg.variables = append(mp.cfg.variables, variable)
		case "resource":
			resource := &Resource{}
			diags = gohcl.DecodeBody(block.Body, ctx, resource)
			if err := handleDiags(mp.fundamental, diags); err != nil {
				break
			}
			resource.Type = block.Labels[0]
			resource.Name = block.Labels[1]
			mp.cfg.resources = append(mp.cfg.resources, resource)
		case "output":
			output := &Output{}
			diags = gohcl.DecodeBody(block.Body, ctx, output)
			if err := handleDiags(mp.fundamental, diags); err != nil {
				break
			}
			output.ID = block.Labels[0]
			mp.cfg.outputs = append(mp.cfg.outputs, output)
		case "locals":
			var locals hcl.Attributes
			locals, diags = block.Body.JustAttributes()
			if err := handleDiags(mp.fundamental, diags); err != nil {
				break
			}
			for name, attr := range locals {
				local := &Local{
					Name:      name,
					attribute: attr,
				}
				mp.cfg.locals = append(mp.cfg.locals, local)
			}
		default:
			mp.Warningf("unrecognized block type %s: %v", block.Type, block.Body)
		}
		if err := handleDiags(mp.fundamental, diags); err != nil {
			return err
		}
	}

	return nil
}

func (mp *moduleParser) Parser() *hclparse.Parser {
	return mp.fundamental
}

func (mp *moduleParser) Module() *tfconfig.Module {
	return mp.module
}
