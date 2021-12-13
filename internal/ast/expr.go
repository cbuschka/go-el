package ast

type Attrib interface{}

type Expr interface {
	Attrib
	Eval(env map[string]interface{}) (interface{}, error)
}
