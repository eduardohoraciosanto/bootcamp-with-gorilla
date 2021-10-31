package viewmodels_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/viewmodels"
)

func TestDecodeHealthRequestOK(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	req, err := viewmodels.DecodeHealthRequest(context.TODO(), r)
	_, ok := req.(viewmodels.HealthRequest)
	if !ok {
		t.Fatalf("Unexpected type found")
	}
	if err != nil {
		t.Fatalf("Not expected to fail")
	}
}

func TestEncodeHealthResponseOK(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeHealthResponse(context.TODO(), r, []models.Health{
		{
			Name:  "test",
			Alive: true,
		},
	})
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}

func TestEncodeHealthResponseBadResponse(t *testing.T) {
	r := httptest.NewRecorder()
	err := viewmodels.EncodeHealthResponse(context.TODO(), r, "stringResponse")
	if err != nil {
		t.Fatalf("error expected to be nil")
	}
	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}
