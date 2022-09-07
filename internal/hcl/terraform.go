package hcl

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	log "github.com/sirupsen/logrus"
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

// ParseHCLFile parses an HCL file...
func ParseHCLFile(filename string) error {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(filename)
	if err := handleDiags(parser, diags); err != nil {
		return fmt.Errorf("%#v", err)
	}
	log.Debugf("%#v\n", file.Body)

	// log.Debugf("%s\n", file.Bytes)

	// attributes, diags := file.Body.JustAttributes()
	// if diags.HasErrors() {
	// 	return fmt.Errorf("%#v", diags.Error())
	// }
	// log.Debugf("%#v\n", attributes)
	return nil
}
