package ast

import (
	"fmt"
)

type BoolBinaryOp int

const (
	OR  = BoolBinaryOp(1)
	AND = BoolBinaryOp(2)
)

type BoolBinaryExpr struct {
	A  Expr
	B  Expr
	Op BoolBinaryOp
}

func NewAndExpr(a, b Attrib) (*BoolBinaryExpr, error) {
	return &BoolBinaryExpr{a.(Expr), b.(Expr), AND}, nil
}

func NewOrExpr(a, b Attrib) (*BoolBinaryExpr, error) {
	return &BoolBinaryExpr{a.(Expr), b.(Expr), OR}, nil
}

func (this *BoolBinaryExpr) Eval(env map[string]interface{}) (interface{}, error) {
	aVal, err := this.A.Eval(env)
	if err != nil {
		return nil, err
	}

	bVal, err := this.B.Eval(env)
	if err != nil {
		return nil, err
	}

	aBoolVal, aIsBool := aVal.(bool)
	if !aIsBool {
		return false, fmt.Errorf("not bool '%v'", aBoolVal)
	}

	bBoolVal, bIsBool := bVal.(bool)
	if !bIsBool {
		return false, fmt.Errorf("not bool '%v'", bBoolVal)
	}

	switch this.Op {
	case OR:
		return aBoolVal || bBoolVal, nil
	case AND:
		return aBoolVal && bBoolVal, nil
	}
	return this.A.Eval(env)
}
