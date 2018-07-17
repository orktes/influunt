package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestPlaceholder(t *testing.T) {
	graph := influunt.NewGraph()
	inputNode := influunt.Placeholder(graph)

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(map[influunt.Node]interface{}{
		inputNode: 1,
	}, []influunt.Node{inputNode})
	if err != nil {
		t.Fatal(err)
	}

	if result[0].(int) != 1 {
		t.Error("Wrong result returned", result)
	}
}
