package cli

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbose bool

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

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "make log output more verbose")

	return cmd
}

func init() {
	cobra.OnInitialize(initLogger)
}

func initLogger() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{})
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

// Execute runs the cli in this package
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
