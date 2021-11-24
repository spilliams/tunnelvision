package grapher

import "github.com/spilliams/tunnelvision/src/internal"

func LoadFromFile(filename string) (internal.Grapher, error) {
	return &grapher{}, nil
}

type grapher struct {
}

func (g *grapher) String() string {
	return "grapher"
}
