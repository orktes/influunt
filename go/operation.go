package influunt

import "fmt"

// Operation describes a single operation in the graph
type Operation struct {
	OpIndex int
	Cache   bool
	Spec    OpSpec
}

// Output returns an output node for a given index
func (o *Operation) Output(i int) Node {
	return Node{
		Name: fmt.Sprintf("%s:%d", o.Spec.Name, i),
		ID:   [2]int{o.OpIndex, i},
	}
}
