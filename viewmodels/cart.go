package viewmodels

import (
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
)

type Cart struct {
	ID    string `json:"id"`
	Items []Item `json:"items"`
}

type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity,omitempty"`
	Price    float32 `json:"price"`
}

type CartResponse struct {
	Cart Cart `json:"cart"`
}

func CartModelToViewmodel(cart models.Cart) Cart {
	vmItems := []Item{}

	for _, item := range cart.Items {
		vmItems = append(vmItems, Item{
			ID:       item.ID,
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}

	return Cart{
		ID:    cart.ID,
		Items: vmItems,
	}
}

type AddItemToCartRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type ModifyItemQuantityRequest struct {
	Quantity int `json:"quantity"`
}
