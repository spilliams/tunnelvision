package hcl

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
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
		{Type: "resource", LabelNames: []string{"type", "name"}},
		{Type: "data", LabelNames: []string{"type", "name"}},
	},
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
}

type moduleParser struct {
	fundamental *hclparse.Parser
	parsed      map[string]*hcl.File
}

// NewModuleParser builds an object that conforms to the ModuleParser interface
func NewModuleParser() ModuleParser {
	return &moduleParser{
		fundamental: hclparse.NewParser(),
		parsed:      make(map[string]*hcl.File, 0),
	}
}

func (mp *moduleParser) ParseModuleDirectory(dirname string) error {
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
	if mp.parsed[filename] != nil {
		return nil
	}
	file, diags := mp.fundamental.ParseHCLFile(filename)
	if err := handleDiags(mp.fundamental, diags); err != nil {
		return err
	}
	mp.parsed[filename] = file

	// log.Debugf("%s\n", file.Bytes)

	// attributes, diags := file.Body.JustAttributes()
	// if diags.HasErrors() {
	// 	return fmt.Errorf("%#v", diags.Error())
	// }
	// log.Debugf("%#v\n", attributes)
	return nil
}

func (mp *moduleParser) Parser() *hclparse.Parser {
	return mp.fundamental
}
