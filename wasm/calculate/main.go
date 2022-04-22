package calculate

import (
	"strconv"
	"syscall/js"
	dom "honnef.co/go/js/dom/v2"
)

func Add(this js.Value, i []js.Value) any {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

	value1 := js.Global().Get("document").Call("querySelector", value1Selector).Get("value").String()
    value2 := js.Global().Get("document").Call("querySelector", value2Selector).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

	resultEl := *(dom.GetWindow().Document().QuerySelector(resultSelector)).(*dom.HTMLInputElement)

    resultEl.Set("value", int1 + int2)

	println("add: ", resultEl.Value())

	return ""
}

func Subtract(this js.Value, i []js.Value) any {
	value1Selector := "#" + i[0].String()
	value2Selector := "#" + i[1].String()
	resultSelector := "#" + i[2].String()

	value1 := js.Global().Get("document").Call("querySelector", value1Selector).Get("value").String()
    value2 := js.Global().Get("document").Call("querySelector", value2Selector).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

	resultEl := *(dom.GetWindow().Document().QuerySelector(resultSelector)).(*dom.HTMLInputElement)

    resultEl.Set("value", int1 - int2)

	println("subtract: ", resultEl.Value())

	return ""
}

func GetOperator() string {
	return (*(dom.GetWindow().Document().QuerySelector("#operator")).(*dom.HTMLSelectElement)).Value()
}