package cli

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnelvision",
		Short: "A way to visualize your terraform plans",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("hello, world")
			return nil
		},
	}

	// cmd.AddCommand(cmds ...*cobra.Command)

	return cmd
}

func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
