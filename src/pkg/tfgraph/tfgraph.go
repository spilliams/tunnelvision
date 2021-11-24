package tfgraph

import (
	"fmt"

	"github.com/spilliams/tunnelvision/src/pkg/grapher"
)

func TfGraph(inFile string) error {
	gg := grapher.NewGrapher()
	gvLoader := grapher.NewGraphvizLoader()
	gg.RegisterLoader("dot", gvLoader)
	gg.RegisterLoader("gv", gvLoader)

	g, err := gg.LoadGraphFromFile(inFile)
	if err != nil {
		return err
	}
	fmt.Println(g)

	return nil
}
