package grapher

import (
	"fmt"
	"strings"

	"github.com/spilliams/tunnelvision/src/internal"
)

func NewGrapher() internal.Grapher {
	return &grapher{
		loaders: map[string]internal.Loader{},
	}
}

type grapher struct {
	loaders map[string]internal.Loader
}

func (gg *grapher) RegisterLoader(extension string, l internal.Loader) {
	gg.loaders[extension] = l
}

func (gg *grapher) LoadGraphFromFile(filename string) (internal.Graph, error) {
	parts := strings.Split(filename, ".")
	if len(parts) == 1 {
		return nil, fmt.Errorf("filename must have an extension")
	}
	ext := parts[len(parts)-1]
	loader, ok := gg.loaders[ext]
	if !ok {
		return nil, fmt.Errorf("grapher does not recognize file extension %s. Call RegisterLoader?", ext)
	}
	if loader == nil {
		return nil, fmt.Errorf("grapher had a nil Loader mapped to extension %s", ext)
	}
	return loader.LoadGraphFromFile(filename)
}
