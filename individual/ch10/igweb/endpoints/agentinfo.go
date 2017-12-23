package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/EngineerKamesh/igb/igweb/bot"
	"github.com/EngineerKamesh/igb/igweb/common"
)

func GetAgentInfoEndpoint(env *common.Env, chatbot bot.Bot) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		m := make(map[string]string)
		m["AgentName"] = chatbot.Name()
		m["AgentTitle"] = chatbot.Title()
		m["AgentThumbImagePath"] = chatbot.ThumbnailPath()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)
	})
}
