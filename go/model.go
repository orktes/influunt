package influunt

// Model contains a graph and a usage signature
type Model struct {
	Graph   *Graph
	Inputs  map[string]Node
	Outputs map[string]Node
}
