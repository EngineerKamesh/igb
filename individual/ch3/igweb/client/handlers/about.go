package handlers

import (
	"context"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/isomorphicgo/isokit"
)

func AboutHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {

		env.TemplateSet.Render("about_content", &isokit.RenderParams{Data: nil, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: "About"})
		//InitializeAboutPage(env)
	})
}

func InitializeAboutPage(env *common.Env) {
}
