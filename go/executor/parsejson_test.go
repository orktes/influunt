package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestParseJSON(t *testing.T) {
	graph := influunt.NewGraph()
	json := influunt.Const(graph, `{"foo":123}`)
	jsonNode := influunt.ParseJSON(graph, json)
	attNode := influunt.GetAttr(graph, jsonNode, influunt.Const(graph, "foo"))

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(nil, []influunt.Node{attNode})
	if err != nil {
		t.Fatal(err)
	}

	if result[0] != float64(123) {
		t.Errorf("Wrong result returned %+v", attNode)
	}
}
