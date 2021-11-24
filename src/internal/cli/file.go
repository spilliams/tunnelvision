package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/tunnelvision/src/internal/grapher"
	"github.com/spilliams/tunnelvision/src/internal/hcl"
)

func newFileCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "file",
		Aliases: []string{"f"},
		Short:   "commands relating to a single file",
	}

	cmd.AddCommand(newParseFileCommand())
	cmd.AddCommand(newGraphFileCommand())

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

func newGraphFileCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "graph FILE",
		Short: "perform operations on a single .dot or .gv file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			gg := grapher.NewGrapher()
			gvLoader := grapher.NewGraphvizLoader()
			gg.RegisterLoader("dot", gvLoader)
			gg.RegisterLoader("gv", gvLoader)

			g, err := gg.LoadGraphFromFile(args[0])
			if err != nil {
				return err
			}
			fmt.Println(g)
			return nil
		},
	}
}
