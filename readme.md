# go-el

### A simple expression language created with gocc

## Usage

```
package main

import (
	"github.com/cbuschka/go-el"
	"fmt"
)

func main() {
	expr := el.MustCompile("true")
	env := el.NewEvaluationContext()
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
