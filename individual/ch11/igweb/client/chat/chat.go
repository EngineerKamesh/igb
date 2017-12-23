package chat

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket/websocketjs"
	"github.com/isomorphicgo/isokit"
	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

var ws *websocketjs.WebSocket
var agentInfo map[string]string

const ENTERKEY int = 13

func getServerPort(env *common.Env) string {

	if env.Location.Port != "" {
		return env.Location.Port
	}

	if env.Location.Protocol == "https" {
		return "443"
	} else {
		return "80"
	}

}

func StartLiveChat(env *common.Env) {

	agentInfoChannel := make(chan map[string]string)
	go GetAgentInfoRequest(agentInfoChannel)
	agentInfo = <-agentInfoChannel

	var err error
	serverEndpoint := "ws://" + env.Location.Hostname + ":" + getServerPort(env) + "/ws"
	ws, err = websocketjs.New(serverEndpoint)
	if err != nil {
		println("Encountered error when attempting to connect to the websocket: ", err)
	}

	chatContainer := env.Document.GetElementByID("chatboxContainer").(*dom.HTMLDivElement)
	chatContainer.SetClass("containerPulse")

	env.TemplateSet.Render("partials/chatbox_partial", &isokit.RenderParams{Data: agentInfo, Disposition: isokit.PlacementReplaceInnerContents, Element: chatContainer})

	chatboxHeaderBar := env.Document.GetElementByID("chatboxHeaderBar").(*dom.HTMLDivElement)
	chatboxHeaderBar.SetClass("chatboxHeader")

	InitializeChatEventHandlers(env)

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

func HandleOnMessage(env *common.Env, ev *js.Object) {

	response := ev.Get("data").String()
	UpdateChatBox(env, response, agentInfo["AgentName"])
}

func ChatSendMessage(env *common.Env, message string) {
	ws.Send([]byte(message))
	UpdateChatBox(env, message, "Me")
}

func CloseChat(env *common.Env) {
	ws.Close()
	chatboxContainer := env.Document.GetElementByID("chatboxContainer").(*dom.HTMLDivElement)
	chatboxContainer.RemoveChild(chatboxContainer.ChildNodes()[0])

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
		CloseChat(env)
	})

	ws.AddEventListener("message", false, func(ev *js.Object) {
		go HandleOnMessage(env, ev)
	})

	ws.AddEventListener("close", false, func(ev *js.Object) {

		chatboxContainer := env.Document.GetElementByID("chatboxContainer").(*dom.HTMLDivElement)
		if len(chatboxContainer.ChildNodes()) > 0 {
			go HandleDisconnection(env)
		}
	})

	env.Window.AddEventListener("offline", false, func(event dom.Event) {
		go HandleDisconnection(env)
	})

}

func HandleDisconnection(env *common.Env) {

	chatContainer := env.Document.GetElementByID("chatboxContainer").(*dom.HTMLDivElement)
	chatContainer.SetClass("")

	chatboxHeaderBar := env.Document.GetElementByID("chatboxHeaderBar").(*dom.HTMLDivElement)
	chatboxHeaderBar.SetClass("chatboxHeader disconnected")

	chatboxTitleDiv := env.Document.GetElementByID("chatboxTitle").(*dom.HTMLDivElement)
	if chatboxTitleDiv != nil {
		titleSpan := chatboxTitleDiv.ChildNodes()[0].(*dom.HTMLSpanElement)
		if titleSpan != nil {
			var countdown uint64 = 6
			tickerForCountdown := time.NewTicker(1 * time.Second)
			timerToCloseChat := time.NewTimer(6 * time.Second)
			go func() {
				for _ = range tickerForCountdown.C {
					atomic.AddUint64(&countdown, ^uint64(0))
					safeCountdownValue := atomic.LoadUint64(&countdown)
					titleSpan.SetInnerHTML("Disconnected! - Closing LiveChat in " + strconv.FormatUint(safeCountdownValue, 10) + " seconds.")
				}
			}()
			go func() {
				<-timerToCloseChat.C
				tickerForCountdown.Stop()
				CloseChat(env)
			}()
		}
	}
}

func GetAgentInfoRequest(agentInfoChannel chan map[string]string) {
	data, err := xhr.Send("GET", "/restapi/get-agent-info", nil)
	if err != nil {
		println("Encountered error: ", err)
	}
	var agentInfo map[string]string
	json.NewDecoder(strings.NewReader(string(data))).Decode(&agentInfo)
	agentInfoChannel <- agentInfo
}
