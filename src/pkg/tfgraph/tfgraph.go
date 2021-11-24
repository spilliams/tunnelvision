package tfgraph

import (
	"github.com/spilliams/tunnelvision/src/pkg/grapher"
)

func TfGraph(inFile string, outFile string) error {
	gg := grapher.NewGrapher()
	gvReader := grapher.NewGraphvizReader()
	gg.RegisterReader("dot", gvReader)
	gg.RegisterReader("gv", gvReader)

	if err := gg.ReadGraphFromFile(inFile); err != nil {
		return err
	}

	gvWriter := grapher.NewGraphvizWriter()
	gg.RegisterWriter("dot", gvWriter)
	gg.RegisterWriter("gv", gvWriter)
	return gg.WriteGraphToFile(outFile)
}
