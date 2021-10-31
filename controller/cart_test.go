package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/controller"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"
)

func TestCreateCartOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.CreateCart(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestCreateCartError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.CreateCart(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

func TestGetCartOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.GetCart(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestGetCartError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.GetCart(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

func TestDeleteCartOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	c.DeleteCart(r, req)

	if r.Result().StatusCode != http.StatusAccepted {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestDeleteCartError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	c.DeleteCart(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

func TestAddItemOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	bodyBytes, _ := json.Marshal(viewmodels.AddItemToCartRequest{})
	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(bodyBytes))
	c.AddItem(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestAddItemBadRequest(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader([]byte("badBody")))
	c.AddItem(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}
func TestAddItemError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	bodyBytes, _ := json.Marshal(viewmodels.AddItemToCartRequest{})
	req, _ := http.NewRequest(http.MethodPost, "", bytes.NewReader(bodyBytes))
	c.AddItem(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

func TestUpdateQuantityOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	bodyBytes, _ := json.Marshal(viewmodels.ModifyItemQuantityRequest{})
	req, _ := http.NewRequest(http.MethodPut, "", bytes.NewReader(bodyBytes))
	c.UpdateQuantity(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestUpdateQuantityBadRequest(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodPut, "", bytes.NewReader([]byte("badBody")))
	c.UpdateQuantity(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestUpdateQuantityError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	bodyBytes, _ := json.Marshal(viewmodels.ModifyItemQuantityRequest{})
	req, _ := http.NewRequest(http.MethodPut, "", bytes.NewReader(bodyBytes))
	c.UpdateQuantity(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

func TestRemoveItemOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	c.RemoveItem(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestRemoveItemError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	c.RemoveItem(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

func TestRemoveAllItemsOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	c.RemoveAllItems(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestRemoveAllItemsError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.CartController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodDelete, "", nil)
	c.RemoveAllItems(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code: %d", r.Result().StatusCode)
	}
}

// Mock service

type mockService struct {
	shouldFail bool
}

func (ms *mockService) CreateCart(ctx context.Context) (models.Cart, error) {
	if ms.shouldFail {
		return models.Cart{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Cart{}, nil
}
func (ms *mockService) GetCart(ctx context.Context, cartID string) (models.Cart, error) {
	if ms.shouldFail {
		return models.Cart{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Cart{}, nil
}
func (ms *mockService) GetAvailableItems(ctx context.Context) ([]models.Item, error) {
	if ms.shouldFail {
		return []models.Item{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return []models.Item{
		{
			ID:    "someItem",
			Name:  "Some Item",
			Price: 12.34,
		},
	}, nil
}
func (ms *mockService) GetItem(ctx context.Context, id string) (models.Item, error) {
	if ms.shouldFail {
		return models.Item{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Item{
		ID:    "someItem",
		Name:  "Some Item",
		Price: 12.34,
	}, nil
}
func (ms *mockService) AddItemToCart(ctx context.Context, cartID, itemID string, quantity int) (models.Cart, error) {
	if ms.shouldFail {
		return models.Cart{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Cart{}, nil
}
func (ms *mockService) ModifyItemInCart(ctx context.Context, cartID, itemID string, newQuantity int) (models.Cart, error) {
	if ms.shouldFail {
		return models.Cart{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Cart{}, nil
}
func (ms *mockService) DeleteItemInCart(ctx context.Context, cartID, itemID string) (models.Cart, error) {
	if ms.shouldFail {
		return models.Cart{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Cart{}, nil
}
func (ms *mockService) DeleteAllItemsInCart(ctx context.Context, cartID string) (models.Cart, error) {
	if ms.shouldFail {
		return models.Cart{}, fmt.Errorf("Mock Service was asked to fail")
	}

	return models.Cart{}, nil
}
func (ms *mockService) DeleteCart(ctx context.Context, cartID string) error {
	if ms.shouldFail {
		return fmt.Errorf("Mock Service was asked to fail")
	}

	return nil
}
