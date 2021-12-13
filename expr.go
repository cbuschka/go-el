package el

import (
	"github.com/cbuschka/go-el/internal/ast"
	"github.com/cbuschka/go-el/internal/generated/lexer"
	"github.com/cbuschka/go-el/internal/generated/parser"
)

type Expression struct {
	program *ast.Expr
}

type EvaluationContext interface {
	AddFunction(name string, function interface{})
	AddValue(name string, value interface{})
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

func NewEvaluationContext() EvaluationContext {
	return ast.NewEvaluationContext()
}

func (e *Expression) Evaluate(env EvaluationContext) (interface{}, error) {
	return (*e.program).Eval(env.(*ast.EvaluationContext))
}
