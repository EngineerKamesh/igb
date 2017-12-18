package handlers

import (
	"context"
	"encoding/json"
	"strings"

	"honnef.co/go/js/xhr"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
)

func AboutHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {

		gopherTeamChannel := make(chan []*models.Gopher)
		go FetchGopherTeam(gopherTeamChannel)
		gophers := <-gopherTeamChannel
		templateData := templatedata.About{PageTitle: "About", Gophers: gophers}
		env.TemplateSet.Render("about_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
		//InitializeAboutPage(env)
	})
}

func FetchGopherTeam(gopherTeamChannel chan []*models.Gopher) {
	data, err := xhr.Send("GET", "/restapi/get-gopher-team", nil)
	if err != nil {
		println("Encountered error: ", err)
	}
	var gophers []*models.Gopher
	json.NewDecoder(strings.NewReader(string(data))).Decode(&gophers)
	gopherTeamChannel <- gophers
}

func InitializeAboutPage(env *common.Env) {
}
