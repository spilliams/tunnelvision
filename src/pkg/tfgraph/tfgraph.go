package tfgraph

import (
	"github.com/spilliams/tunnelvision/src/pkg/grapher"
)

func TfGraph(inFile string, outFile string) error {
	gg := grapher.NewGrapher()
	gvLoader := grapher.NewGraphvizLoader()
	gg.RegisterLoader("dot", gvLoader)
	gg.RegisterLoader("gv", gvLoader)

	if err := gg.LoadGraphFromFile(inFile); err != nil {
		return err
	}

	gvWriter := grapher.NewGraphvizWriter()
	gg.RegisterWriter("dot", gvWriter)
	gg.RegisterWriter("gv", gvWriter)
	return gg.WriteGraphToFile(outFile)
}
