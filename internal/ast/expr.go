package ast

type Attrib interface{}

type Expr interface {
	Attrib
	Eval(env *EvaluationContext) (interface{}, error)
}
