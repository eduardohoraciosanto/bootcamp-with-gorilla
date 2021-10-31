package viewmodels_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/viewmodels"
	"github.com/gorilla/mux"
)

func TestDecodeCreateCartRequestOK(t *testing.T) {
	r := &http.Request{}
	req, err := viewmodels.DecodeCreateCartRequest(context.TODO(), r)
	_, ok := req.(viewmodels.CreateCartRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestEncodeCreateCartResponseOK(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeCartResponse(context.TODO(), r, models.Cart{
		ID: "someCartID",
		Items: []models.Item{
			{
				ID:       "someItemID",
				Name:     "Some Item",
				Quantity: 2,
				Price:    32.12,
			},
		},
	})
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestEncodeCreateCartResponseBadResponse(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeCartResponse(context.TODO(), r, "badResponse")
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestDecodeAddItemToCartRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap,
		viewmodels.AddItemToCartRequest{
			ID:       "someItemID",
			Quantity: 3,
		},
	)

	req, err := viewmodels.DecodeAddItemToCartRequest(context.TODO(), r)
	_, ok := req.(viewmodels.AddItemToCartRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDecodeAddItemToCartRequestIDMissing(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap,
		viewmodels.AddItemToCartRequest{
			ID:       "someItemID",
			Quantity: 3,
		},
	)

	_, err := viewmodels.DecodeAddItemToCartRequest(context.TODO(), r)

	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDecodeAddItemToCartRequestBadRequest(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap,
		"BadRequest",
	)

	_, err := viewmodels.DecodeAddItemToCartRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDecodeGetCartRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap,
		viewmodels.GetCartRequest{},
	)

	req, err := viewmodels.DecodeGetCartRequest(context.TODO(), r)
	_, ok := req.(viewmodels.GetCartRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDecodeGetCartRequestMissingCartID(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap,
		viewmodels.GetCartRequest{},
	)

	_, err := viewmodels.DecodeGetCartRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDecodeModifyItemQuantityRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"
	muxVarsMap["item_id"] = "SomeItemID"

	r := newRequest(
		muxVarsMap,
		viewmodels.ModifyItemQuantityRequest{},
	)

	req, err := viewmodels.DecodeModifyItemQuantityRequest(context.TODO(), r)
	_, ok := req.(viewmodels.ModifyItemQuantityRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDecodeModifyItemQuantityRequestMissingCartID(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap,
		viewmodels.ModifyItemQuantityRequest{},
	)

	_, err := viewmodels.DecodeModifyItemQuantityRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDecodeModifyItemQuantityRequestMissingItemID(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap,
		viewmodels.ModifyItemQuantityRequest{},
	)

	_, err := viewmodels.DecodeModifyItemQuantityRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDecodeModifyItemQuantityRequestBadRequest(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"
	muxVarsMap["item_id"] = "SomeItemID"

	r := newRequest(
		muxVarsMap,
		"BadRequest",
	)

	_, err := viewmodels.DecodeModifyItemQuantityRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDeleteItemRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"
	muxVarsMap["item_id"] = "SomeItemID"

	r := newRequest(
		muxVarsMap, nil,
	)

	req, err := viewmodels.DecodeDeleteItemRequest(context.TODO(), r)
	_, ok := req.(viewmodels.DeleteItemRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDeleteItemRequestMissingCartID(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap, nil,
	)

	_, err := viewmodels.DecodeDeleteItemRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDeleteItemRequestMissingItemID(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap, nil,
	)

	_, err := viewmodels.DecodeDeleteItemRequest(context.TODO(), r)

	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDeleteAllItemsRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap, nil,
	)

	req, err := viewmodels.DecodeDeleteAllItemsRequest(context.TODO(), r)
	_, ok := req.(viewmodels.DeleteAllItemsRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDeleteAllItemsRequestMissingCartID(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap, nil,
	)

	_, err := viewmodels.DecodeDeleteAllItemsRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestDeleteCartRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["cart_id"] = "SomeCartID"

	r := newRequest(
		muxVarsMap, nil,
	)

	req, err := viewmodels.DecodeDeleteCartRequest(context.TODO(), r)
	_, ok := req.(viewmodels.DeleteCartRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDeleteCartRequestMissingCartID(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap, nil,
	)

	_, err := viewmodels.DecodeDeleteCartRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestEncodeDeleteCartResponseOK(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeDeleteCartResponse(context.TODO(), r, nil)
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusAccepted {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestDecodeGetAvailableItemsRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)

	r := newRequest(
		muxVarsMap,
		nil,
	)

	req, err := viewmodels.DecodeGetAvailableItemsRequest(context.TODO(), r)
	_, ok := req.(viewmodels.GetAvailableItemsRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestEncodeGetAvailableItemsRequestOK(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeGetAvailableItemsResponse(context.TODO(), r, []models.Item{
		{
			ID:    "SomeItemID",
			Name:  "Some Item Name",
			Price: 12.34,
		},
		{
			ID:    "SomeOtherItemID",
			Name:  "Some Other Item Name",
			Price: 54.67,
		},
	})
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestDecodeGetItemRequestOK(t *testing.T) {
	muxVarsMap := make(map[string]string)
	muxVarsMap["item_id"] = "someItemID"
	r := newRequest(
		muxVarsMap,
		nil,
	)

	req, err := viewmodels.DecodeGetItemRequest(context.TODO(), r)
	_, ok := req.(viewmodels.GetItemRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestDecodeGetItemRequestBadRequest(t *testing.T) {
	muxVarsMap := make(map[string]string)
	r := newRequest(
		muxVarsMap,
		nil,
	)

	_, err := viewmodels.DecodeGetItemRequest(context.TODO(), r)
	if err == nil {
		t.Fatalf("Expected to fail")
	}
}

func TestEncodeGetItemRequestOK(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeGetItemResponse(context.TODO(), r, models.Item{
		ID:    "SomeItemID",
		Name:  "Some Item Name",
		Price: 12.34,
	})
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}

func newRequest(vars map[string]string, body interface{}) *http.Request {
	r := &http.Request{}

	r = mux.SetURLVars(r, vars)

	b, _ := json.Marshal(body)

	r.Body = ioutil.NopCloser(bytes.NewReader(b))

	return r
}
