package chat

import (
	"flag"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket/websocketjs"
	"github.com/isomorphicgo/isokit"
	"honnef.co/go/js/dom"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var ws *websocketjs.WebSocket
var agentInfo map[string]string

const ENTERKEY int = 13

func StartLiveChat(env *common.Env) {

	var err error
	ws, err = websocketjs.New("ws://localhost:8080/ws") // Does not block.
	if err != nil {
		// handle error
	}

	agentInfo = make(map[string]string)
	agentInfo["AgentName"] = "Case"
	agentInfo["AgentThumbImagePath"] = "/static/images/chat/Case.png"
	agentInfo["AgentTitle"] = "Resident Isomorphic Gopher Agent"

	chatContainer := env.Document.GetElementByID("chatboxContainer").(*dom.HTMLDivElement)
	env.TemplateSet.Render("partials/chatbox_partial", &isokit.RenderParams{Data: agentInfo, Disposition: isokit.PlacementReplaceInnerContents, Element: chatContainer})

	InitializeChatEventHandlers(env)

	for {
		select {}
	}
}

func UpdateChatBox(env *common.Env, message string, sender string) {

	m := make(map[string]string)
	m["Name"] = sender
	m["Message"] = message
	conversationContainer := env.Document.GetElementByID("chatboxConversationContainer").(*dom.HTMLDivElement)
	env.TemplateSet.Render("partials/livechatmsg_partial", &isokit.RenderParams{Data: m, Disposition: isokit.PlacementAppendTo, Element: conversationContainer})
	scrollHeight := conversationContainer.Underlying().Get("scrollHeight")
	conversationContainer.Underlying().Set("scrollTop", scrollHeight)
}

func handleOnMessage(env *common.Env, ev *js.Object) {

	response := ev.Get("data").String()
	UpdateChatBox(env, response, agentInfo["AgentName"])
}

func ChatSendMessage(env *common.Env, message string) {
	ws.Send([]byte(message))
	UpdateChatBox(env, message, "Me")
}

func CloseChat() {
	ws.Close()
}

func InitializeChatEventHandlers(env *common.Env) {

	msgInput := env.Document.GetElementByID("chatboxInputField").(*dom.HTMLInputElement)
	msgInput.AddEventListener("keypress", false, func(event dom.Event) {
		if event.Underlying().Get("keyCode").Int() == ENTERKEY {
			event.PreventDefault()
			go ChatSendMessage(env, msgInput.Value)
			msgInput.Value = ""
		}

	})

	closeControl := env.Document.GetElementByID("chatboxCloseControl").(*dom.HTMLDivElement)
	closeControl.AddEventListener("click", false, func(event dom.Event) {
		CloseChat()
		chatboxContainer := env.Document.GetElementByID("chatboxContainer").(*dom.HTMLDivElement)
		chatboxContainer.RemoveChild(chatboxContainer.ChildNodes()[0])
	})

	ws.AddEventListener("message", false, func(ev *js.Object) {
		handleOnMessage(env, ev)
	})

}
