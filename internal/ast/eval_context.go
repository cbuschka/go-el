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

func (e *EvaluationContext) CallOn(name string, target interface{}) (interface{}, error) {
	baseObject := reflect.ValueOf(target)

	methodMeta := baseObject.MethodByName(name)
	if methodMeta == *new(reflect.Value) {
		return nil, fmt.Errorf("method %s not found", name)
	}

	result := methodMeta.Call([]reflect.Value{})
	if result == nil || len(result) != 1 {
		return nil, fmt.Errorf("%s returned invalid result %v", name, result)
	}

	return result[0].Interface(), nil
}
