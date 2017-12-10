package handlers

import (
	"context"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/isomorphicgo/isokit"
)

func ContactHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {

		env.TemplateSet.Render("contact_content", &isokit.RenderParams{Data: nil, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: "Contact"})
	})
	// InitializeContactPage(env)
}

func InitializeContactPage(env *common.Env) {

}
