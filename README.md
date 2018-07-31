WORK IN PROGRESS

[![Build Status](https://travis-ci.org/orktes/influunt.svg?branch=master)](https://travis-ci.org/orktes/influunt)

[![GoDoc](https://godoc.org/github.com/orktes/influunt/go?status.svg)](https://godoc.org/github.com/orktes/influunt/go)

# influunt
Dataflow programming for Golang and Python

## Example (create in Python & run in Go)

First create the graph in python and save it to a file

```python
import influunt

with influunt.Graph() as graph:
    # input placeholder
    inputItems = influunt.placeholder()

    # Map items and returns value + 1
    values = inputItems.map(lambda item, i : item.value + 1)

    # Map items and return names
    names = inputItems.map(lambda item, i : item.name)

    # Save graph to file
    influunt.save_graph(graph, "example.graph")
```

Load graph in go and execute

```go
import (
	"fmt"
	"os"

	"github.com/orktes/influunt/go"
	"github.com/orktes/influunt/go/executor"
)

func main() {
	file, _ := os.Open("example.graph")
	graph, _ := influunt.ReadGraph(file)

	input := graph.NodeByName("Placeholder:0", 0)
	names := graph.NodeByName("Map:12", 0)
	values := graph.NodeByName("Map:7", 0)

	exec, _ := executor.NewExecutor(graph)
	results, _ := exec.Run(map[influunt.Node]interface{}{
		*input: []map[string]interface{}{
			{"name": "foo", "value": 100},
			{"name": "bar", "value": 200},
		},
	}, []influunt.Node{*names, *values})

	fmt.Printf("%+v\n", results)
	// [[foo bar] [101 201]]
}
```