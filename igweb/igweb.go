package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/EngineerKamesh/igb/igweb/bot"
	"github.com/EngineerKamesh/igb/igweb/chat"
	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/common/datastore"
	"github.com/EngineerKamesh/igb/igweb/endpoints"
	"github.com/EngineerKamesh/igb/igweb/handlers"
	"github.com/EngineerKamesh/igb/igweb/shared/templatefuncs"

	"github.com/EngineerKamesh/igb/igweb/shared/cogs/carousel"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/datepicker"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/liveclock"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/notify"
	"github.com/EngineerKamesh/igb/igweb/shared/cogs/timeago"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/isomorphicgo/isokit"
)

var WebAppRoot string
var WebAppMode string
var WebServerPort string
var DBConnectionString string
var StaticAssetsPath string

// initializeTemplateSet is responsible for initializing the template set on the server-side
func initializeTemplateSet(env *common.Env, oneTimeStaticAssetsGeneration bool) {
	isokit.WebAppRoot = WebAppRoot
	isokit.TemplateFilesPath = WebAppRoot + "/shared/templates"
	isokit.StaticAssetsPath = StaticAssetsPath
	isokit.StaticTemplateBundleFilePath = StaticAssetsPath + "/templates/igweb.tmplbundle"

	if WebAppMode == "production" && oneTimeStaticAssetsGeneration == false {
		isokit.UseStaticTemplateBundleFile = true
		isokit.ShouldBundleStaticAssets = false
	}

	ts := isokit.NewTemplateSet()
	funcMap := template.FuncMap{"rubyformat": templatefuncs.RubyDate, "unixformat": templatefuncs.UnixTime, "productionmode": templatefuncs.IsProduction}
	ts.Funcs = funcMap
	ts.GatherTemplates()
	env.TemplateSet = ts
}

// initializeDatastore is responsible for initializing the datastore for our web application's data persistence needs
func initializeDatastore(env *common.Env) {
	db, err := datastore.NewDatastore(datastore.REDIS, DBConnectionString)
	if err != nil {
		log.Fatalf("Could not connect to the Redis Datastore! Encountered the following error when attempting to create a datastore instance: ", err)
	}
	env.DB = db
}

func initializeSessionstore(env *common.Env) {
	if _, err := os.Stat("/tmp/igweb-sessions"); os.IsNotExist(err) {
		os.Mkdir("/tmp/igweb-sessions", 711)
	}
	env.Store = sessions.NewFilesystemStore("/tmp/igweb-sessions", []byte(os.Getenv("IGWEB_HASH_KEY")))
}

// initializeCogs is responsible for initializing all the cogs that will be used in the web app
func initializeCogs(ts *isokit.TemplateSet) {
	timeago.NewTimeAgo().CogInit(ts)
	liveclock.NewLiveClock().CogInit(ts)
	datepicker.NewDatePicker().CogInit(ts)
	carousel.NewCarousel().CogInit(ts)
	notify.NewNotify().CogInit(ts)
	isokit.BundleStaticAssets()
}

// startChatHub is responsible for starting the chat hub used for the live chat feature
func startChatHub(hub *chat.Hub) {
	go hub.Run()
}

