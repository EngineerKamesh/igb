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
		println("Error encountered when attempting to start the notify cog: ", err)
	}

	liveChatIcon := env.Document.GetElementByID("liveChatIcon").(*dom.HTMLImageElement)
	liveChatIcon.AddEventListener("click", false, func(event dom.Event) {

		chatbox := env.Document.GetElementByID("chatbox")
		if chatbox != nil {
			return
		}
		go chat.StartLiveChat(env)
	})

}
