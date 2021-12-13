package expr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompileAndEvaluate(t *testing.T) {

	expr, err := CompileExpression("true")
	if err != nil {
		t.Fatal(err)
		return
	}

	result, err := expr.Evaluate(map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateFalse(t *testing.T) {

	expr, err := CompileExpression("false")
	if err != nil {
		t.Fatal(err)
		return
	}

	result, err := expr.Evaluate(map[string]interface{}{})
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, false, result)
}

func TestCompileAndEvaluateLookup(t *testing.T) {

	expr, err := CompileExpression("flag")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := map[string]interface{}{}
	env["flag"] = true
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateCompositeAnd(t *testing.T) {

	expr, err := CompileExpression("( flag && false ) || (flag2 && flag )")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := map[string]interface{}{}
	env["flag"] = true
	env["flag2"] = true
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateCompareInt(t *testing.T) {

	expr, err := CompileExpression("value == 1")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := map[string]interface{}{}
	env["value"] = 1
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateCompareStrings(t *testing.T) {

	expr, err := CompileExpression("value == \"yay\"")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := map[string]interface{}{}
	env["value"] = "yay"
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

type TestStruct struct {
	Value string
}

func TestCompileAndEvaluateDeref(t *testing.T) {

	expr, err := CompileExpression("map.struct.Value")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := map[string]interface{}{}
	aMap := map[string]interface{}{}
	aMap["struct"] = TestStruct{Value: "yay"}
	env["map"] = aMap
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "yay", result)
}

func (t TestStruct) String() string {
	return "Yay!"
}

func TestCompileAndEvaluateCall(t *testing.T) {

	expr, err := CompileExpression("struct.String()")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := map[string]interface{}{}
	env["struct"] = TestStruct{}
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "Yay!", result)
}
