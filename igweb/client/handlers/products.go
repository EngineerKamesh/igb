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

func ProductsHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {

		productsChannel := make(chan []*models.Product)
		go FetchProducts(productsChannel)
		products := <-productsChannel
		templateData := &templatedata.Products{PageTitle: "Products", Products: products}
		env.TemplateSet.Render("products_content", &isokit.RenderParams{Data: templateData, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent, PageTitle: templateData.PageTitle})
		InitializeProductsPage(env)
		env.Router.RegisterLinks("#primaryContent a")
	})
}

func FetchProducts(productsChannel chan []*models.Product) {

	data, err := xhr.Send("GET", "/restapi/get-products", nil)
	if err != nil {
		println("Encountered error: ", err)
		return
	}
	var products []*models.Product
	json.NewDecoder(strings.NewReader(string(data))).Decode(&products)

	productsChannel <- products
}

func InitializeProductsPage(env *common.Env) {

	buttons := env.Document.GetElementsByClassName("addToCartButton")
	for _, button := range buttons {
		button.AddEventListener("click", false, handleAddToCartButtonClickEvent)
	}

}
