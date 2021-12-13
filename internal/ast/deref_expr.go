package ast

import (
	"fmt"
	"github.com/cbuschka/go-expr/internal/generated/token"
)

type DerefExpr struct {
	Attrib
	Base       Expr
	Identifier string
}

func NewDerefExpr(base, identifier Attrib) (*DerefExpr, error) {
	return &DerefExpr{Base: base.(Expr), Identifier: string(identifier.(*token.Token).Lit)}, nil
}

func (e *DerefExpr) Eval(env map[string]interface{}) (interface{}, error) {
	baseVal, err := e.Base.Eval(env)
	if err != nil {
		return nil, err
	}

	baseMap, isMap := baseVal.(map[string]interface{})
	if isMap {
		val, found := baseMap[e.Identifier]
		if !found {
			return nil, fmt.Errorf("%s not found", e.Identifier)
		}

		return val, nil
	}

	return false, fmt.Errorf("deref %s from %v not implemented", e.Identifier, baseVal)
}
