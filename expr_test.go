package el

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompileAndEvaluate(t *testing.T) {

	expr, err := Compile("true")
	if err != nil {
		t.Fatal(err)
		return
	}

	result, err := expr.Evaluate()
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateFalse(t *testing.T) {

	expr, err := Compile("false")
	if err != nil {
		t.Fatal(err)
		return
	}

	result, err := expr.Evaluate()
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, false, result)
}

func TestCompileAndEvaluateLookup(t *testing.T) {

	expr, err := Compile("flag")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetValue("flag", true)
	result, err := expr.EvaluateWithContext(evalCtx)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateCompositeAnd(t *testing.T) {

	expr, err := Compile("( flag && false ) || (flag2 && flag )")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetValue("flag", true)
	evalCtx.SetValue("flag2", true)
	result, err := expr.EvaluateWithContext(evalCtx)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateCompareInt(t *testing.T) {

	expr, err := Compile("value == 1")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetValue("value", 1)
	result, err := expr.EvaluateWithContext(evalCtx)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, true, result)
}

func TestCompileAndEvaluateCompareStrings(t *testing.T) {

	expr, err := Compile("value == \"yay\"")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetValue("value", "yay")
	result, err := expr.EvaluateWithContext(evalCtx)
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

	expr, err := Compile("map.struct.Value")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	aMap := map[string]interface{}{}
	aMap["struct"] = TestStruct{Value: "yay"}
	evalCtx.SetValue("map", aMap)
	result, err := expr.EvaluateWithContext(evalCtx)
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

	expr, err := Compile("struct.String()")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetValue("struct", TestStruct{})
	result, err := expr.EvaluateWithContext(evalCtx)
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

	expr, err := Compile("struct.Say(\"Hello World!\", \"All of the world\", \"Very loud\")")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetValue("struct", TestStruct{})
	result, err := expr.EvaluateWithContext(evalCtx)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "Said: Hello World! To: All of the world How: Very loud", result)
}

func TestCompileAndEvaluateFunctionCallWithArgs(t *testing.T) {

	expr, err := Compile("say(\"huhu\")")
	if err != nil {
		t.Fatal(err)
		return
	}

	evalCtx := NewEvaluationContext()
	evalCtx.SetFunction("say", func(what string) string {
		return fmt.Sprintf("Said: %s", what)
	})
	result, err := expr.EvaluateWithContext(evalCtx)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "Said: huhu", result)
}
