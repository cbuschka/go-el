package ast

import (
	"github.com/cbuschka/go-el/internal/generated/token"
)

type DerefExpr struct {
	Attrib
	Base       Expr
	Identifier string
}

func NewDerefExpr(base, identifier Attrib) (*DerefExpr, error) {
	return &DerefExpr{Base: base.(Expr), Identifier: string(identifier.(*token.Token).Lit)}, nil
}

func (e *DerefExpr) Eval(env *EvaluationContext) (interface{}, error) {

	baseVal, err := e.Base.Eval(env)
	if err != nil {
		return nil, err
	}

	val, err := env.ResolveFrom(e.Identifier, baseVal)
	if err != nil {
		return nil, err
	}

	return val, nil
}
