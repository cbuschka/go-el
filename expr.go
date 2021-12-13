package el

import (
	"github.com/cbuschka/go-el/internal/ast"
	"github.com/cbuschka/go-el/internal/generated/lexer"
	"github.com/cbuschka/go-el/internal/generated/parser"
)

type Expression interface {
	EvaluateWithContext(evalContext EvaluationContext) (interface{}, error)
	Evaluate() (interface{}, error)
}

type EvaluationContext interface {
	SetFunction(name string, function interface{})
	SetValue(name string, value interface{})
}

type expression struct {
	program *ast.Expr
}

func MustCompile(script string) Expression {
	expr, err := Compile(script)
	if err != nil {
		panic(err)
	}

	return expr
}

func Compile(script string) (Expression, error) {
	l := lexer.NewLexer([]byte(script))
	p := parser.NewParser()
	st, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	boolExpr := st.(ast.Expr)

	return Expression(&expression{program: &boolExpr}), nil
}

func NewEvaluationContext() EvaluationContext {
	return ast.NewEvaluationContext()
}

func (e *expression) Evaluate() (interface{}, error) {
	evalContext := NewEvaluationContext()
	return e.EvaluateWithContext(evalContext)
}

func (e *expression) EvaluateWithContext(evalContext EvaluationContext) (interface{}, error) {
	return (*e.program).Eval(evalContext.(*ast.EvaluationContext))
}
