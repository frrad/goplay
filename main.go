package main

import (
	"fmt"
	"syscall/js"

	"github.com/frrad/gofmt"
)

func main() {
	c := make(chan struct{}, 0)
	registerCallback()
	<-c
}

func registerCallback() {
	js.Global().Set("gofmt", js.NewCallback(format))
}

func format([]js.Value) {
	updateMessage("loading input...")
	userInput := getInputByName("code-input")

	updateMessage("formatting...")
	ans, err := gofmt.Fmt(userInput)
	if err != nil {
		updateMessage(fmt.Sprintf("failed to format: %+v", err))
		return
	}

	updateMessage("formatted!")
	setValue("code-input", ans)
}

func updateMessage(msg string) {
	writeHTML("messageP", msg)
}

func writeHTML(id, html string) {
	set(id, "innerHTML", html)
}

func setValue(id, val string) {
	set(id, "value", val)
}
func set(id, which, val string) {
	js.Global().Get("document").Call("getElementById", id).Set(which, val)
}

func getInputByName(inName string) string {
	return js.Global().Get("document").Call("getElementById", inName).Get("value").String()
}
