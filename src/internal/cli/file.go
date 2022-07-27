package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/tunnelvision/src/internal/hcl"
	"github.com/spilliams/tunnelvision/src/pkg/tfgraph"
)

func newParseFileCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "parse-file FILE",
		Short: "parse a single hcl file [WIP]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return hcl.ParseHCLFile(args[0])
		},
	}
}

func newGraphFileCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "graph-file FILE",
		Short: "perform operations on a single .dot or .gv file [WIP]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			outFilename := "output.dot"
			err := tfgraph.New(args[0], logrus.StandardLogger(), outFilename)
			if err != nil {
				return err
			}
			logrus.Infof("Wrote graph to %s", outFilename)
			return nil
		},
	}
}
