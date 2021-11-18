package cli

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnelvision",
		Short: "visualize and maintain your terraform",
	}

	cmd.AddCommand(newBackendCommand())
	cmd.AddCommand(newFileCommand())
	cmd.AddCommand(newModuleCommand())
	cmd.AddCommand(newRootCommand())
	// cmd.AddCommand(newStackCommand())

	return cmd
}

// Execute runs the cli in this package
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
