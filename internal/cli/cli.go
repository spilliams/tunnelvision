package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnelvision",
		Short: "A way to visualize your terraform plans",
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	log.Println("hello, world")
		// 	return nil
		// },
	}

	cmd.AddCommand(newSimpleParseCommand())

	return cmd
}

func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func newSimpleParseCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "parse FILE",
		Short: "Parse a single hcl file",
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
