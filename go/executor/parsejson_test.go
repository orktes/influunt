package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestParseJSON(t *testing.T) {
	graph := influunt.NewGraph()
	jsonNode := influunt.ParseJSON(graph, influunt.Const(graph, `{"foo":123}`))

	executor, err := executor.NewExecutor(graph)
	if err != nil {
		t.Fatal(err)
	}

	result, err := executor.Run(nil, []influunt.Node{jsonNode})
	if err != nil {
		t.Fatal(err)
	}

	m := result[0].(map[string]interface{})

	if m["foo"] != float64(123) {
		t.Errorf("Wrong result returned %+v", m["foo"])
	}
}
