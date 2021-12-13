package ast

import (
	"github.com/cbuschka/go-expr/internal/generated/token"
)

type LookupExpr struct {
	Attrib
	Identifier string
}

func NewLookupExpr(a Attrib) (*LookupExpr, error) {
	return &LookupExpr{Identifier: string(a.(*token.Token).Lit)}, nil
}

func (e *LookupExpr) Eval(env *EvaluationContext) (interface{}, error) {
	val, err := env.Resolve(e.Identifier)
	if err != nil {
		return nil, err
	}

	return val, nil
}
