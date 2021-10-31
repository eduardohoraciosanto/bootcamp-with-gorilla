package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/service"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"
	"github.com/gorilla/mux"
)

type CartController struct {
	Service service.CartService
}

//CreateCart creates a cart on the DB
func (c *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {

	cart, err := c.Service.CreateCart(r.Context())
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}

	response := viewmodels.CartResponse{
		Cart: viewmodels.CartModelToViewmodel(cart),
	}
	viewmodels.RespondWithData(w, http.StatusOK, response)
}

//GetCart creates a cart on the DB
func (c *CartController) GetCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cart_id"]
	cart, err := c.Service.GetCart(r.Context(), cartID)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}
	response := viewmodels.CartResponse{
		Cart: viewmodels.CartModelToViewmodel(cart),
	}
	viewmodels.RespondWithData(w, http.StatusOK, response)
}

//DeleteCart removes all items from the cart
func (c *CartController) DeleteCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cart_id"]

	err := c.Service.DeleteCart(r.Context(), cartID)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}
	viewmodels.RespondWithData(w, http.StatusAccepted, nil)
}

//AddItem Adds an item to a cart
func (c *CartController) AddItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cart_id"]

	vm := viewmodels.AddItemToCartRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&vm)
	if err != nil {
		log.Printf("Error decoding body: %v", err)
		viewmodels.RespondWithError(w, viewmodels.StandardBadBodyRequest)
		return
	}
	cart, err := c.Service.AddItemToCart(r.Context(), cartID, vm.ID, vm.Quantity)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}

	response := viewmodels.CartResponse{
		Cart: viewmodels.CartModelToViewmodel(cart),
	}
	viewmodels.RespondWithData(w, http.StatusOK, response)
}

//UpdateQuantity changes the amount of a single item in the cart
func (c *CartController) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cart_id"]
	itemID := vars["item_id"]

	vm := viewmodels.ModifyItemQuantityRequest{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&vm)
	if err != nil {
		log.Printf("Error decoding body: %v", err)
		viewmodels.RespondWithError(w, viewmodels.StandardBadBodyRequest)
		return
	}

	cart, err := c.Service.ModifyItemInCart(r.Context(), cartID, itemID, vm.Quantity)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}

	response := viewmodels.CartResponse{
		Cart: viewmodels.CartModelToViewmodel(cart),
	}
	viewmodels.RespondWithData(w, http.StatusOK, response)
}

//RemoveItem removes an item from the cart
func (c *CartController) RemoveItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cart_id"]
	itemID := vars["item_id"]

	cart, err := c.Service.DeleteItemInCart(r.Context(), cartID, itemID)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}

	response := viewmodels.CartResponse{
		Cart: viewmodels.CartModelToViewmodel(cart),
	}

	viewmodels.RespondWithData(w, http.StatusOK, response)
}

//RemoveAllItems removes all items from the cart
func (c *CartController) RemoveAllItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["cart_id"]

	cart, err := c.Service.DeleteAllItemsInCart(r.Context(), cartID)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}

	response := viewmodels.CartResponse{
		Cart: viewmodels.CartModelToViewmodel(cart),
	}
	viewmodels.RespondWithData(w, http.StatusOK, response)
}
