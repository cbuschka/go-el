package ast

import (
	"fmt"
	"reflect"
)

type EvaluationContext struct {
	Root interface{}
}

func NewEvaluationContext(root interface{}) *EvaluationContext {
	return &EvaluationContext{Root: root}
}

func (e *EvaluationContext) Resolve(name string) (interface{}, error) {
	return e.ResolveFrom(name, e.Root)
}

func (e *EvaluationContext) ResolveFrom(name string, parent interface{}) (interface{}, error) {

	baseMap, isMap := parent.(map[string]interface{})
	if isMap {
		val, found := baseMap[name]
		if !found {
			return nil, fmt.Errorf("%s not found", name)
		}

		return val, nil
	}

	typeMeta := reflect.TypeOf(parent)
	if typeMeta.Kind() == reflect.Ptr {
		typeMeta = typeMeta.Elem()
	}
	baseObject := reflect.ValueOf(parent)
	if baseObject.Kind() == reflect.Ptr {
		baseObject = baseObject.Elem()
	}

	fieldMeta := baseObject.FieldByName(name)
	if fieldMeta == *new(reflect.Value) {
		return false, fmt.Errorf("%s not found", name)
	}

	return fieldMeta.Interface(), nil

}
