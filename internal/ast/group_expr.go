package ast

type GroupExpr struct {
	A Expr
}

func NewGroupExpr(a Attrib) (*GroupExpr, error) {
	return &GroupExpr{a.(Expr)}, nil
}

func (g *GroupExpr) Eval(env map[string]interface{}) (interface{}, error) {
	return g.A.Eval(env)
}
