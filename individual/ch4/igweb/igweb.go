package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/common/datastore"
	"github.com/EngineerKamesh/igb/igweb/endpoints"
	"github.com/EngineerKamesh/igb/igweb/handlers"
	"github.com/EngineerKamesh/igb/igweb/shared/templatefuncs"
	"github.com/gorilla/mux"
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

	ts := isokit.NewTemplateSet()
	funcMap := template.FuncMap{"rubyformat": templatefuncs.RubyDate, "unixformat": templatefuncs.UnixTime}
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

// registerRoutes is responsible for regisetering the server-side request handlers
func registerRoutes(env *common.Env, r *mux.Router) {
	// Register Handler for Loading the Sample Dataset into a Redis Instance
	r.Handle("/config/load-sample-data", handlers.LoadSampleDataHandler(env)).Methods("GET")

	// Register Handlers for Client-Side JavaScript Application
	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot)).Methods("GET")
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot)).Methods("GET")

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
	r.Handle("/about", handlers.AboutHandler(env)).Methods("GET")
	r.Handle("/contact", handlers.ContactHandler(env)).Methods("GET", "POST")

	// Register REST API Endpoints for IGWEB
	r.Handle("/restapi/get-gopher-team", endpoints.GetGopherTeamEndpoint(env)).Methods("GET")

}

func main() {

	env := common.Env{}

	if WebAppRoot == "" {
		fmt.Println("The IGWEB_APP_ROOT environment variable must be set before the web server instance can be started.")
		os.Exit(1)
	}

	initializeTemplateSet(&env, false)
	initializeDatastore(&env)

	r := mux.NewRouter()
	registerRoutes(&env, r)

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
