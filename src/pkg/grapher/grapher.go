package grapher

import (
	"fmt"
	"strings"

	"github.com/spilliams/tunnelvision/src/pkg"
)

func NewGrapher() pkg.Grapher {
	return &grapher{
		loaders: map[string]pkg.Loader{},
		writers: map[string]pkg.GraphWriter{},
	}
}

type grapher struct {
	loaders map[string]pkg.Loader
	writers map[string]pkg.GraphWriter
	graph   pkg.Graph
}

func (gg *grapher) RegisterLoader(extension string, l pkg.Loader) {
	gg.loaders[extension] = l
}

func (gg *grapher) RegisterWriter(extension string, w pkg.GraphWriter) {
	gg.writers[extension] = w
}

func (gg *grapher) LoadGraphFromFile(filename string) error {
	ext, err := extension(filename)
	if err != nil {
		return err
	}
	loader, ok := gg.loaders[ext]
	if !ok {
		return fmt.Errorf("grapher does not recognize file extension %s", ext)
	}
	if loader == nil {
		return fmt.Errorf("grapher had a nil loader registered to extension %s", ext)
	}
	gg.graph, err = loader.LoadGraphFromFile(filename)
	return err
}

func (gg *grapher) WriteGraphToFile(filename string) error {
	ext, err := extension(filename)
	if err != nil {
		return err
	}
	writer, ok := gg.writers[ext]
	if !ok {
		return fmt.Errorf("grapher does not recognize file extension %s", ext)
	}
	if writer == nil {
		return fmt.Errorf("grapher had a nil writer registered to extension %s", ext)
	}
	return writer.Write(gg.graph, filename)
}

func extension(filename string) (string, error) {
	parts := strings.Split(filename, ".")
	if len(parts) == 1 {
		return "", fmt.Errorf("filename must have an extension")
	}
	return parts[len(parts)-1], nil
}
