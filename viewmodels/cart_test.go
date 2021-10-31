package viewmodels_test

import (
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"
)

func TestCartToViewmodel(t *testing.T) {
	c := models.Cart{
		ID: "someCart",
		Items: []models.Item{
			{
				ID:       "someItem",
				Name:     "Some Item",
				Quantity: 2,
				Price:    12.34,
			},
			{
				ID:       "someItem2",
				Name:     "Some Item 2",
				Quantity: 4,
				Price:    24.68,
			},
		},
	}
	cVM := viewmodels.CartModelToViewmodel(c)

	if cVM.ID != c.ID {
		t.Fatalf("Cart ID incorrectly converted")
	}

	if cVM.Items[0].Name != c.Items[0].Name {
		t.Fatalf("Cart Item 0 Name converted incorrectly")
	}

	if cVM.Items[1].Price != c.Items[1].Price {
		t.Fatalf("Cart Item 1 Price converted incorrectly")
	}
}
