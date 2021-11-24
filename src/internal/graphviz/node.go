package graphviz

import "github.com/awalterschulze/gographviz"

type node struct {
	f9l *gographviz.Node
}

func (n *node) String() string {
	return n.f9l.Name
}
