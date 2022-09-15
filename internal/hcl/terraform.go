package hcl

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/sirupsen/logrus"
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
	attributes hcl.Attributes
	blocks     hcl.Blocks
}

func newConfiguration() *configuration {
	return &configuration{
		attributes: hcl.Attributes{},
		blocks:     hcl.Blocks{},
	}
}

// ModuleParser is something that can parse a terraform module (configuration
// directory)
type ModuleParser interface {
	Configuration() *configuration
	Module() *tfconfig.Module
	ParseModuleDirectory(string) error
	Parser() *hclparse.Parser
	ParseTerraformFile(string) error
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

	content, _, diags := file.Body.PartialContent(configFileSchema)
	if err := handleDiags(mp.fundamental, diags); err != nil {
		return err
	}

	// read the contents into the receiver
	for _, block := range content.Blocks {
		var diags hcl.Diagnostics
		if block.Type == "locals" {
			var locals hcl.Attributes
			locals, diags = block.Body.JustAttributes()
			if err := handleDiags(mp.fundamental, diags); err != nil {
				return err
			}
			for name, attr := range locals {
				mp.cfg.attributes[name] = attr
			}
		} else {
			mp.cfg.blocks = append(mp.cfg.blocks, block)
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

func (mp *moduleParser) Configuration() *configuration {
	return mp.cfg
}
