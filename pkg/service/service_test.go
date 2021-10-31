package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/service"
)

func TestCreateCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.CreateCart(context.TODO())

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestCreateCartCacheFail(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldSetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.CreateCart(context.TODO())

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestGetCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.GetCart(context.TODO(), "testCartID")

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestGetCartCacheFail(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldGetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.GetCart(context.TODO(), "testCartID")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestGetCartExternalFail(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: true,
		})

	_, err := svc.GetCart(context.TODO(), "testCartID")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestGetAvailableItemsOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.GetAvailableItems(context.TODO())

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestGetAvailableItemsExternalFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: true,
		})

	_, err := svc.GetAvailableItems(context.TODO())

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestGetItemOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.GetItem(context.TODO(), "someItem")

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestGetItemExternalFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: true,
		})

	_, err := svc.GetItem(context.TODO(), "someItem")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestAddItemToCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.AddItemToCart(context.TODO(), "someCart", "someItem", 1)

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestAddItemToCartFailItemAlreadyAdded(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.AddItemToCart(context.TODO(), "someCart", "1-simple-Item", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestAddItemToCartCacheFailureGet(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldGetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.AddItemToCart(context.TODO(), "someCart", "someItem", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestAddItemToCartCacheFailureSet(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldSetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.AddItemToCart(context.TODO(), "someCart", "someItem", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestAddItemToCartExternalFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: true,
		})

	_, err := svc.AddItemToCart(context.TODO(), "someCart", "someItem", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestModifyItemInCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.ModifyItemInCart(context.TODO(), "someCart", "1-simple-Item", 1)

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestModifyItemInCartItemNotFound(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.ModifyItemInCart(context.TODO(), "someCart", "SomeItem", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestModifyItemInCartCacheGetFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldGetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.ModifyItemInCart(context.TODO(), "someCart", "1-simple-Item", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestModifyItemInCartCacheSetFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldSetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.ModifyItemInCart(context.TODO(), "someCart", "1-simple-Item", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestModifyItemInCartExternalFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: true,
		})

	_, err := svc.ModifyItemInCart(context.TODO(), "someCart", "1-simple-Item", 1)

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteItemInCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteItemInCart(context.TODO(), "someCart", "1-simple-Item")

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestDeleteItemInCartItemNotFound(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteItemInCart(context.TODO(), "someCart", "SomeItem")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteItemInCartCacheGetFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldGetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteItemInCart(context.TODO(), "someCart", "1-simple-Item")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteItemInCartCacheSetFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldSetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteItemInCart(context.TODO(), "someCart", "1-simple-Item")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteItemInCartExternalFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: true,
		})

	_, err := svc.DeleteItemInCart(context.TODO(), "someCart", "1-simple-Item")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteAllItemsInCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteAllItemsInCart(context.TODO(), "someCart")

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestDeleteAllItemsInCartCacheGetFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldGetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteAllItemsInCart(context.TODO(), "someCart")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteAllItemsInCartCacheSetFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldSetFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	_, err := svc.DeleteAllItemsInCart(context.TODO(), "someCart")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

func TestDeleteCartOK(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{},
		&externalMock{
			shouldFail: false,
		})

	err := svc.DeleteCart(context.TODO(), "someCart")

	if err != nil {
		t.Fatalf("Service not Expected to fail")
	}
}

func TestDeleteCartCacheFailure(t *testing.T) {
	svc := service.NewCartService("unit-testing",
		&cacheMock{
			shouldDelFail: true,
		},
		&externalMock{
			shouldFail: false,
		})

	err := svc.DeleteCart(context.TODO(), "someCart")

	if err == nil {
		t.Fatalf("Service Expected to fail")
	}
}

//*************************Mocks********************

//******** Cache Mock

type cacheMock struct {
	shouldSetFail   bool
	shouldGetFail   bool
	shouldDelFail   bool
	shouldAliveFail bool
}

func (c *cacheMock) Set(key string, value interface{}) error {
	if c.shouldSetFail {
		return fmt.Errorf("Mock was asked to fail")
	}
	return nil
}
func (c *cacheMock) Get(key string, here interface{}) error {
	if c.shouldGetFail {
		return fmt.Errorf("Mock was asked to fail")
	}
	m := here.(*models.Cart)

	m.Items = []models.Item{
		{
			ID: "1-simple-Item",
		},
		{
			ID: "2-simple-Item",
		},
	}
	return nil
}
func (c *cacheMock) Del(key string) error {
	if c.shouldDelFail {
		return fmt.Errorf("Mock was asked to fail")
	}

	return nil
}
func (c *cacheMock) Alive() bool {
	return !c.shouldAliveFail
}

//External Service Mock
type externalMock struct {
	shouldFail bool
}

func (e *externalMock) Health() error {
	if e.shouldFail {
		return fmt.Errorf("External API Mock was asked to fail")
	}
	return nil
}

func (e *externalMock) GetItem(id string) (models.Item, error) {
	if e.shouldFail {
		return models.Item{}, fmt.Errorf("External Mock was asked to Fail")
	}
	return models.Item{}, nil
}
func (e *externalMock) GetAllItems() ([]models.Item, error) {
	if e.shouldFail {
		return []models.Item{}, fmt.Errorf("External Mock was asked to Fail")
	}

	return []models.Item{}, nil
}
