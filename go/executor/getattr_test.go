package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"

	"github.com/davecgh/go-spew/spew"
)

func TestGetAttr(t *testing.T) {
	graph := influunt.NewGraph()
	attNode := influunt.GetAttr(graph, influunt.Const(graph, map[string]interface{}{"foo": 123}), influunt.Const(graph, "foo"))

	spew.Dump(graph)

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(nil, []influunt.Node{attNode})
	if err != nil {
		t.Fatal(err)
	}

	if result[0] != 123 {
		t.Errorf("Wrong result returned %+v", result)
	}
}
