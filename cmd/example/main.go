package main

import (
	"fmt"
	"github.com/cbuschka/go-el"
)

func main() {
	expr := el.MustCompile("true")
	env := map[string]interface{}{}
	result, err := expr.Evaluate(env)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result is '%v'\n", result)
}
