package hcl

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

func handleDiags(parser *hclparse.Parser, diags hcl.Diagnostics) error {
	if diags == nil {
		return nil
	}
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
