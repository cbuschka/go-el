package ast

import (
	"fmt"
	"regexp"
)

type RelationalOp int

const (
	EQ      = RelationalOp(0)
	NE      = RelationalOp(1)
	MATCHES = RelationalOp(2)
)

type RelationalExpr struct {
	A  Expr
	B  Expr
	Op RelationalOp
}

func (this *RelationalExpr) Eval(env *EvaluationContext) (interface{}, error) {
	aVal, err := this.A.Eval(env)
	if err != nil {
		return nil, err
	}

	bVal, err := this.B.Eval(env)
	if err != nil {
		return nil, err
	}

	switch this.Op {
	case EQ:
		return aVal == bVal, nil
	case NE:
		return aVal == bVal, nil
	case MATCHES:
		aStrVal, aIsString := aVal.(string)
		if !aIsString {
			return nil, fmt.Errorf("left side of =~ requires string")
		}
		bStrVal, bIsString := bVal.(string)
		if !bIsString {
			return nil, fmt.Errorf("right side of =~ requires pattern string")
		}

		return regexp.MatchString(bStrVal, aStrVal)
	}
	return nil, fmt.Errorf("unknown op")
}

func NewEqExpr(a, b Attrib) (*RelationalExpr, error) {
	return &RelationalExpr{a.(Expr), b.(Expr), EQ}, nil
}

func NewNeExpr(a, b Attrib) (*RelationalExpr, error) {
	return &RelationalExpr{a.(Expr), b.(Expr), NE}, nil
}

func NewMatchesExpr(a, b Attrib) (*RelationalExpr, error) {
	return &RelationalExpr{a.(Expr), b.(Expr), MATCHES}, nil
}
