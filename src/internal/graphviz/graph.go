package graphviz

import (
	"fmt"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/sirupsen/logrus"
	"github.com/spilliams/tunnelvision/src/pkg"
)

type graph struct {
	*logrus.Logger
	f9l *gographviz.Graph
}

func (g *graph) String() string {
	return g.f9l.String()
}

func (g *graph) SetLogger(l *logrus.Logger) {
	g.Logger = l
}

func (g *graph) Nodes() []pkg.Node {
	from := g.f9l.Nodes.Nodes
	nodes := make([]pkg.Node, len(from))
	for i := 0; i < len(from); i++ {
		nodes[i] = &node{from[i]}
	}
	return nodes
}

// WalkNodes iterates over the receiver's nodes, calling the given func on each.
// If the iterator returns nil, the graph will trim that node from itself.
func (g *graph) WalkNodes(f func(pkg.Node) pkg.Node) (int, int) {
	total := len(g.Nodes())
	returned := 0
	for _, n := range g.Nodes() {
		name := n.String()
		n = f(n)
		if n == nil {
			g.Debugf("node %s didn't filter: removing", name)
			if err := g.RemoveNode(name); err != nil {
				g.Warnf("%v", err)
			}
			continue
		}
		returned++
	}
	return total, returned
}

func (g *graph) RemoveNode(name string) error {
	parents := g.f9l.Relations.ChildToParents[name]
	var parent string
	for parentName, ok := range parents {
		if ok {
			parent = parentName
			break
		}
	}
	if err := g.f9l.RemoveNode(parent, name); err != nil {
		return fmt.Errorf("could not remove node %s from parent %s: %v", name, parent, err)
	}
	return g.checkErr()
}

func (g *graph) checkErr() error {
	check := g.f9l.String()
	if strings.HasPrefix(check, "error: ") {
		return fmt.Errorf(check)
	}

	return nil
}