// registerRoutes is responsible for regisetering the server-side request handlers
func registerRoutes(env *common.Env, r *mux.Router, hub *chat.Hub) {
	// Register Handler for Loading the Sample Dataset into a Redis Instance
	r.Handle("/config/load-sample-data", handlers.LoadSampleDataHandler(env)).Methods("GET")

	// Register Handlers for Client-Side JavaScript Application
	if WebAppMode != "production" {
		r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot)).Methods("GET")
		r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot)).Methods("GET")
	}

	// Register handler for the delivery of the template bundle
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(env)).Methods("POST")

	// Register Request Handlers for GopherJS Examples
	r.Handle("/front-end-examples-demo", handlers.FrontEndExamplesHandler(env)).Methods("GET")
	r.Handle("/cars-demo", handlers.CarsDemoHandler(env)).Methods("GET")
	r.Handle("/cars-data", handlers.CarsDataHandler(env)).Methods("POST")
	r.Handle("/localstorage-demo", handlers.LocalStorageDemoHandler(env)).Methods("GET")

	// Register Request Handlers for IGWEB
	r.Handle("/", handlers.IndexHandler(env)).Methods("GET")
	r.Handle("/index", handlers.IndexHandler(env)).Methods("GET")
	r.Handle("/lowercase-text", handlers.LowercaseTextTransformHandler(env)).Methods("POST")
	r.Handle("/products", handlers.ProductsHandler(env)).Methods("GET")
	r.Handle("/product-detail/{productTitle}", handlers.ProductDetailHandler(env)).Methods("GET")
	r.Handle("/about", handlers.AboutHandler(env)).Methods("GET")
	r.Handle("/contact", handlers.ContactHandler(env)).Methods("GET", "POST")
	r.Handle("/contact-confirmation", handlers.ContactConfirmationHandler(env)).Methods("GET")
	r.Handle("/shopping-cart", handlers.ShoppingCartHandler(env))

	// Register REST API Endpoints for IGWEB
	r.Handle("/restapi/get-gopher-team", endpoints.GetGopherTeamEndpoint(env)).Methods("GET")
	r.Handle("/restapi/get-products", endpoints.GetProductsEndpoint(env)).Methods("GET")
	r.Handle("/restapi/get-product-detail/{productTitle}", endpoints.GetProductDetailEndpoint(env)).Methods("GET")
	r.Handle("/restapi/get-cart-items", endpoints.GetShoppingCartItemsEndpoint(env)).Methods("GET")
	r.Handle("/restapi/add-item-to-cart", endpoints.AddItemToShoppingCartEndpoint(env)).Methods("PUT")
	r.Handle("/restapi/remove-item-from-cart", endpoints.RemoveItemFromShoppingCartEndpoint(env)).Methods("DELETE")
	r.Handle("/restapi/contact-form", endpoints.ContactFormEndpoint(env)).Methods("POST")
	r.Handle("/restapi/get-agent-info", endpoints.GetAgentInfoEndpoint(env, hub.ChatBot())).Methods("GET")

	// Register Request Handler for the Websocket Connection used by the live chat feature
	r.Handle("/ws", chat.ServeWs(hub))
}

func main() {

	env := common.Env{}

	if WebAppRoot == "" {
		fmt.Println("The IGWEB_APP_ROOT environment variable must be set before the web server instance can be started.")
		os.Exit(1)
	}

	shouldGenerateStaticAssetsAndExit := flag.Bool("generate-static-assets", false, "One time generation of static assets for production use.")
	flag.Parse()

	if *shouldGenerateStaticAssetsAndExit == true {
		generateStaticAssetsAndExit(&env)
	}

	initializeTemplateSet(&env, false)
	initializeCogs(env.TemplateSet)
	initializeDatastore(&env)
	initializeSessionstore(&env)

	chatbot := bot.NewAgentCase()
	hub := chat.NewHub(chatbot)
	startChatHub(hub)

	r := mux.NewRouter()
	registerRoutes(&env, r, hub)

	// Register Request Handler for Static Assetcs
	fs := http.FileServer(http.Dir(StaticAssetsPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.Handle("/", r)
	http.ListenAndServe(":"+WebServerPort, nil)

}

func init() {

	WebAppRoot = os.Getenv("IGWEB_APP_ROOT")
	WebAppMode = os.Getenv("IGWEB_MODE")
	WebServerPort = os.Getenv("IGWEB_SERVER_PORT")
	DBConnectionString = os.Getenv("IGWEB_DB_CONNECTION_STRING")

	// Set the default web server port if it hasn't been set already
	if WebServerPort == "" {
		WebServerPort = "8080"
	}

	// Set the default database connection string
	if DBConnectionString == "" {
		DBConnectionString = "localhost:6379"
	}

	StaticAssetsPath = WebAppRoot + "/static"

}

func generateStaticAssetsAndExit(env *common.Env) {

	fmt.Print("Generating static assets...")
	isokit.ShouldMinifyStaticAssets = true
	isokit.ShouldBundleStaticAssets = true
	initializeTemplateSet(env, true)
	initializeCogs(env.TemplateSet)
	fmt.Println("Done")
	os.Exit(0)

}
