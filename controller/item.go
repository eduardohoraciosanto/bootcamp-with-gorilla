package controller

import (
	"net/http"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/service"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"
	"github.com/gorilla/mux"
)

type ItemController struct {
	Service service.CartService
}

//GetAllItems returns all items from the external API
func (c *ItemController) GetAllItems(w http.ResponseWriter, r *http.Request) {

	items, err := c.Service.GetAvailableItems(r.Context())
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}
	vmItems := []viewmodels.Item{}
	for _, item := range items {
		vmItem := viewmodels.Item{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		}
		vmItems = append(vmItems, vmItem)
	}

	viewmodels.RespondWithData(w, http.StatusOK, vmItems)
}

//GetItem returns a particular item from the external API
func (c *ItemController) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["item_id"]
	item, err := c.Service.GetItem(r.Context(), itemID)
	if err != nil {
		viewmodels.RespondWithError(w, err)
		return
	}
	vmItem := viewmodels.Item{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}

	viewmodels.RespondWithData(w, http.StatusOK, vmItem)
}
