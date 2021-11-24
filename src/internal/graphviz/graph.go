package graphviz

import (
	"github.com/awalterschulze/gographviz"
	"github.com/spilliams/tunnelvision/src/pkg"
)

type graph struct {
	f9l *gographviz.Graph
}

func (g *graph) String() string {
	return g.f9l.String()
}

func (g *graph) Nodes() []pkg.Node {
	from := g.f9l.Nodes.Nodes
	nodes := make([]pkg.Node, len(from))
	for i := 0; i < len(from); i++ {
		nodes[i] = &node{from[i]}
	}
	return nodes
}
