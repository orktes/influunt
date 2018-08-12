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

    # Save graph as a model and define inputs and outputs
    influunt.save_model(
		graph, 
		{"input": inputItems},
		{"values": values, "names": names}, 
		"example.model"
	)
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
	file, _ := os.Open("example.model")
	model, _ := influunt.ReadModel(file)

	exec, _ := executor.NewModelExecutor(model)
	results, _ := exec.Run(map[string]interface{}{
		"input": []map[string]interface{}{
			{"name": "foo", "value": 100},
			{"name": "bar", "value": 200},
		},
	})

	fmt.Printf("%+v\n", results["names"])
	// [foo bar]

	fmt.Printf("%+v\n", results["values"])
	// [101 201]
}
```