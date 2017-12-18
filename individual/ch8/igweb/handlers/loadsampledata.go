package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/EngineerKamesh/igb/igweb/common"
	"github.com/EngineerKamesh/igb/igweb/shared/models"
)

func LoadSampleDataHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		molly := &models.Gopher{Name: "Molly", Title: "Founder and CEO", Biodata: "Maecenas dui neque, sagittis non purus eget, vehicula imperdiet quam. Nulla condimentum urna libero, non viverra ligula ultrices vel. Sed dictum et felis id ultrices. Suspendisse imperdiet mollis est nec commodo. Vivamus volutpat nunc non quam iaculis vestibulum non ut sapien. Morbi venenatis mi eu ligula tempus ultrices. Aenean eget tellus vitae nulla accumsan rhoncus non ut eros. Nam ultricies interdum risus et pretium. Maecenas dignissim purus nec auctor vestibulum.", ImageURI: "/static/images/gophers/Molly.png", StartTime: time.Date(2017, 5, 24, 17, 9, 0, 0, time.UTC)}
		caseg := &models.Gopher{Name: "Case", Title: "Resident Isomorphic Gopher Agent", Biodata: "Nullam gravida, dolor ac sollicitudin elementum, augue ex luctus libero, ac fringilla turpis odio nec lacus. Vestibulum fermentum tincidunt tortor a fermentum. Nam at pellentesque justo. Fusce ac quam massa. Suspendisse a tellus est. Suspendisse eget leo sollicitudin, volutpat nunc sit amet, placerat massa. Vestibulum maximus urna velit, eget vehicula sapien mollis at. Proin venenatis lobortis mattis. Duis consequat lectus lectus, ac sagittis leo bibendum at.", ImageURI: "/static/images/gophers/Case.png", StartTime: time.Date(2017, 7, 14, 18, 36, 0, 0, time.UTC)}
		wintermute := &models.Gopher{Name: "Wintermute", Title: "Developer", Biodata: "Quisque imperdiet velit non metus commodo dignissim. Nulla auctor tempus felis, sed euismod risus dictum quis. Aenean nunc erat, rutrum eget quam eu, vehicula commodo ante. Nam lacus justo, laoreet ac molestie quis, tempus ut lacus. Morbi in lectus quis justo convallis euismod vitae eu orci. Sed euismod viverra tincidunt. Curabitur a convallis ex. Praesent vitae purus et eros malesuada sollicitudin. Fusce id luctus sapien. Aliquam placerat et arcu in efficitur. Nulla eu consectetur ante.", ImageURI: "/static/images/gophers/Wintermute.png", StartTime: time.Date(2017, 9, 19, 19, 16, 0, 0, time.UTC)}

		gophers := []*models.Gopher{molly, caseg, wintermute}
		err := env.DB.CreateGopherTeam(gophers)
		if err != nil {
			log.Println("Encountered an error when attempting to create the Gopher Team: ", err)
			w.Write([]byte("Failed to load sample data!"))
			return
		}

		chair := &models.Product{SKU: "1", Name: "Hand Crafted Chair", Description: "A hand-crafted chair.", ThumbnailPreviewURI: "/static/images/products/chair_thumb.png", ImagePreviewURI: "/static/images/products/chair.png", Price: 79.99, Route: "/product-detail/hand-crafted-chair", SummaryDetail: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer sed rutrum lacus, a tincidunt felis. Quisque vitae tincidunt nibh. Sed vel ultricies dui. Vivamus porttitor sem et sapien pretium suscipit. Aenean tortor nulla, tincidunt a varius vel, feugiat vel ligula. Aenean mattis ligula a massa mollis auctor. Curabitur porttitor ipsum id auctor luctus. Fusce aliquam sem magna, eget mollis nisl luctus ut. Etiam ac pellentesque justo."}
		coffeemug := &models.Product{SKU: "2", Name: "Coffee Mug", Description: "A clever coffee mug.", ThumbnailPreviewURI: "/static/images/products/coffeemug_thumb.jpg", ImagePreviewURI: "/static/images/products/coffeemug.jpg", Price: 14.99, Route: "/product-detail/coffee-mug", SummaryDetail: "Mauris tincidunt quam vel purus mollis, eleifend egestas turpis aliquet. Ut nec turpis consequat risus luctus rhoncus. Nunc posuere, nisi et convallis placerat, leo ipsum ultrices nisi, quis aliquam enim dui eu magna. Etiam porta lorem at ante ultrices, vel placerat ante rhoncus. Donec et magna non risus hendrerit condimentum quis eget enim. Nunc congue mattis risus non maximus. Curabitur molestie elit at est maximus dapibus. In hac habitasse platea dictumst."}
		desklamp := &models.Product{SKU: "3", Name: "Desk Lamp", Description: "A handy desk lamp.", ThumbnailPreviewURI: "/static/images/products/desklamp_thumb.jpg", ImagePreviewURI: "/static/images/products/desklamp.jpg", Price: 39.99, Route: "/product-detail/desk-lamp", SummaryDetail: "Cras sem tortor, mollis nec pretium et, molestie sit amet orci. Sed accumsan porttitor turpis sit amet lacinia. Fusce leo leo, condimentum sed tincidunt a, tristique sit amet metus. Cras gravida justo vel consequat ornare. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Aenean volutpat ipsum at purus feugiat, a rhoncus risus consectetur. Curabitur eget dignissim nibh, sed dignissim augue. Quisque dignissim eget nibh sit amet lobortis. Sed eget ornare dolor. Duis bibendum urna sed turpis posuere tempor."}
		flowervase := &models.Product{SKU: "4", Name: "Flower Vase", Description: "A beautiful flower vase.", ThumbnailPreviewURI: "/static/images/products/flowervase_thumb.jpg", ImagePreviewURI: "/static/images/products/flowervase.jpg", Price: 10.89, Route: "/product-detail/flower-vase", SummaryDetail: "Nunc dapibus mollis efficitur. Duis nec malesuada magna, non elementum ipsum. Nam vel eros porttitor, sagittis lorem non, commodo nisi. Donec dictum ex convallis condimentum placerat. Vivamus sodales, arcu eget congue consequat, quam sem congue nunc, nec maximus ipsum lorem in odio. Etiam congue porttitor odio, et cursus odio consequat eu. Morbi tempor et urna id maximus."}
		lawnmower := &models.Product{SKU: "5", Name: "Lawn Mower", Description: "A handy lawn mower.", ThumbnailPreviewURI: "/static/images/products/lawnmower_thumb.jpg", ImagePreviewURI: "/static/images/products/lawnmower.jpg", Price: 139.99, Route: "/product-detail/lawn-mower", SummaryDetail: "Nullam accumsan purus eu rutrum iaculis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec sit amet maximus turpis. Morbi dapibus orci mauris. Proin iaculis bibendum felis, id convallis ipsum pellentesque at. Cras in sem dictum, auctor tortor ut, posuere sem. Suspendisse gravida lacus a lobortis venenatis. Suspendisse ante lacus, dignissim vitae ullamcorper a, sollicitudin commodo elit. Nunc venenatis, lacus nec molestie sodales, lacus justo mollis sapien, ut tempus leo turpis non nisl. Etiam dapibus, est in ullamcorper congue, sem nunc rutrum diam, et rutrum nisi odio ac massa. Aliquam diam orci, laoreet nec lectus vel, aliquet venenatis ipsum. Proin porttitor lobortis ipsum in hendrerit. Quisque venenatis sapien eu magna laoreet ultrices. Mauris non lacinia enim. Vivamus in urna porta, aliquet ipsum id, suscipit enim. Pellentesque vulputate libero et lectus bibendum accumsan."}
		swissarmyknife := &models.Product{SKU: "6", Name: "Swiss Army Knife", Description: "A useful swiss army knife.", ThumbnailPreviewURI: "/static/images/products/swissarmyknife_thumb.jpg", ImagePreviewURI: "/static/images/products/swissarmyknife.jpg", Price: 9.99, Route: "/product-detail/swiss-army-knife", SummaryDetail: "Maecenas tincidunt orci erat, eu scelerisque nisl fermentum vitae. Pellentesque luctus augue egestas, malesuada mauris quis, blandit mi. Fusce pharetra, odio a tincidunt tincidunt, tortor elit blandit ligula, sit amet commodo sem felis nec quam. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Nam sed maximus erat. Nam in vehicula sem. Donec posuere vehicula facilisis. Proin lacinia hendrerit urna. Aenean tincidunt orci quis urna viverra, nec sagittis libero pharetra. Aenean id tempor justo, ac efficitur arcu. Fusce urna ipsum, euismod ac lorem sed, interdum pellentesque lectus. Suspendisse mattis neque eget laoreet tincidunt. Proin arcu mauris, lobortis a accumsan sed, mattis id leo."}
		tieset := &models.Product{SKU: "7", Name: "Tie Set", Description: "A classy tie set.", ThumbnailPreviewURI: "/static/images/products/tieset_thumb.jpg", ImagePreviewURI: "/static/images/products/tieset.jpg", Price: 22.99, Route: "/product-detail/tie-set", SummaryDetail: "Donec pulvinar tincidunt semper. Sed ac leo dolor. Mauris elementum posuere enim. Aliquam laoreet auctor nisl a tempus. Integer et erat sit amet nisi ullamcorper dapibus. Maecenas ut est turpis. Morbi velit urna, maximus sed bibendum ut, pretium luctus metus. Cras quis sollicitudin urna, a scelerisque felis. Mauris mollis mattis odio mollis facilisis. Ut at mi eu leo maximus auctor nec vel arcu."}
		watch := &models.Product{SKU: "8", Name: "Watch", Description: "An elegant watch.", ThumbnailPreviewURI: "/static/images/products/watch_thumb.jpg", ImagePreviewURI: "/static/images/products/watch.jpg", Price: 135.99, Route: "/product-detail/watch", SummaryDetail: "Maecenas eu velit sagittis, tristique diam sit amet, convallis massa. Aenean feugiat fringilla efficitur. Cras nec accumsan est. Phasellus vel tortor suscipit, pellentesque velit a, maximus odio. Suspendisse hendrerit cursus felis sit amet pretium. Suspendisse euismod porttitor mi, egestas porta sapien interdum quis. Phasellus at urna ac sapien lobortis sagittis. Nam convallis quam id placerat rhoncus. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Ut cursus pretium urna, in tincidunt tortor gravida non. Suspendisse porta mi in lectus elementum consequat. Nunc ut vulputate sem. Nam bibendum finibus metus sed tempus. Donec euismod metus eu quam congue gravida."}
		shirt := &models.Product{SKU: "9", Name: "Shirt", Description: "A cheerful shirt.", ThumbnailPreviewURI: "/static/images/products/shirt_thumb.jpg", ImagePreviewURI: "/static/images/products/shirt.jpg", Price: 17.99, Route: "/product-detail/shirt", SummaryDetail: "Quisque augue urna, scelerisque in maximus non, faucibus sit amet magna. Interdum et malesuada fames ac ante ipsum primis in faucibus. Aliquam quis urna purus. Integer in malesuada justo. Morbi non mi eu elit facilisis gravida. Sed dictum dui ac lectus pharetra, non viverra neque aliquet. Morbi risus risus, ullamcorper in quam pulvinar, sagittis venenatis dolor. Vestibulum dolor urna, condimentum et consequat sit amet, bibendum sed sapien. Quisque efficitur libero id odio feugiat commodo. Nam viverra vulputate ipsum ut sollicitudin. In ultrices lorem eget nisl rutrum, quis mollis neque dictum. Nam sodales bibendum metus, sed porttitor nisi laoreet at."}

		var productKeys []string = make([]string, 0)

		products := []*models.Product{chair, coffeemug, desklamp, flowervase, lawnmower, swissarmyknife, tieset, watch, shirt}

		for i := 0; i < len(products); i++ {
			err := env.DB.CreateProduct(products[i])
			if err != nil {
				log.Println("Failed to insert record for ", products[i].Route, " and encountered the following error: ", err)
				w.Write([]byte("Failed to load sample data!"))
				return
			}
			productKeys = append(productKeys, products[i].Route)
		}

		err = env.DB.CreateProductRegistry(productKeys)
		if err != nil {
			log.Println("Encountered an error when attempting to create the product registry: ", err)
			w.Write([]byte("Failed to load sample data!"))
			return
		}

		w.Write([]byte("Sample data loaded."))

	})
}
