package executor

import (
	"testing"
)

func Test_getMathFunc(t *testing.T) {
	add := getMathFunc("add")

	res, err := add(1, 2)

	if err != nil {
		t.Error(err)
	}

	if res.(int) != 3 {
		t.Error("invalid response received")
	}
}
