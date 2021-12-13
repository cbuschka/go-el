package ast

import (
	"github.com/cbuschka/go-el/internal/generated/token"
)

type CallExpr struct {
	Attrib
	Base       Expr
	Identifier string
	ArgList    *ArgList
}

func NewMethodCallExpr(base, identifier Attrib, argList Attrib) (*CallExpr, error) {
	return &CallExpr{Base: base.(Expr), Identifier: string(identifier.(*token.Token).Lit), ArgList: argList.(*ArgList)}, nil
}

func NewFunctionCallExpr(identifier Attrib, argList Attrib) (*CallExpr, error) {
	return &CallExpr{Base: nil, Identifier: string(identifier.(*token.Token).Lit)}, nil
}

func toArgs(env *EvaluationContext, argList *ArgList, args []interface{}) ([]interface{}, error) {
	if argList == nil {
		return args, nil
	}

	if argList.ArgList != nil {
		var err error
		args, err = toArgs(env, argList.ArgList, args)
		if err != nil {
			return nil, err
		}
	}

	argVal, err := argList.Expr.Eval(env)
	if err != nil {
		return nil, err
	}
	args = append(args, argVal)

	return args, nil
}

func (e *CallExpr) Eval(env *EvaluationContext) (interface{}, error) {

	baseVal, err := e.Base.Eval(env)
	if err != nil {
		return nil, err
	}

	args, err := toArgs(env, e.ArgList, []interface{}{})
	if err != nil {
		return nil, err
	}

	val, err := env.CallOn(e.Identifier, args, baseVal)
	if err != nil {
		return nil, err
	}

	return val, nil
}

type ArgList struct {
	Attrib
	ArgList *ArgList
	Expr    Expr
}

func NewArgList(a, b Attrib) (*ArgList, error) {
	if a == nil {
		return &ArgList{ArgList: nil, Expr: b.(Expr)}, nil
	}

	argList := a.(*ArgList)
	return &ArgList{ArgList: argList, Expr: b.(Expr)}, nil
}
