package influunt

// OpSpec defines specification for an operation to be a added
type OpSpec struct {
	Type   string
	Inputs []Node
	Attrs  map[string]interface{}
}
