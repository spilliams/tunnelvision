package tunnelvision

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newGraphModuleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "graph [DIRECTORY]",
		Short: "graphs the root module at the given directory (`.` by default)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			rootDir, err := os.Getwd()
			if err != nil {
				return err
			}
			if len(args) > 0 {
				rootDir = args[0]
				// clean path to absolute
			}
			outFilename := "output.dot"

			logrus.Infof("reading configuration at %s", rootDir)
			logrus.Infof("outputting graph in file %s", outFilename)
			// err := tfgraph.New(args[0], logrus.StandardLogger(), outFilename)
			// if err != nil {
			// 	return err
			// }
			// logrus.Infof("Wrote graph to %s", outFilename)
			return nil
		},
	}
}
