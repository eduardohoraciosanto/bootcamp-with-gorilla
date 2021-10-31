package viewmodels_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	serviceErrors "github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/errors"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"
)

func TestRespondWithData(t *testing.T) {
	r := httptest.NewRecorder()
	data := viewmodels.Cart{
		ID: "testCart",
	}
	viewmodels.RespondWithData(r, http.StatusOK, data)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrNotFound(t *testing.T) {
	r := httptest.NewRecorder()
	mErr := &serviceErrors.ServiceError{
		Code: serviceErrors.CartNotFoundCode,
	}
	viewmodels.RespondWithError(r, mErr)

	if r.Result().StatusCode != http.StatusNotFound {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrItemNotFound(t *testing.T) {
	r := httptest.NewRecorder()
	mErr := &serviceErrors.ServiceError{
		Code: serviceErrors.ItemNotFoundCode,
	}
	viewmodels.RespondWithError(r, mErr)

	if r.Result().StatusCode != http.StatusNotFound {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrUnprocessableEntity(t *testing.T) {
	r := httptest.NewRecorder()
	mErr := &serviceErrors.ServiceError{
		Code: serviceErrors.ItemAlreadyInCartCode,
	}
	viewmodels.RespondWithError(r, mErr)

	if r.Result().StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrInternal(t *testing.T) {
	r := httptest.NewRecorder()
	mErr := &serviceErrors.ServiceError{
		Code: "some Error Code",
	}
	viewmodels.RespondWithError(r, mErr)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrInternalDefault(t *testing.T) {
	r := httptest.NewRecorder()
	viewmodels.RespondWithError(r, fmt.Errorf("Some error"))

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrBadReq(t *testing.T) {
	r := httptest.NewRecorder()
	mErr := &viewmodels.Error{
		Code: viewmodels.ErrCodeBadRequest,
	}
	viewmodels.RespondWithError(r, mErr)

	if r.Result().StatusCode != http.StatusBadRequest {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestRespondWithErrViewInternalError(t *testing.T) {
	r := httptest.NewRecorder()
	mErr := &viewmodels.Error{
		Code: "SomeCode",
	}
	viewmodels.RespondWithError(r, mErr)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}
