package graphviz

import "github.com/awalterschulze/gographviz"

type node struct {
	f9l *gographviz.Node
}

func (n *node) String() string {
	return n.f9l.Name
}

func (n *node) SetName(s string) {
	n.f9l.Name = s
}
