package graphviz

import (
	"io/ioutil"
	"os"

	"github.com/awalterschulze/gographviz"
	"github.com/spilliams/tunnelvision/src/pkg"
)

type reader struct{}

// NewReader returns a new file-reader that knows how to read graphviz
// files. It makes an assumption that the first node in the file is the root of
// the graph
func NewReader() pkg.GraphReader {
	return &reader{}
}

func (r *reader) Read(filename string) (pkg.Graph, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	graphAst, _ := gographviz.ParseString(string(contents))
	g := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, g); err != nil {
		return nil, err
	}
	return &graph{g}, nil
}

type writer struct{}

func NewWriter() pkg.GraphWriter {
	return &writer{}
}

func (w *writer) Write(g pkg.Graph, filename string) error {
	return os.WriteFile(filename, []byte(g.String()), 0777)
}
