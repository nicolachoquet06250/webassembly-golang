package calculate

import (
	"strconv"
	"syscall/js"
	dom "honnef.co/go/js/dom/v2"
)

func GetOperator() string {
	return (*(dom.GetWindow().Document().QuerySelector("#operator")).(*dom.HTMLSelectElement)).Value()
}

func GetInputElement(selector string) dom.HTMLInputElement {
	v := dom.GetWindow().Document().QuerySelector(selector)
	vi := v.(*dom.HTMLInputElement)
	return *vi
}

func GetValue(selector string) (int, error) {
	v := (*(dom.GetWindow().Document().QuerySelector(selector)).(*dom.HTMLInputElement)).Value()

	return strconv.Atoi(v)
}

func Add(this js.Value, i []js.Value) interface{} {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

    int1, _ := GetValue(value1Selector)
    int2, _ := GetValue(value2Selector)

	resultEl := GetInputElement(resultSelector)

    resultEl.Set("value", int1 + int2)

	rValue, _ := GetValue(resultSelector)

	println("subtract: ", rValue)

	return rValue
}

func Subtract(this js.Value, i []js.Value) interface{} {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

    int1, _ := GetValue(value1Selector)
    int2, _ := GetValue(value2Selector)

	resultEl := GetInputElement(resultSelector)

    resultEl.Set("value", int1 - int2)

	rValue, _ := GetValue(resultSelector)

	println("subtract: ", rValue)

	return rValue
}