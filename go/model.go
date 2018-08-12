package influunt

import (
	"encoding/gob"
	"io"
)

// Model contains a Model and a usage signature
type Model struct {
	Graph   *Graph
	Inputs  map[string]Node
	Outputs map[string]Node
}

// ReadModel reads Model for io.Reader
func ReadModel(r io.Reader) (*Model, error) {
	model := &Model{}
	dec := gob.NewDecoder(r)
	return model, dec.Decode(model)
}

// WriteModel writes Model to io.Writer
func WriteModel(m *Model, w io.Writer) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(m)
}
