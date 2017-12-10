package gopherjsprimer

import (
	"honnef.co/go/js/dom"
)

func InitializePage() {

	d := dom.GetWindow().Document()

	messageInput := d.GetElementByID("messageInput").(*dom.HTMLInputElement)

	alertButtonJS := d.GetElementByID("alertMessageJSGlobal").(*dom.HTMLButtonElement)
	alertButtonJS.AddEventListener("click", false, func(event dom.Event) {
		DisplayAlertMessageJSGlobal(messageInput.Value)
	})

	alertButtonDOM := d.GetElementByID("alertMessageDOM").(*dom.HTMLButtonElement)
	alertButtonDOM.AddEventListener("click", false, func(event dom.Event) {
		DisplayAlertMessageDOM(messageInput.Value)
	})

	showGopherButton := d.GetElementByID("showGopher").(*dom.HTMLButtonElement)
	showGopherButton.AddEventListener("click", false, func(event dom.Event) {
		ShowIsomorphicGopher()
	})

	hideGopherButton := d.GetElementByID("hideGopher").(*dom.HTMLButtonElement)
	hideGopherButton.AddEventListener("click", false, func(event dom.Event) {
		HideIsomorphicGopher()
	})

	builtinDemoButton := d.GetElementByID("builtinDemoButton").(*dom.HTMLButtonElement)
	builtinDemoButton.AddEventListener("click", false, func(event dom.Event) {
		builtinDemo(event.Target())
	})

	lowercaseTransformButton := d.GetElementByID("lowercaseTransformButton").(*dom.HTMLButtonElement)
	lowercaseTransformButton.AddEventListener("click", false, func(event dom.Event) {
		go lowercaseTextTransformer()
	})

}
