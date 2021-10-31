package viewmodels

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
	"github.com/gorilla/mux"
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

type CreateCartRequest struct {
}

func DecodeCreateCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := CreateCartRequest{}
	return req, nil
}

type CartResponse struct {
	Cart Cart `json:"cart"`
}

func EncodeCartResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	epRes, ok := response.(models.Cart)
	if !ok {
		return RespondWithError(w, StandardInternalServerError)
	}
	vmItems := []Item{}

	for _, item := range epRes.Items {
		vmItems = append(vmItems, Item{
			ID:       item.ID,
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
		})
	}
	vmResponse := CartResponse{
		Cart: Cart{
			ID:    epRes.ID,
			Items: vmItems,
		},
	}
	return RespondWithData(w, http.StatusOK, vmResponse)
}

type AddItemToCartRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
	CartID   string `json:"-"`
}

func DecodeAddItemToCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := AddItemToCartRequest{}
	vars := mux.Vars(r)
	cartID, ok := vars["cart_id"]
	if !ok || cartID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestBody,
		}
	}
	req.CartID = cartID
	return req, nil
}

type GetCartRequest struct {
	CartID string
}

func DecodeGetCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := GetCartRequest{}
	vars := mux.Vars(r)
	cartID, ok := vars["cart_id"]
	if !ok || cartID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}
	req.CartID = cartID
	return req, nil
}

type ModifyItemQuantityRequest struct {
	ID       string `json:"-"`
	Quantity int    `json:"quantity"`
	CartID   string `json:"-"`
}

func DecodeModifyItemQuantityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := ModifyItemQuantityRequest{}
	vars := mux.Vars(r)
	cartID, ok := vars["cart_id"]
	if !ok || cartID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}
	itemID, ok := vars["item_id"]
	if !ok || itemID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestBody,
		}
	}
	req.CartID = cartID
	req.ID = itemID
	return req, nil
}

type DeleteItemRequest struct {
	ID     string `json:"-"`
	CartID string `json:"-"`
}

func DecodeDeleteItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := DeleteItemRequest{}
	vars := mux.Vars(r)
	cartID, ok := vars["cart_id"]
	if !ok || cartID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}
	itemID, ok := vars["item_id"]
	if !ok || itemID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}

	req.CartID = cartID
	req.ID = itemID
	return req, nil
}

type DeleteAllItemsRequest struct {
	CartID string `json:"-"`
}

func DecodeDeleteAllItemsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := DeleteAllItemsRequest{}
	vars := mux.Vars(r)
	cartID, ok := vars["cart_id"]
	if !ok || cartID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}

	req.CartID = cartID
	return req, nil
}

type DeleteCartRequest struct {
	CartID string `json:"-"`
}

func DecodeDeleteCartRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := DeleteCartRequest{}
	vars := mux.Vars(r)
	cartID, ok := vars["cart_id"]
	if !ok || cartID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}

	req.CartID = cartID
	return req, nil
}

func EncodeDeleteCartResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return RespondWithData(w, http.StatusAccepted, nil)
}

type GetAvailableItemsRequest struct {
}

func DecodeGetAvailableItemsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := GetAvailableItemsRequest{}
	return req, nil
}

type GetAvailableItemsResponse struct {
	Items []Item `json:"items"`
}

func EncodeGetAvailableItemsResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	items := response.([]models.Item)
	res := GetAvailableItemsResponse{Items: []Item{}}

	for _, item := range items {
		item := Item{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		}
		res.Items = append(res.Items, item)
	}
	return RespondWithData(w, http.StatusOK, res)
}

type GetItemRequest struct {
	ID string `json:"-"`
}

func DecodeGetItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	itemID, ok := vars["item_id"]
	if !ok || itemID == "" {
		return nil, Error{
			Code:        ErrCodeBadRequest,
			Description: ErrDescriptionBadRequestURL,
		}
	}

	return GetItemRequest{ID: itemID}, nil
}

type GetItemResponse struct {
	Item Item `json:"item"`
}

func EncodeGetItemResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	item := response.(models.Item)
	res := GetItemResponse{
		Item: Item{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		},
	}
	return RespondWithData(w, http.StatusOK, res)
}
