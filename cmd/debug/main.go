package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spilliams/tunnelvision/internal/hcl"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	rootDir := "/Users/spencer/spilliams/tunnelvision/fixtures/examples/simple-graph"
	logrus.Infof("reading configuration at %s", rootDir)
	// logrus.Infof("outputting graph in file %s", outFilename)

	parser := hcl.NewModuleParser(logrus.StandardLogger())
	if err := parser.ParseModuleDirectory(rootDir); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Debugf("%#v", parser.Parser())
	logrus.Debugf("%#v", parser.Module())

	logrus.Info("Done")
}
