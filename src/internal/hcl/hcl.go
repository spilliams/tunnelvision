package hcl

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	log "github.com/sirupsen/logrus"
)

func ParseHCLFile(filename string) error {
	p := hclparse.NewParser()
	file, diags := p.ParseHCLFile(filename)
	if diags.HasErrors() {
		return fmt.Errorf("%#v", diags)
	}
	log.Debugf("%#v\n", file.Body)

	log.Debugf("%s\n", file.Bytes)

	// attributes, diags := file.Body.JustAttributes()
	// if diags.HasErrors() {
	// 	return fmt.Errorf("%#v", diags)
	// }
	// log.Debugf("%#v\n", attributes)
	return nil
}

var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "terraform",
		},
		{
			// This one is not really valid, but we include it here so we
			// can create a specialized error message hinting the user to
			// nest it inside a "terraform" block.
			Type: "required_providers",
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type: "locals",
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type: "moved",
		},
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
		{
			Type:       "backend",
			LabelNames: []string{"type"},
		},
		{
			Type: "cloud",
		},
		{
			Type: "required_providers",
		},
		{
			Type:       "provider_meta",
			LabelNames: []string{"provider"},
		},
	},
}
