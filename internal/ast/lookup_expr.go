package ast

import (
	"fmt"
	"github.com/cbuschka/go-expr/internal/generated/token"
)

type LookupExpr struct {
	Attrib
	Identifier string
}

func NewLookupExpr(a Attrib) (*LookupExpr, error) {
	return &LookupExpr{Identifier: string(a.(*token.Token).Lit)}, nil
}

func (e *LookupExpr) Eval(env map[string]interface{}) (interface{}, error) {
	val, found := env[e.Identifier]
	if !found {
		return false, fmt.Errorf("%s not found", e.Identifier)
	}

	return val, nil
}
