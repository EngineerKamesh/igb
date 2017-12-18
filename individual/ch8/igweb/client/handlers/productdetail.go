package handlers

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
	"github.com/EngineerKamesh/igb/igweb/shared/templatedata"
	"github.com/isomorphicgo/isokit"
	"honnef.co/go/js/xhr"
)

func ProductDetailHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		routeVars := ctx.Value(isokit.RouteVarsKey("Vars")).(map[string]string)
		productTitle := routeVars[`product-detail/{productTitle}`]
		productChannel := make(chan *models.Product)
		go FetchProductDetail(productChannel, productTitle)
		product := <-productChannel
		templateData := &templatedata.ProductDetail{PageTitle: product.Name, Product: product}
		env.TemplateSet.Render("product_detail_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
		InitializeProductDetailPage(env)
	})
}

func FetchProductDetail(productChannel chan *models.Product, productTitle string) {

	data, err := xhr.Send("GET", "/restapi/get-product-detail"+"/"+productTitle, nil)
	if err != nil {
		println("Encountered error: ", err)
		println(err)
	}
	var product *models.Product
	json.NewDecoder(strings.NewReader(string(data))).Decode(&product)

	productChannel <- product
}

func InitializeProductDetailPage(env *common.Env) {

	buttons := env.Document.GetElementsByClassName("addToCartButton")
	for _, button := range buttons {
		button.AddEventListener("click", false, handleAddToCartButtonClickEvent)
	}
}
