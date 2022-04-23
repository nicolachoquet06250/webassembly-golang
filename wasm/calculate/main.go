package calculate

import (
	dom "honnef.co/go/js/dom/v2"
	"strconv"
	"syscall/js"
)

type HTMLInputElement = dom.HTMLInputElement
type HTMLSelectElement = dom.HTMLSelectElement
type Value = js.Value

func GetOperator() string {
	return (*(dom.GetWindow().Document().QuerySelector("#operator")).(*HTMLSelectElement)).Value()
}

func getInputElement(selector string) HTMLInputElement {
	v := dom.GetWindow().Document().QuerySelector(selector)
	vi := v.(*HTMLInputElement)

	return *vi
}

func getValue(selector string) (int, error) {
	v := (*(dom.GetWindow().Document().QuerySelector(selector)).(*HTMLInputElement)).Value()

	return strconv.Atoi(v)
}

func add(this Value, i []Value) interface{} {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

	int1, _ := getValue(value1Selector)
	int2, _ := getValue(value2Selector)

	resultEl := getInputElement(resultSelector)

	resultEl.Set("value", int1+int2)

	rValue, _ := getValue(resultSelector)

	println("subtract: ", rValue)

	return rValue
}

func subtract(this Value, i []Value) interface{} {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

	int1, _ := getValue(value1Selector)
	int2, _ := getValue(value2Selector)

	resultEl := getInputElement(resultSelector)

	resultEl.Set("value", int1-int2)

	rValue, _ := getValue(resultSelector)

	println("subtract: ", rValue)

	return rValue
}

func RegisterCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
