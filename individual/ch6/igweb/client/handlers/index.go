package handlers

import (
	"context"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
)

func IndexHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		templateData := templatedata.Index{PageTitle: "IGWEB"}
		env.TemplateSet.Render("index_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
	})
	//InitializeIndexPage(env)
}

func InitializeIndexPage(env *common.Env) {

}
