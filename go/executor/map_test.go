package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestMap(t *testing.T) {
	graph := influunt.NewGraph()
	list := influunt.Const(graph, []int{1, 2, 3})
	res := influunt.Map(graph, list, func(item influunt.Node, index influunt.Node) influunt.Node {
		return influunt.Add(graph, item, influunt.Const(graph, 1))
	})

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(nil, []influunt.Node{res})
	if err != nil {
		t.Fatal(err)
	}

	resultInts := result[0].([]int)
	if resultInts[0] != 2 {
		t.Error("Wrong value returned", resultInts)
	}
	if resultInts[1] != 3 {
		t.Error("Wrong value returned", resultInts)
	}

	if resultInts[2] != 4 {
		t.Error("Wrong value returned", resultInts)
	}

}
