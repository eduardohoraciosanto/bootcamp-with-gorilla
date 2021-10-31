package errors_test

import (
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/errors"
)

func TestErrorCode(t *testing.T) {
	err := errors.ServiceError{
		Code: "TestingError",
	}

	if err.Error() != "TestingError" {
		t.Fatalf("Error code unexpected")
	}
}
