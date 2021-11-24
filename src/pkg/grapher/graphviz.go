package grapher

import (
	"io/ioutil"
	"os"

	"github.com/awalterschulze/gographviz"
	"github.com/spilliams/tunnelvision/src/pkg"
)

type graphvizReader struct{}

// NewGraphvizReader returns a new file-reader that knows how to read graphviz
// files. It makes an assumption that the first node in the file is the root of
// the graph
func NewGraphvizReader() pkg.GraphReader {
	return &graphvizReader{}
}

func (gvl *graphvizReader) Read(filename string) (pkg.Graph, error) {
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

type graphvizWriter struct{}

func NewGraphvizWriter() pkg.GraphWriter {
	return &graphvizWriter{}
}

func (gvw *graphvizWriter) Write(g pkg.Graph, filename string) error {
	return os.WriteFile(filename, []byte(g.String()), 0777)
}
