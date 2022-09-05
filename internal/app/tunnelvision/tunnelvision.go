package tunnelvision

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbose bool

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tunnelvision",
		Short: "visualize and maintain your terraform configurations",
	}

	cmd.AddCommand(newParseFileCommand())
	cmd.AddCommand(newGraphFileCommand())

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "make log output more verbose")

	return cmd
}

func init() {
	cobra.OnInitialize(initLogger)
}

func initLogger() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// Execute runs the cli in this package
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
