package ast

import (
	"fmt"
	"github.com/cbuschka/go-expr/internal/generated/token"
	"reflect"
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

	typeMeta := reflect.TypeOf(baseVal)
	if typeMeta.Kind() == reflect.Ptr {
		typeMeta = typeMeta.Elem()
	}
	baseObject := reflect.ValueOf(baseVal)
	if baseObject.Kind() == reflect.Ptr {
		baseObject = baseObject.Elem()
	}

	fieldMeta := baseObject.FieldByName(e.Identifier)
	if fieldMeta == *new(reflect.Value) {
		return false, fmt.Errorf("%s not found in %v", e.Identifier, baseVal)
	}

	return fieldMeta.Interface(), nil
}
