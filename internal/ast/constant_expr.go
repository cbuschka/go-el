package ast

import (
	"github.com/cbuschka/go-el/internal/generated/token"
	"strconv"
)

type ConstantExpr struct {
	Value interface{}
}

func (e *ConstantExpr) Eval(*EvaluationContext) (interface{}, error) {
	return e.Value, nil
}

func NewConstantStringExpr(a Attrib) (*ConstantExpr, error) {
	literal := string(a.(*token.Token).Lit)
	literal = literal[1 : len(literal)-1]
	return &ConstantExpr{Value: literal}, nil
}

func NewConstantBoolExpr(value bool) (*ConstantExpr, error) {
	return &ConstantExpr{Value: value}, nil
}

func NewConstantIntExpr(a Attrib) (*ConstantExpr, error) {
	value, err := strconv.Atoi(string(a.(*token.Token).Lit))
	if err != nil {
		return nil, err
	}
	return &ConstantExpr{Value: value}, nil
}
