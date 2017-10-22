package main

import (
	"html/template"
	"log"
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

func initializePage(env *common.Env) {

	l := strings.Split(env.Window.Location().Pathname, "/")
	routeName := l[1]

	if routeName == "" {
		routeName = "index"
	}

	if strings.Contains(routeName, "-demo") == false {
		handlers.InitializePageLayoutControls(env)
	}

	switch routeName {

	case "front-end-examples-demo":
		gopherjsprimer.InitializePage()

	case "cars-demo":
		carsdemo.InitializePage()

	case "localstorage-demo":
		localstoragedemo.InitializePage()

	case "index":
		handlers.InitializeIndexPage(env)

	case "about":
		handlers.InitializeAboutPage(env)

	case "products":
		handlers.InitializeProductsPage(env)

	case "product-detail":
		handlers.InitializeProductDetailPage(env)

	case "shopping-cart":
		handlers.InitializeShoppingCartPage(env)

	case "contact":
		contactForm := forms.NewContactForm(nil)
		handlers.InitializeContactPage(env, contactForm)

	default:
		log.Println("Encountered unknown route name: ", routeName)
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

	// Fetch the template set
	templateSetChannel := make(chan *isokit.TemplateSet)
	funcMap := template.FuncMap{"rubyformat": templatefuncs.RubyDate, "unixformat": templatefuncs.UnixTime, "productionmode": templatefuncs.IsProduction}
	go isokit.FetchTemplateBundleWithSuppliedFunctionMap(templateSetChannel, funcMap)
	ts := <-templateSetChannel

	env := common.Env{}
	env.TemplateSet = ts
	env.Window = dom.GetWindow()
	env.Document = dom.GetWindow().Document()
	env.PrimaryContent = env.Document.GetElementByID("primaryContent")
	env.Location = env.Window.Location()

	registerRoutes(&env)
	initializePage(&env)
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
		println("Encountered unexpected document ready state value!")
	}

}
