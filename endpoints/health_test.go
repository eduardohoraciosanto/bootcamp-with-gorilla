package endpoints_test

import (
	"context"
	"testing"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/endpoints"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
)

func TestHealth(t *testing.T) {
	ms := mockedService{shouldFail: false}
	response, err := endpoints.Health(&ms)(context.TODO(), nil)

	if err != nil {
		t.Fatalf("error is supposed to be nil")
	}
	_, ok := response.([]models.Health)
	if !ok {
		t.Fatalf("response should be a models.Health array")
	}
}
