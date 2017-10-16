package main

import (
	"html/template"
	"strings"

	"github.com/isomorphicgo/isokit"
	"honnef.co/go/js/dom"

	"github.com/EngineerKamesh/igb/igweb/client/carsdemo"
	"github.com/EngineerKamesh/igb/igweb/client/common"
	"github.com/EngineerKamesh/igb/igweb/client/gopherjsprimer"
	"github.com/EngineerKamesh/igb/igweb/client/handlers"
	"github.com/EngineerKamesh/igb/igweb/client/localstoragedemo"
	"github.com/EngineerKamesh/igb/igweb/shared/forms"
	"github.com/EngineerKamesh/igb/igweb/shared/templatefuncs"
)

func initializeEventHandlers(env *common.Env) {

	l := strings.Split(env.Window.Location().Pathname, "/")
	pageName := l[1]

	if pageName == "" {
		pageName = "index"
	}

	switch pageName {

	case "front-end-examples":
		gopherjsprimer.InitializeEventHandlers()

	case "cars-demo":
		carsdemo.InitializePage()

	case "localstorage-demo":
		localstoragedemo.InitializeEventHandlers()

	case "index":
		handlers.InitializeDefaultHandlers(env)
		handlers.InitializeIndexEventHandlers(env)

	case "about":
		handlers.InitializeDefaultHandlers(env)
		handlers.InitializeAboutEventHandlers(env)

	case "products":
		handlers.InitializeDefaultHandlers(env)
		handlers.InitializeProductsEventHandlers(env)

	case "product-detail":
		handlers.InitializeProductDetailEventHandlers(env)

	case "shopping-cart":
		handlers.InitializeDefaultHandlers(env)
		handlers.InitializeShoppingCartEventHandlers(env)

	case "contact":
		contactForm := forms.NewContactForm(nil)
		handlers.InitializeContactFormEventHandlers(env, contactForm)

	}
}

func registerRoutes(env *common.Env) {

	r := isokit.NewRouter()
	r.Handle("/index", handlers.IndexHandler(env))
	r.Handle("/products", handlers.ProductsHandler(env))
	r.Handle("/product-detail/{productTitle}", handlers.ProductDetailHandler(env))
	r.Handle("/about", handlers.AboutHandler(env))
	r.Handle("/contact", handlers.ContactHandler(env))
	r.Handle("/shopping-cart", handlers.ShoppingCartHandler(env))
	r.Listen()
	env.Router = r
}

func run() {
	println("IGWEB Client Application")
	templateSetChannel := make(chan *isokit.TemplateSet)

	funcMap := template.FuncMap{"rubyformat": templatefuncs.RubyDate, "unixformat": templatefuncs.UnixTime}

	go isokit.FetchTemplateBundleWithSuppliedFunctionMap(templateSetChannel, funcMap)
	ts := <-templateSetChannel

	env := common.Env{}
	env.TemplateSet = ts

	env.Window = dom.GetWindow()
	env.Document = dom.GetWindow().Document()
	env.PrimaryContent = env.Document.GetElementByID("primaryContent")

	registerRoutes(&env)
	initializeEventHandlers(&env)

}

func main() {

	var D = dom.GetWindow().Document().(dom.HTMLDocument)
	switch readyState := D.ReadyState(); readyState {
	case "loading":
		D.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		println("Unexpected document.ReadyState value!")
	}

}
