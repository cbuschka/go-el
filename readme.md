# go-expr

### A simple expression evaluator created with gocc

## Usage

```
package main

import (
	exprPkg "github.com/cbuschka/go-expr"
	"fmt"
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
```
[see cmd/example/main.go](./cmd/example/main.go)


## License
Copyright (c) 2021 by [Cornelius Buschka](https://github.com/cbuschka).

[MIT](./license.txt)
