package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestCond(t *testing.T) {
	graph := influunt.NewGraph()

	aNode := influunt.Add(graph, influunt.Const(graph, 1), influunt.Const(graph, 2))
	bNode := influunt.Add(graph, aNode, aNode)

	resultNode := influunt.Cond(graph, influunt.Const(graph, true), aNode, bNode)
	result2Node := influunt.Cond(graph, influunt.Const(graph, false), aNode, bNode)

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(nil, []influunt.Node{resultNode, result2Node})
	if err != nil {
		t.Fatal(err)
	}

	if result[0].(int) != 3 {
		t.Error("Wrong result returned", result)
	}
	if result[1].(int) != 6 {
		t.Error("Wrong result returned", result)
	}
}
