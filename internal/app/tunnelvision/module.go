package tunnelvision

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/tunnelvision/internal/hcl"
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
				rootDir, err = filepath.Abs(rootDir)
				if err != nil {
					return err
				}
			}
			// outFilename := "output.dot"

			logrus.Infof("reading configuration at %s", rootDir)
			// logrus.Infof("outputting graph in file %s", outFilename)

			parser := hcl.NewModuleParser()
			parser.ParseModuleDirectory(rootDir)

			fmt.Printf("%#v\n", parser.Parser())
			// err := tfgraph.New(args[0], logrus.StandardLogger(), outFilename)
			// if err != nil {
			// 	return err
			// }
			// logrus.Infof("Wrote graph to %s", outFilename)
			return nil
		},
	}
}
