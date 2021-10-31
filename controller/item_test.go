package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/controller"
)

func TestGetAllItemsOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.ItemController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.GetAllItems(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestGetAllItemsError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.ItemController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.GetAllItems(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestGetItemOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.ItemController{
		Service: &mockService{
			shouldFail: false,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.GetItem(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestGetItemError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.ItemController{
		Service: &mockService{
			shouldFail: true,
		},
	}
	req, _ := http.NewRequest(http.MethodPost, "", nil)
	c.GetItem(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}
