package influunt

// Operation describes a single operation in the graph
type Operation struct {
	OpIndex int
	Cache   bool
	Spec    *OpSpec
}

// Output returns an output node for a given index
func (o *Operation) Output(i int) Node {
	return Node{
		ID: [2]int{o.OpIndex, i},
	}
}
