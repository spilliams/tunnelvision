package tfgraph

import (
	"fmt"
	"strings"

	"github.com/spilliams/tunnelvision/src/internal/graphviz"
	"github.com/spilliams/tunnelvision/src/pkg"
	"github.com/spilliams/tunnelvision/src/pkg/grapher"
)

func TfGraph(inFile string, outFile string) error {
	gg := grapher.NewGrapher()
	gvReader := graphviz.NewReader()
	gg.RegisterReader("dot", gvReader)
	gg.RegisterReader("gv", gvReader)

	if err := gg.ReadGraphFromFile(inFile); err != nil {
		return err
	}

	// as of terraform 1.0.11, the graph has a *lot* of extra stuff. We should
	// work on them one at a time, so that later we could put them under feature
	// flags.
	g := gg.Graph()
	g.WalkNodes(func(n pkg.Node) pkg.Node {
		n.SetName(strings.TrimPrefix(strings.TrimSuffix(strings.TrimPrefix(n.String(), "\""), "\""), "[root] "))
		return n
	})

	for _, node := range g.Nodes() {
		fmt.Println(node)
	}

	gvWriter := graphviz.NewWriter()
	gg.RegisterWriter("dot", gvWriter)
	gg.RegisterWriter("gv", gvWriter)
	return gg.WriteGraphToFile(outFile)
}
