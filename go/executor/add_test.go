package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestAdd(t *testing.T) {
	graph := influunt.NewGraph()
	resultNode := influunt.Add(graph, influunt.Const(graph, 1), influunt.Const(graph, 2))
	resultNode2 := influunt.Add(graph, resultNode, influunt.Const(graph, 2.5))
	resultNode3 := influunt.Add(graph, influunt.Const(graph, "foo"), influunt.Const(graph, "bar"))

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(nil, []influunt.Node{resultNode, resultNode2, resultNode3})
	if err != nil {
		t.Fatal(err)
	}

	if result[0].(int) != 3 {
		t.Error("Wrong result returned", result)
	}
	if result[1].(float64) != 5.5 {
		t.Error("Wrong result returned", result)
	}
	if result[2].(string) != "foobar" {
		t.Error("Wrong result returned", result)
	}
}

func BenchmarkAdd(b *testing.B) {
	graph := influunt.NewGraph()
	resultNode := influunt.Add(graph, influunt.Const(graph, 1), influunt.Const(graph, 2))
	resultNode2 := influunt.Add(graph, resultNode, influunt.Const(graph, 2.5))
	resultNode3 := influunt.Add(graph, influunt.Const(graph, "foo"), influunt.Const(graph, "bar"))

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		b.Fatal(err)
	}

	outputs := []influunt.Node{resultNode, resultNode2, resultNode3}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		executor.Run(nil, outputs)
	}
}
