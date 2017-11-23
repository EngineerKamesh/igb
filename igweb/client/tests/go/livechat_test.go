package main

import (
	"github.com/EngineerKamesh/igb/igweb/client/tests/go/caspertest"
	"github.com/gopherjs/gopherjs/js"
)

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

var casper = js.Global.Get("casper")

func main() {

	viewportParams := &caspertest.ViewportParams{Object: js.Global.Get("Object").New()}
	viewportParams.Width = 1440
	viewportParams.Height = 960
	casper.Get("options").Set("viewportSize", viewportParams)

	casper.Get("test").Call("begin", "Live Chat Test Suite", 3, func(test *js.Object) {
		casper.Call("start", "http://localhost:8080/index", wait)
	})

	// Activate the live chat feature
	casper.Call("then", func() {
		casper.Call("click", "#livechatContainer img")
	})

	casper.Call("then", waitChat)

	// Verify that the chat box has opened up
	casper.Call("wait", 1800, func() {
		casper.Call("capture", "screenshots/livechat_test_chatbox_open.png")
		casper.Get("test").Call("assertSelectorHasText", "#chatboxTitle span", "Chat with", "Display chatbox.")
	})

	// Ask a question to the live chat bot, and verify that the chat bot has provided an answer
	casper.Call("then", askQuestion)
	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/livechat_test_answer_question.png")
		casper.Get("test").Call("assertSelectorHasText", "#chatboxConversationContainer", "Isomorphic Go is the methodology to create isomorphic web applications", "Display the answer to \"What is Isomorphic Go?\"")
	})

	// Navigate to another web page (the conversation should be retained)
	casper.Call("then", func() {
		casper.Call("click", "a[href^='/about']")
	})

	casper.Call("then", wait)

	// Verify that the conversation has been retained after navigating to another web page
	casper.Call("wait", 450, func() {
		casper.Call("capture", "screenshots/livechat_test_conversation_retained.png")
		casper.Get("test").Call("assertSelectorHasText", "#chatboxConversationContainer", "Isomorphic Go is the methodology to create isomorphic web applications", "Verify that the conversation is retained when navigating to another page in the website.")
	})

	casper.Call("then", func() {
		casper.Call("click", "#chatboxCloseControl")
	})

	casper.Call("run", func() {
		casper.Get("test").Call("done")
	})

}
