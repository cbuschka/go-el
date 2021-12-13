package ast

import (
	"github.com/cbuschka/go-el/internal/generated/token"
)

type CallExpr struct {
	Attrib
	Base       Expr
	Identifier string
}

func NewCallExpr(base, identifier Attrib) (*CallExpr, error) {
	return &CallExpr{Base: base.(Expr), Identifier: string(identifier.(*token.Token).Lit)}, nil
}

func (e *CallExpr) Eval(env *EvaluationContext) (interface{}, error) {

	baseVal, err := e.Base.Eval(env)
	if err != nil {
		return nil, err
	}

	val, err := env.CallOn(e.Identifier, baseVal)
	if err != nil {
		return nil, err
	}

	return val, nil
}
