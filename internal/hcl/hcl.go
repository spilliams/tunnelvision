package hcl

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	log "github.com/sirupsen/logrus"
)

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

func handleDiags(parser *hclparse.Parser, diags hcl.Diagnostics) error {
	if diags.HasErrors() {
		wr := hcl.NewDiagnosticTextWriter(
			os.Stderr,
			parser.Files(),
			100,  // wrapping width
			true, // colors
		)
		wr.WriteDiagnostics(diags)
		return fmt.Errorf("errors found")
	}
	return nil
}
