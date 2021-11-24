package pkg

type Grapher interface {
	RegisterReader(extension string, r GraphReader)
	RegisterWriter(extension string, w GraphWriter)
	ReadGraphFromFile(filename string) error
	WriteGraphToFile(filename string) error
}

type GraphReader interface {
	Read(filename string) (Graph, error)
}

type GraphWriter interface {
	Write(g Graph, filename string) error
}

type Graph interface {
	String() string
}
