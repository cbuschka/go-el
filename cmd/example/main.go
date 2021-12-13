package main

import (
	"fmt"
	"github.com/cbuschka/go-el"
)

func main() {
	expr := el.MustCompile("true")
	result, err := expr.Evaluate()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result is '%v'\n", result)
}
