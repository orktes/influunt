package executor_test

import (
	"testing"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func TestModelExecutor(t *testing.T) {
	graph := influunt.NewGraph()
	inputNode := influunt.Placeholder(graph)

	model := &influunt.Model{
		Graph: graph,
		Inputs: map[string]influunt.Node{
			"input": inputNode,
		},
		Outputs: map[string]influunt.Node{
			"echo": inputNode,
		},
	}

	modelExecutor, err := executor.NewModelExecutor(model)
	if err != nil {
		t.Error(err)
	}

	res, err := modelExecutor.Run(map[string]interface{}{
		"input": "foo",
	})
	if err != nil {
		t.Error(err)
	}

	if res["echo"].(string) != "foo" {
		t.Error("Wrong result received", res)
	}

}
