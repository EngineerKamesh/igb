package handlers

import (
	"github.com/EngineerKamesh/igb/igweb/client/chat"
	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/notify"
	"honnef.co/go/js/dom"
)

func InitializePageLayoutControls(env *common.Env) {

	n := notify.NewNotify()
	err := n.Start()
	if err != nil {
		println("Error encontered when attempting to start the notify cog: ", err)
	}

	chatbox := env.Document.GetElementByID("chatbox")

	if chatbox != nil {
		return
	}

	liveChatIcon := env.Document.GetElementByID("liveChatIcon").(*dom.HTMLImageElement)
	liveChatIcon.AddEventListener("click", false, func(event dom.Event) {
		go chat.StartLiveChat(env)
	})

}
