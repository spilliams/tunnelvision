package pkg

type Grapher interface {
	RegisterLoader(extension string, l Loader)
	RegisterWriter(extension string, w GraphWriter)
	LoadGraphFromFile(filename string) error
	WriteGraphToFile(filename string) error
}

type Loader interface {
	LoadGraphFromFile(filename string) (Graph, error)
}

type GraphWriter interface {
	Write(g Graph, filename string) error
}

type Graph interface {
	String() string
}
