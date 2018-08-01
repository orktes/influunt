package executor

func addStringAndString(a, b string) string {
	return a + b
}

func equalStringAndString(a, b string) bool {
	return a == b
}

func notEqualStringAndString(a, b string) bool {
	return a == b
}

var _ = addMathFunc("add", addStringAndString)
var _ = addMathFunc("eq", equalStringAndString)
var _ = addMathFunc("ne", notEqualStringAndString)
