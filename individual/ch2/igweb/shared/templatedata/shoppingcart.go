package templatedata

import "github.com/EngineerKamesh/igb/igweb/shared/models"

type ShoppingCart struct {
	PageTitle string
	Products  []*models.Product
}
