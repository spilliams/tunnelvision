package cli

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/cobra"
)

func newFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "file",
		Aliases: []string{"f"},
		Short:   "commands relating to a single file",
	}

	cmd.AddCommand(newParseFileCommand())

	return cmd
}

func newParseFileCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "parse FILE",
		Short: "parse a single hcl file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			p := hclparse.NewParser()
			file, diagnostics := p.ParseHCLFile(args[0])
			fmt.Println(file)
			fmt.Println(diagnostics)
			return nil
		},
	}
}
