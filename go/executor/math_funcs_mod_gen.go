// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package executor

func modintAndint(a int, b int) int {
	return a % int(b)
}

var _ = addMathFunc("mod", modintAndint)

func modintAndint32(a int, b int32) int {
	return a % int(b)
}

var _ = addMathFunc("mod", modintAndint32)

func modint32Andint(a int32, b int) int32 {
	return a % int32(b)
}

var _ = addMathFunc("mod", modint32Andint)

func modint32Andint32(a int32, b int32) int32 {
	return a % int32(b)
}

var _ = addMathFunc("mod", modint32Andint32)
