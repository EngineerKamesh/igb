package datastore

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/EngineerKamesh/igb/igweb/shared/models"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisDatastore struct {
	*pool.Pool
}

func NewRedisDatastore(address string) (*RedisDatastore, error) {

	connectionPool, err := pool.New("tcp", address, 10)
	if err != nil {
		return nil, err
	}
	return &RedisDatastore{
		Pool: connectionPool,
	}, nil
}

func (r *RedisDatastore) GetProducts() []*models.Product {

	registryKey := "product-registry"
	exists, err := r.Cmd("EXISTS", registryKey).Int()

	if err != nil {
		log.Println("Encountered error: ", err)
		return nil
	} else if exists == 0 {
		return nil
	}

	var productKeys []string
	jsonData, err := r.Cmd("GET", registryKey).Str()
	if err != nil {
		log.Print("Encountered error when attempting to fetch product registry data from Redis instance: ", err)
		return nil
	}

	if err := json.Unmarshal([]byte(jsonData), &productKeys); err != nil {
		log.Print("Encountered error when attempting to unmarshal JSON product registry data: ", err)
		return nil
	}

	products := make([]*models.Product, 0)

	for i := 0; i < len(productKeys); i++ {

		productTitle := strings.Replace(productKeys[i], "/product-detail/", "", -1)
		product := r.GetProductDetail(productTitle)
		products = append(products, product)

	}
	return products
}

func (r *RedisDatastore) GetProductDetail(productTitle string) *models.Product {

	productKey := "/product-detail/" + productTitle
	exists, err := r.Cmd("EXISTS", productKey).Int()

	if err != nil {
		log.Println("Encountered error: ", err)
		return nil
	} else if exists == 0 {
		return nil
	}

	var p models.Product
	jsonData, err := r.Cmd("GET", productKey).Str()

	if err != nil {
		log.Print("Encountered error when attempting to fetch product data from Redis instance: ", err)
		return nil
	}

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Print("Encountered error when attempting to unmarshal JSON product data: ", err)
		return nil
	}

	return &p

}

func (r *RedisDatastore) GenerateProductsMap(products []*models.Product) map[string]*models.Product {

	productsMap := make(map[string]*models.Product)
	for i := 0; i < len(products); i++ {
		productsMap[products[i].SKU] = products[i]
	}

	return productsMap
}

func (r *RedisDatastore) GetProductsInShoppingCart(cart *models.ShoppingCart) []*models.Product {

	products := r.GetProducts()
	productsMap := r.GenerateProductsMap(products)

	result := make([]*models.Product, 0)
	for _, v := range cart.Items {
		product := &models.Product{}
		product = productsMap[v.ProductSKU]
		product.Quantity = v.Quantity

		result = append(result, product)
	}

	return result

}

func (r *RedisDatastore) CreateProduct(product *models.Product) error {

	jsonData, err := json.Marshal(*product)
	if err != nil {
		return err
	}

	if r.Cmd("SET", product.Route, string(jsonData)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

func (r *RedisDatastore) CreateProductRegistry(products []string) error {

	jsonData, err := json.Marshal(products)
	if err != nil {
		return err
	}

	if r.Cmd("SET", "product-registry", string(jsonData)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

func (r *RedisDatastore) CreateGopherTeam(team []*models.Gopher) error {

	jsonData, err := json.Marshal(team)
	if err != nil {
		return err
	}

	if r.Cmd("SET", "gopher-team", string(jsonData)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil

}

func (r *RedisDatastore) GetGopherTeam() []*models.Gopher {

	exists, err := r.Cmd("EXISTS", "gopher-team").Int()

	if err != nil {
		log.Println("Encountered error: ", err)
		return nil
	} else if exists == 0 {
		return nil
	}

	var t []*models.Gopher
	jsonData, err := r.Cmd("GET", "gopher-team").Str()

	if err != nil {
		log.Print("Encountered error when attempting to fetch gopher team data from Redis instance: ", err)
		return nil
	}

	if err := json.Unmarshal([]byte(jsonData), &t); err != nil {
		log.Print("Encountered error when attempting to unmarshal JSON gopher team data: ", err)
		return nil
	}

	return t

}

func (r *RedisDatastore) CreateContactRequest(contactRequest *models.ContactRequest) error {

	now := time.Now()
	nowFormatted := now.Format(time.RFC822Z)

	jsonData, err := json.Marshal(contactRequest)
	if err != nil {
		return err
	}

	if r.Cmd("SET", "contact-request|"+contactRequest.Email+"|"+nowFormatted, string(jsonData)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil

}

func (r *RedisDatastore) Close() {

	r.Close()
}
