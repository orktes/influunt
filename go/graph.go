package influunt

import (
	"encoding/gob"
	"fmt"
	"io"
)

// Graph structure for containing the graph
type Graph struct {
	Operations []*Operation
}

// NewGraph returns a new graph
func NewGraph() *Graph {
	return &Graph{}
}

// ReadGraph reads graph for io.Reader
func ReadGraph(r io.Reader) (*Graph, error) {
	graph := NewGraph()
	dec := gob.NewDecoder(r)
	return graph, dec.Decode(graph)
}

// WriteGraph writes graph to io.Writer
func WriteGraph(g *Graph, w io.Writer) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(g)
}

// NodeByName returns node for a given name
func (g *Graph) NodeByName(name string, index int) *Node {
	for _, op := range g.Operations {
		if op.Spec.Name == name {
			n := op.Output(index)
			return &n
		}
	}

	return nil
}

// AddOperation adds a new operation to graph
func (g *Graph) AddOperation(spec OpSpec) *Operation {
	i := len(g.Operations)

	if spec.Name == "" {
		spec.Name = fmt.Sprintf("%s:%d", spec.Type, i)
	}

	op := &Operation{
		OpIndex: i,
		Spec:    spec,
		Cache:   true,
	}
	g.Operations = append(g.Operations, op)
	return op
}

func init() {
	gob.Register(map[string]interface{}{})
	gob.Register(Node{})
}
