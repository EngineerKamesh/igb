package main

import (
	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

type ChatFormParams struct {
	*js.Object
	FirstName   string `js:"firstName"`
	LastName    string `js:"lastName"`
	Email       string `js:"email"`
	MessageBody string `js:"messageBody"`
}

var wait = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#primaryContent")
	return nil
})

var waitChat = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Call("waitForSelector", "#chatbox")
	return nil
})

var askQuestion = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {

	this.Call("sendKeys", "input#chatboxInputField", "What is Isomorphic Go?")
	this.Call("sendKeys", "input#chatboxInputField", casper.Get("page").Get("event").Get("key").Get("Enter"))
	return nil
})

var currentQuestion = ""
var casper = js.Global.Get("casper")

func main() {

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Live Chat Test", 3, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/index", wait)
	})

	casper.Call("then", func() {
		casper.Call("click", "#livechatContainer img")
	})

	casper.Call("then", waitChat)

	casper.Call("then", func() {
		casper.Call("capture", "screenshots/livechat_test_chatbox_open.png")
		casper.Get("test").Call("assertSelectorHasText", "#chatboxTitle span", "Chat with", "Display chatbox.")
	})

	casper.Call("then", askQuestion)
	casper.Call("then", func() {
		casper.Call("capture", "screenshots/livechat_test_answer_question.png")
		casper.Get("test").Call("assertSelectorHasText", "#chatboxConversationContainer", "Isomorphic Go is the methodology to create isomorphic web applications", "Display the answer to \"What is Isomorphic Go?\"")
	})

	casper.Call("then", func() {
		casper.Call("click", "a[href^='/about']")
	})

	casper.Call("then", wait)

	casper.Call("then", func() {
		casper.Call("capture", "screenshots/livechat_test_conversation_retained.png")
		casper.Get("test").Call("assertSelectorHasText", "#chatboxConversationContainer", "Isomorphic Go is the methodology to create isomorphic web applications", "Verify that the conversation is retained when navigating to another page in the website.")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
