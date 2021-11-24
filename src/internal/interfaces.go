package internal

type Grapher interface {
	RegisterLoader(extension string, l Loader)
	LoadGraphFromFile(filename string) (Graph, error)
}

type Loader interface {
	LoadGraphFromFile(filename string) (Graph, error)
}

type Graph interface {
}
