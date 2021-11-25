package pkg

import "github.com/sirupsen/logrus"

type Logger interface {
	SetLogger(*logrus.Logger)
}

type Grapher interface {
	Logger
	RegisterReader(extension string, r GraphReader)
	RegisterWriter(extension string, w GraphWriter)
	ReadGraphFromFile(filename string) error
	WriteGraphToFile(filename string) error
	Graph() Graph
}

type GraphReader interface {
	Logger
	Read(filename string) (Graph, error)
}

type GraphWriter interface {
	Logger
	Write(g Graph, filename string) error
}

type Graph interface {
	Logger
	String() string
	Nodes() []Node
	WalkNodes(func(Node) Node)
}

type Node interface {
	String() string
	Attribute(AttributeKey) string
	SetAttribute(AttributeKey, string)
}

type AttributeKey interface {
	String() string
}
