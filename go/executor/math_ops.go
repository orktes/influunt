package executor

import (
	"github.com/orktes/influunt/go"
	"github.com/pkg/errors"
)

//go:generate genny -in=math_funcs.tmpl -out=math_funcs_gen.go gen "AType=int,int32,float32,float64 BType=int,int32"
//go:generate genny -in=math_funcs_float.tmpl -out=math_funcs_float_gen.go gen "AType=int,int32,float32,float64 BType=float32,float64"
//go:generate genny -in=math_funcs_mod.tmpl -out=math_funcs_mod_gen.go gen "AType=int,int32 BType=int,int32"

func generateMathOp(typ string) func(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
	mathFunc := getMathFunc(typ)

	return func(c *Context, e *Executor, inputs []influunt.Node, attrs map[string]interface{}) ([]interface{}, error) {
		if len(inputs) != 2 {
			return nil, errors.New("operation takes two arguments")
		}

		a, err := e.executeOp(c, inputs[0])
		if err != nil {
			return nil, err
		}
		b, err := e.executeOp(c, inputs[1])
		if err != nil {
			return nil, err
		}

		val, err := mathFunc(a, b)
		if err != nil {
			return nil, err
		}

		return []interface{}{val}, nil
	}
}

func init() {
	AddOperation("Add", generateMathOp("add"))
	AddOperation("Div", generateMathOp("div"))
	AddOperation("Sub", generateMathOp("substract"))
	AddOperation("Mod", generateMathOp("mod"))
	AddOperation("Mul", generateMathOp("mul"))
	AddOperation("LessThan", generateMathOp("lt"))
	AddOperation("LessOrEqual", generateMathOp("lte"))
	AddOperation("GreaterThan", generateMathOp("gt"))
	AddOperation("GreaterOrEqual", generateMathOp("gte"))
	AddOperation("Equal", generateMathOp("eq"))
	AddOperation("NotEqual", generateMathOp("ne"))
}
