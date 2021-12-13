package expr

import (
	"github.com/cbuschka/go-expr/internal/ast"
	"github.com/cbuschka/go-expr/internal/generated/lexer"
	"github.com/cbuschka/go-expr/internal/generated/parser"
)

type Expression struct {
	program *ast.Expr
}

func MustCompile(script string) *Expression {
	expr, err := CompileExpression(script)
	if err != nil {
		panic(err)
	}

	return expr
}

func CompileExpression(script string) (*Expression, error) {
	l := lexer.NewLexer([]byte(script))
	p := parser.NewParser()
	st, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	boolExpr := st.(ast.Expr)

	return &Expression{program: &boolExpr}, nil
}

func (e *Expression) Evaluate(env map[string]interface{}) (interface{}, error) {
	return (*e.program).Eval(env)
}
