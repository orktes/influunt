package influunt

// Graph structure for containing the graph
type Graph struct {
	Operations []*Operation
}

// NewGraph returns a new graph
func NewGraph() *Graph {
	return &Graph{}
}

// AddOperation adds a new operation to graph
func (g *Graph) AddOperation(spec *OpSpec) *Operation {
	i := len(g.Operations)
	op := &Operation{
		OpIndex: i,
		Spec:    spec,
		Cache:   true,
	}
	g.Operations = append(g.Operations, op)
	return op
}
