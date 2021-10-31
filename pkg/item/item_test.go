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

func TestGetItem(t *testing.T) {

	svc := item.NewExternalService(
		logrus.New(),
		&itemClientMock{
			shouldFail: false,
			response: viewmodels.ExternalItem{
				ID:    "someItemID",
				Title: "Some Item ID",
				Price: "12.34",
			},
		},
	)

	_, err := svc.GetItem("someItemID")
	if err != nil {
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
			response: viewmodels.ExternalItem{
				ID:    "someItemID",
				Title: "Some Item ID",
				Price: "ThisisNotANumber",
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
			response: []viewmodels.ExternalItem{
				{
					ID:    "someItemID",
					Title: "Some Item ID",
					Price: "12.34",
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
			response: []viewmodels.ExternalItem{
				{
					ID:    "someItemID",
					Title: "Some Item ID",
					Price: "ThisisNotANumber",
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
	response   interface{}
	shouldFail bool
}

func (i *itemClientMock) Get(url string) (*http.Response, error) {
	if i.shouldFail {
		return nil, fmt.Errorf("Mock asked to fail")
	}
	b, _ := json.Marshal(i.response)
	resp := &http.Response{}
	resp.Body = ioutil.NopCloser(bytes.NewReader(b))
	return resp, nil
}
