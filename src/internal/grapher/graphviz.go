package grapher

import (
	"io/ioutil"

	"github.com/awalterschulze/gographviz"
	"github.com/spilliams/tunnelvision/src/internal"
)

type graphvizLoader struct{}

// NewGraphvizLoader returns a new file loader that knows how to read graphviz
// files. It makes an assumption that the first node in the file is the root of
// the graph
func NewGraphvizLoader() internal.Loader {
	return &graphvizLoader{}
}

func (gvl *graphvizLoader) LoadGraphFromFile(filename string) (internal.Graph, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	graphAst, _ := gographviz.ParseString(string(contents))
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		return nil, err
	}
	return graph, nil
}
