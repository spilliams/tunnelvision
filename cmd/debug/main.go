package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spilliams/tunnelvision/internal/hcl"
)

func main() {
	rootDir := "/Users/spencer/spilliams/tunnelvision/fixtures/examples/simple-graph"
	logrus.Infof("reading configuration at %s", rootDir)
	// logrus.Infof("outputting graph in file %s", outFilename)

	parser := hcl.NewModuleParser()
	parser.ParseModuleDirectory(rootDir)

	logrus.Debugf("%#v", parser.Parser())
	logrus.Debugf("%#v", parser.Module())

	logrus.Info("Done")
}
