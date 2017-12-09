package handlers

import (
	"context"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/isomorphicgo/isokit"
)

func ProductsHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {

		env.TemplateSet.Render("products_content", &isokit.RenderParams{Data: nil, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: "Products"})
		//InitializeProductsPage(env)
	})
}

func InitializeProductsPage(env *common.Env) {

}
