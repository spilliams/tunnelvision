package cli

import (
	"github.com/spf13/cobra"
	"github.com/spilliams/tunnelvision/src/internal/hcl"
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
			return hcl.ParseHCLFile(args[0])
		},
	}
}
