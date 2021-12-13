package ast

import (
	"fmt"
	"reflect"
)

type EvaluationContext struct {
	Root      map[string]interface{}
	Functions map[string]interface{}
}

func NewEvaluationContext() *EvaluationContext {
	return &EvaluationContext{Root: map[string]interface{}{}, Functions: map[string]interface{}{}}
}

func (c *EvaluationContext) AddFunction(name string, f interface{}) {
	c.Functions[name] = f
}

func (c *EvaluationContext) AddValue(name string, value interface{}) {
	c.Root[name] = value
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

func (e *EvaluationContext) CallMethodOn(name string, args []interface{}, target interface{}) (interface{}, error) {
	baseObject := reflect.ValueOf(target)

	methodMeta := baseObject.MethodByName(name)
	if methodMeta == *new(reflect.Value) {
		return nil, fmt.Errorf("method %s not found", name)
	}

	argValues := []reflect.Value{}
	for _, arg := range args {
		argValues = append(argValues, reflect.ValueOf(arg))
	}

	result := methodMeta.Call(argValues)
	if result == nil || len(result) != 1 {
		return nil, fmt.Errorf("%s returned invalid result %v", name, result)
	}

	return result[0].Interface(), nil
}

func (c *EvaluationContext) CallFunction(identifier string, args []interface{}) (interface{}, error) {
	function, found := c.Functions[identifier]
	if !found {
		return nil, fmt.Errorf("function %s npot found", identifier)
	}

	argValues := []reflect.Value{}
	for _, arg := range args {
		argValues = append(argValues, reflect.ValueOf(arg))
	}

	functionValue := reflect.ValueOf(function)
	result := functionValue.Call(argValues)
	if result == nil || len(result) != 1 {
		return nil, fmt.Errorf("%s returned invalid result %v", identifier, result)
	}

	return result[0].Interface(), nil
}
