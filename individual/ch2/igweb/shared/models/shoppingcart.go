package models

import (
	"encoding/json"
	"log"
	"strings"
)

type ShoppingCartItem struct {
	ProductSKU string `json:"productSKU"`
	Quantity   int    `json:"quantity"`
}

type ShoppingCart struct {
	Items map[string]*ShoppingCartItem `json:"items"`
}

func NewShoppingCart() *ShoppingCart {
	items := make(map[string]*ShoppingCartItem)
	return &ShoppingCart{Items: items}
}

func (s *ShoppingCart) ItemTotal() int {
	return len(s.Items)
}

func (s *ShoppingCart) IsEmpty() bool {

	if len(s.Items) > 0 {
		return false
	} else {
		return true
	}

}

func (s *ShoppingCart) AddItem(sku string) {

	if s.Items == nil {
		s.Items = make(map[string]*ShoppingCartItem)
	}

	_, ok := s.Items[sku]
	if ok {
		s.Items[sku].Quantity += 1

	} else {
		item := ShoppingCartItem{ProductSKU: sku, Quantity: 1}
		s.Items[sku] = &item
	}

}

func (s *ShoppingCart) RemoveItem(sku string) bool {

	_, ok := s.Items[sku]
	if ok {
		delete(s.Items, sku)
		return true
	} else {
		return false
	}

}

func (s *ShoppingCart) UpdateItemQuantity(sku string, quantity int) bool {

	//, item *ShoppingCartItem

	_, ok := s.Items[sku]
	if ok {
		s.Items[sku].Quantity += 1
		return true
	} else {

		return false
	}

}

func SerializeShoppingCart(s *ShoppingCart) []byte {

	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Print("Encountered error when attempting to marshal json: ", err)
	}

	return jsonData
}

func UnserializeShoppingCart(jsonData []byte) *ShoppingCart {

	var s *ShoppingCart = &ShoppingCart{}
	json.NewDecoder(strings.NewReader(string(jsonData))).Decode(s)
	return s
}
