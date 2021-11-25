package graphviz

import (
	"fmt"
	"strings"

	"github.com/awalterschulze/gographviz"
	"github.com/spilliams/tunnelvision/src/pkg"
)

type node struct {
	f9l *gographviz.Node
}

func (n *node) String() string {
	return n.f9l.Name
}

func (n *node) Attribute(key pkg.AttributeKey) string {
	val, ok := n.f9l.Attrs[gographviz.Attr(key.String())]
	if !ok {
		return val
	}
	return strings.TrimPrefix(strings.TrimSuffix(val, `"`), `"`)
}

func (n *node) SetAttribute(key pkg.AttributeKey, value string) {
	strings.Join(strings.Split(value, `"`), `\"`)
	n.f9l.Attrs.Add(key.String(), fmt.Sprintf(`"%s"`, value))
}
