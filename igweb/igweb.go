package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/EngineerKamesh/igb/igweb/chat"
	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/common/datastore"
	"github.com/EngineerKamesh/igb/igweb/endpoints"
	"github.com/EngineerKamesh/igb/igweb/handlers"
	"github.com/EngineerKamesh/igb/igweb/shared/templatefuncs"

	"github.com/EngineerKamesh/igb/igweb/shared/cogs/button"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/carousel"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/datepicker"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/liveclock"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/timeago"
	"github.com/gorilla/mux"
	"github.com/isomorphicgo/isokit"
)

var WebAppRoot string = os.Getenv("IGWEB_APP_ROOT")

func main() {

	if WebAppRoot == "" {
		fmt.Println("The IGWEB_APP_ROOT environment variable must be set before the web server instance can be started.")
		os.Exit(1)
	}

	db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")
	if err != nil {
		log.Fatalf("Could not connect to the Redis Datastore! Encountered this error when attempting to create a datastore instance: ", err)
	}

	env := common.Env{}
	isokit.WebAppRoot = WebAppRoot
	isokit.TemplateFilesPath = WebAppRoot + "/shared/templates"
	isokit.StaticAssetsPath = WebAppRoot + "/static"
	ts := isokit.NewTemplateSet()
	funcMap := template.FuncMap{"rubyformat": templatefuncs.RubyDate, "unixformat": templatefuncs.UnixTime}
	ts.Funcs = funcMap

	ts.GatherTemplates()
	InitializeCogs(ts)
	env.TemplateSet = ts
	env.DB = db

	hub := chat.NewHub()
	go hub.Run()

	r := mux.NewRouter()

	// Register Handler for Loading the Sample Dataset into a Redis Instance
	r.Handle("/config/load-sample-data", handlers.LoadSampleDataHandler(&env)).Methods("GET")

	// Register Handlers for Client-Side Application
	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot))
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	// Register Request Handlers for GopherJS Examples
	r.Handle("/front-end-examples-demo", handlers.FrontEndExamplesHandler(&env)).Methods("GET")
	r.Handle("/cars-demo", handlers.CarsDemoHandler(&env)).Methods("GET")
	r.Handle("/cars-data", handlers.CarsDataHandler(&env)).Methods("POST")
	r.Handle("/localstorage-demo", handlers.LocalStorageDemoHandler(&env)).Methods("GET")

	// Register Request Handlers for IGWEB
	r.Handle("/", handlers.IndexHandler(&env)).Methods("GET")
	r.Handle("/index", handlers.IndexHandler(&env)).Methods("GET")
	r.Handle("/lowercase-text", handlers.LowercaseTextTransformHandler(&env))
	r.Handle("/products", handlers.ProductsHandler(&env)).Methods("GET")
	r.Handle("/product-detail/{productTitle}", handlers.ProductDetailHandler(&env)).Methods("GET")
	r.Handle("/about", handlers.AboutHandler(&env)).Methods("GET")
	r.Handle("/contact", handlers.ContactHandler(&env)).Methods("GET", "POST")
	r.Handle("/contact-confirmation", handlers.ContactConfirmationHandler(&env)).Methods("GET")
	r.Handle("/shopping-cart", handlers.ShoppingCartHandler(&env))

	// Register Request Handler for the Websocket Connection (used in the Live Chat feature)
	http.Handle("/ws", chat.ServeWs(hub))

	// Register REST API Endpoints for IGWEB
	r.Handle("/restapi/get-gopher-team", endpoints.GetGopherTeamEndpoint(&env)).Methods("GET")
	r.Handle("/restapi/get-products", endpoints.GetProductsEndpoint(&env)).Methods("GET")
	r.Handle("/restapi/get-product-detail/{productTitle}", endpoints.GetProductDetailEndpoint(&env)).Methods("GET")
	r.Handle("/restapi/get-cart-items", endpoints.GetShoppingCartItemsEndpoint(&env)).Methods("GET")
	r.Handle("/restapi/add-item-to-cart", endpoints.AddItemToShoppingCartEndpoint(&env)).Methods("PUT")
	r.Handle("/restapi/remove-item-from-cart", endpoints.RemoveItemFromShoppingCartEndpoint(&env)).Methods("DELETE")
	r.Handle("/restapi/contact-form", endpoints.ContactFormEndpoint(&env)).Methods("POST")

	// Register Request Handler for Static Assetcs
	fs := http.FileServer(http.Dir(WebAppRoot + "/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

// InitializeCogs takes care of initializing all the cogs that we will be using in the web app
func InitializeCogs(ts *isokit.TemplateSet) {

	button.NewButton().CogInit(ts)
	liveclock.NewLiveClock().CogInit(ts)
	timeago.NewTimeAgo().CogInit(ts)
	datepicker.NewDatePicker().CogInit(ts)
	carousel.NewCarousel().CogInit(ts)
	isokit.BundleStaticAssets()
}
