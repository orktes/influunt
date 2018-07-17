package influunt

// Add operations adds two nodes together (a + b)
func Add(g *Graph, a, b Node) Node {
	return g.AddOperation(OpSpec{
		Type:   "Add",
		Inputs: []Node{a, b},
	}).Output(0)
}

// Placeholder returns a new placeholder node
func Placeholder(g *Graph) Node {
	return g.AddOperation(OpSpec{
		Type: "Placeholder",
	}).Output(0)
}

// Map maps over a given list
func Map(g *Graph, list Node, fn func(item Node, index Node) Node) Node {
	start := len(g.Operations)
	argPh := Placeholder(g)
	indexPh := Placeholder(g)
	fnRes := fn(argPh, indexPh)

	for i := start; i < len(g.Operations); i++ {
		g.Operations[i].Cache = false
	}

	return g.AddOperation(OpSpec{
		Type:   "Map",
		Inputs: []Node{list, fnRes},
		Attrs: map[string]interface{}{
			"argument": argPh,
			"index":    indexPh,
		},
	}).Output(0)

}

// ParseJSON parses json
func ParseJSON(g *Graph, json Node) Node {
	return g.AddOperation(OpSpec{
		Type:   "ParseJSON",
		Inputs: []Node{json},
	}).Output(0)
}

// GetAttr returns an attribute from an object or map
func GetAttr(g *Graph, m Node, key Node) Node {
	return g.AddOperation(OpSpec{
		Type:   "GetAttr",
		Inputs: []Node{m, key},
	}).Output(0)
}

// Cond returs a if pred is "true" and b if pred is "false"
func Cond(g *Graph, pred Node, a, b func() Node) Node {
	return g.AddOperation(OpSpec{
		Type:   "Cond",
		Inputs: []Node{pred, a(), b()},
	}).Output(0)
}

// Const returns a constant value node
func Const(g *Graph, val interface{}) Node {
	return g.AddOperation(OpSpec{
		Type: "Const",
		Attrs: map[string]interface{}{
			"value": val,
		},
	}).Output(0)
}
