package handlers

import (
	"github.com/EngineerKamesh/igb/igweb/client/chat"
	"github.com/EngineerKamesh/igb/igweb/client/common"
	"honnef.co/go/js/dom"
)

func InitializePageLayoutControls(env *common.Env) {

	chatbox := env.Document.GetElementByID("chatbox")

	if chatbox != nil {
		return
	}

	liveChatIcon := env.Document.GetElementByID("liveChatIcon").(*dom.HTMLImageElement)
	liveChatIcon.AddEventListener("click", false, func(event dom.Event) {
		go chat.StartLiveChat(env)
	})

}
