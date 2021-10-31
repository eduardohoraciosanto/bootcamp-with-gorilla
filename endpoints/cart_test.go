package endpoints_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/endpoints"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/viewmodels"
)

func TestCreateCart(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.CreateCart(&ms)(context.TODO(), nil)

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Cart)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestGetCart(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.GetCart(&ms)(context.TODO(), viewmodels.GetCartRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Cart)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestGetAvailableItems(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.GetAvailableItems(&ms)(context.TODO(), nil)

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.([]models.Item)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestGetItem(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.GetItem(&ms)(context.TODO(), viewmodels.GetItemRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Item)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestAddItemToCart(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.AddItemToCart(&ms)(context.TODO(), viewmodels.AddItemToCartRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Cart)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestModifyItemQuantity(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.ModifyItemQuantity(&ms)(context.TODO(), viewmodels.ModifyItemQuantityRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Cart)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestDeleteItem(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.DeleteItem(&ms)(context.TODO(), viewmodels.DeleteItemRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Cart)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestDeleteAllItems(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.DeleteAllItems(&ms)(context.TODO(), viewmodels.DeleteAllItemsRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.(models.Cart)
	if !ok {
		t.Fatalf("response should be a models.Cart")
	}
}

func TestDeleteCart(t *testing.T) {
	ms := mockedService{shouldFail: false}
	_, err := endpoints.DeleteCart(&ms)(context.TODO(), viewmodels.DeleteCartRequest{})

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
}

//######Mocked Service

type mockedService struct {
	shouldFail bool
}

func (m *mockedService) Health(ctx context.Context) []models.Health {
	alive := true
	if m.shouldFail {
		alive = false
	}

	return []models.Health{
		{
			Name:  "Mocked Service",
			Alive: alive,
		},
	}
}
func (m *mockedService) CreateCart(ctx context.Context) (models.Cart, error) {
	if m.shouldFail {
		return models.Cart{}, fmt.Errorf("mocked service asked to fail")
	}

	return models.Cart{
		ID:    "mock-id",
		Items: []models.Item{},
	}, nil
}
func (m *mockedService) GetCart(ctx context.Context, cartID string) (models.Cart, error) {
	if m.shouldFail {
		return models.Cart{}, fmt.Errorf("mocked service asked to fail")
	}

	return models.Cart{
		ID: "mock-id",
		Items: []models.Item{
			{
				ID:       "mocked-item",
				Name:     "Mocked Item",
				Quantity: 2,
				Price:    43.21,
			},
		},
	}, nil
}
func (m *mockedService) GetAvailableItems(ctx context.Context) ([]models.Item, error) {
	if m.shouldFail {
		return nil, fmt.Errorf("mocked service was asked to fail")
	}
	return []models.Item{
		{
			ID:       "mocked-item",
			Name:     "Mocked Item",
			Quantity: 2,
			Price:    43.21,
		},
	}, nil
}
func (m *mockedService) GetItem(ctx context.Context, id string) (models.Item, error) {
	if m.shouldFail {
		return models.Item{}, fmt.Errorf("mocked service was asked to fail")
	}
	return models.Item{
		ID:       "mocked-item",
		Name:     "Mocked Item",
		Quantity: 2,
		Price:    43.21,
	}, nil
}
func (m *mockedService) AddItemToCart(ctx context.Context, cartID, itemID string, quantity int) (models.Cart, error) {
	if m.shouldFail {
		return models.Cart{}, fmt.Errorf("mocked service asked to fail")
	}

	return models.Cart{
		ID: "mock-id",
		Items: []models.Item{
			{
				ID:       "mocked-item",
				Name:     "Mocked Item",
				Quantity: 2,
				Price:    43.21,
			},
		},
	}, nil
}
func (m *mockedService) ModifyItemInCart(ctx context.Context, cartID, itemID string, newQuantity int) (models.Cart, error) {
	if m.shouldFail {
		return models.Cart{}, fmt.Errorf("mocked service asked to fail")
	}

	return models.Cart{
		ID: "mock-id",
		Items: []models.Item{
			{
				ID:       "mocked-item",
				Name:     "Mocked Item",
				Quantity: 2,
				Price:    43.21,
			},
		},
	}, nil
}
func (m *mockedService) DeleteItemInCart(ctx context.Context, cartID, itemID string) (models.Cart, error) {
	if m.shouldFail {
		return models.Cart{}, fmt.Errorf("mocked service asked to fail")
	}

	return models.Cart{
		ID: "mock-id",
		Items: []models.Item{
			{
				ID:       "mocked-item",
				Name:     "Mocked Item",
				Quantity: 2,
				Price:    43.21,
			},
		},
	}, nil
}
func (m *mockedService) DeleteAllItemsInCart(ctx context.Context, cartID string) (models.Cart, error) {
	if m.shouldFail {
		return models.Cart{}, fmt.Errorf("mocked service asked to fail")
	}

	return models.Cart{
		ID:    "mock-id",
		Items: []models.Item{},
	}, nil
}
func (m *mockedService) DeleteCart(ctx context.Context, cartID string) error {
	if m.shouldFail {
		return fmt.Errorf("mocked service asked to fail")
	}

	return nil
}
