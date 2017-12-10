package datastore

import (
	"errors"

	"github.com/EngineerKamesh/igb/igweb/shared/models"
)

type Datastore interface {
	CreateGopherTeam(team []*models.Gopher) error
	GetGopherTeam() []*models.Gopher
	CreateProduct(product *models.Product) error
	CreateProductRegistry(products []string) error
	GetProducts() []*models.Product
	GetProductDetail(productTitle string) *models.Product
	GetProductsInShoppingCart(cart *models.ShoppingCart) []*models.Product
	CreateContactRequest(contactRrequest *models.ContactRequest) error
	Close()
}

const (
	REDIS = iota
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {

	case REDIS:
		return NewRedisDatastore(dbConnectionString)

	default:
		return nil, errors.New("Unrecognized Datastore!")

	}
}
