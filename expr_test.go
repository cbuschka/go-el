package el

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompileAndEvaluate(t *testing.T) {

	expr, err := CompileExpression("true")
	if err != nil {
		t.Fatal(err)
		return
	}

	result, err := expr.Evaluate(NewEvaluationContext())
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

	result, err := expr.Evaluate(NewEvaluationContext())
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

	env := NewEvaluationContext()
	env.AddValue("flag", true)
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

	env := NewEvaluationContext()
	env.AddValue("flag", true)
	env.AddValue("flag2", true)
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

	env := NewEvaluationContext()
	env.AddValue("value", 1)
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

	env := NewEvaluationContext()
	env.AddValue("value", "yay")
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

	env := NewEvaluationContext()
	aMap := map[string]interface{}{}
	aMap["struct"] = TestStruct{Value: "yay"}
	env.AddValue("map", aMap)
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

	env := NewEvaluationContext()
	env.AddValue("struct", TestStruct{})
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "Yay!", result)
}

func (t TestStruct) Say(message string, toWhom string, how string) string {
	return fmt.Sprintf("Said: %s To: %s How: %s", message, toWhom, how)
}

func TestCompileAndEvaluateCallWithArgs(t *testing.T) {

	expr, err := CompileExpression("struct.Say(\"Hello World!\", \"All of the world\", \"Very loud\")")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := NewEvaluationContext()
	env.AddValue("struct", TestStruct{})
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "Said: Hello World! To: All of the world How: Very loud", result)
}

func TestCompileAndEvaluateFunctionCallWithArgs(t *testing.T) {

	expr, err := CompileExpression("say(\"huhu\")")
	if err != nil {
		t.Fatal(err)
		return
	}

	env := NewEvaluationContext()
	env.AddFunction("say", func(what string) string {
		return fmt.Sprintf("Said: %s", what)
	})
	result, err := expr.Evaluate(env)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "Said: huhu", result)
}
