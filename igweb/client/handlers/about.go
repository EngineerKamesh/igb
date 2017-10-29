package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"honnef.co/go/js/xhr"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/timeago"
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
		InitializeAboutPage(env)
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
	humanReadableDivs := env.Document.GetElementsByClassName("humanReadableDate")
	for _, div := range humanReadableDivs {
		unixTimestamp, err := strconv.ParseInt(div.GetAttribute("data-starttimeunix"), 10, 64)
		if err != nil {
			log.Println("Encountered error when attempting to parse int64 from string:", err)
		}
		t := time.Unix(unixTimestamp, 0)
		humanTime := timeago.NewTimeAgo()
		humanTime.CogInit(env.TemplateSet)
		humanTime.SetID(div.ID())
		humanTime.SetTime(t)
		err = humanTime.Start()
		if err != nil {
			println("Encountered the following error when attempting to start the timeago cog: ", err)
		}
	}
}
