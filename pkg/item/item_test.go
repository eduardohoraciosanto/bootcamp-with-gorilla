package item_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/item"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"

	"github.com/sirupsen/logrus"
)

func TestHealthOK(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalHealthResponse{
				Meta: viewmodels.ExternalMeta{
					Version: "testing",
				},
				Data: viewmodels.ExternalHealth{
					Status: "OK",
				},
			},
		},
	)

	err := svc.Health()
	if err != nil {
		t.Fatalf("Error was not expected")
	}
}

func TestHealthStatusIncorrect(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalHealthResponse{
				Meta: viewmodels.ExternalMeta{
					Version: "testing",
				},
				Data: viewmodels.ExternalHealth{
					Status: "Incorrect",
				},
			},
		},
	)

	err := svc.Health()
	if err == nil {
		t.Fatalf("Error was expected")
	}
}
func TestHealthError(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: true,
			response:   nil,
		},
	)

	err := svc.Health()
	if err == nil {
		t.Fatalf("Error was expected")
	}
}
func TestHealthDecodingError(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response:   "notAJSON",
		},
	)

	err := svc.Health()
	if err == nil {
		t.Fatalf("Error was expected")
	}
}

func TestGetItem(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalGetItemResponse{
				Meta: viewmodels.ExternalMeta{
					Version: "testing",
				},
				Data: viewmodels.ExternalItem{
					ID:    "someItemID",
					Name:  "Some Item ID",
					Price: "12.34",
				},
			},
		},
	)

	_, err := svc.GetItem("someItemID")
	if err != nil {
		t.Fatalf("Error was not expected")
	}
}
func TestGetItemNotFound(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail:         false,
			response:           nil,
			responseStatusCode: 404,
		},
	)

	_, err := svc.GetItem("someItemID")
	if err == nil {
		t.Fatalf("Error was not expected")
	}
}
func TestGetItemApiFailure(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: true,
			response:   nil,
		},
	)

	_, err := svc.GetItem("someItemID")
	if err == nil {
		t.Fatalf("Error was expected")
	}
}
func TestGetItemWrongResponse(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response:   "WrongResponse",
		},
	)

	_, err := svc.GetItem("someItemID")
	if err == nil {
		t.Fatalf("Error was expected")
	}
}
func TestGetItemFloatParseError(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalGetItemResponse{
				Meta: viewmodels.ExternalMeta{
					Version: "testing",
				},
				Data: viewmodels.ExternalItem{
					ID:    "someItemID",
					Name:  "Some Item ID",
					Price: "notANumber",
				},
			},
		},
	)

	_, err := svc.GetItem("someItemID")
	if err == nil {
		t.Fatalf("Error was expected")
	}
}

func TestGetAllItems(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalGetAllItemsResponse{
				Meta: viewmodels.ExternalMeta{
					Version: "testing",
				},
				Data: []viewmodels.ExternalItem{
					{
						ID:    "someItemID",
						Name:  "Some Item ID",
						Price: "12.34",
					},
				},
			},
		},
	)

	_, err := svc.GetAllItems()
	if err != nil {
		t.Fatalf("Error was not expected")
	}
}
func TestGetAllItemsApiFailure(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: true,
			response:   nil,
		},
	)

	_, err := svc.GetAllItems()
	if err == nil {
		t.Fatalf("Error was expected")
	}
}
func TestGetAllItemsWrongResponse(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response:   "WrongResponse",
		},
	)

	_, err := svc.GetAllItems()
	if err == nil {
		t.Fatalf("Error was expected")
	}
}
func TestGetAllItemsFloatParseError(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalGetAllItemsResponse{
				Meta: viewmodels.ExternalMeta{
					Version: "testing",
				},
				Data: []viewmodels.ExternalItem{
					{
						ID:    "someItemID",
						Name:  "Some Item ID",
						Price: "notANumber",
					},
				},
			},
		},
	)

	_, err := svc.GetAllItems()
	if err == nil {
		t.Fatalf("Error was expected")
	}
}

//*****ItemClientMock

type itemClientMock struct {
	response           interface{}
	responseStatusCode int
	shouldFail         bool
}

func (i *itemClientMock) Get(url string) (*http.Response, error) {
	if i.shouldFail {
		return nil, fmt.Errorf("Mock asked to fail")
	}
	b, _ := json.Marshal(i.response)
	resp := &http.Response{}
	resp.Body = ioutil.NopCloser(bytes.NewReader(b))
	if i.responseStatusCode != 0 {
		resp.StatusCode = i.responseStatusCode
	}
	return resp, nil
}
