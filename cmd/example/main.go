package main

import (
	"fmt"
	exprPkg "github.com/cbuschka/go-expr"
)

func main() {
	expr := exprPkg.MustCompile("true")
	env := map[string]interface{}{}
	result, err := expr.Evaluate(env)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result is '%v'\n", result)
}
