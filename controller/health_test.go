package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/controller"
)

func TestHealthOk(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.HealthController{
		Service: &healthMock{
			shouldServiceFail:  false,
			shouldCacheFail:    false,
			shouldExternalFail: false,
			shouldReturnError:  false,
		},
	}
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	c.Health(r, req)

	if r.Result().StatusCode != http.StatusOK {
		t.Fatalf("Unexpected Status Code")
	}
}
func TestHealthError(t *testing.T) {
	r := httptest.NewRecorder()
	c := controller.HealthController{
		Service: &healthMock{
			shouldServiceFail:  false,
			shouldCacheFail:    false,
			shouldExternalFail: false,
			shouldReturnError:  true,
		},
	}
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	c.Health(r, req)

	if r.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Unexpected Status Code")
	}
}

//******** Health Service Mock

type healthMock struct {
	shouldServiceFail  bool
	shouldCacheFail    bool
	shouldExternalFail bool
	shouldReturnError  bool
}

func (hm *healthMock) HealthCheck() (service bool, externalAPI bool, cache bool, err error) {
	if hm.shouldReturnError {
		return hm.shouldServiceFail, hm.shouldExternalFail, hm.shouldCacheFail, fmt.Errorf("Health Mock was asked to fail")
	}
	return hm.shouldServiceFail, hm.shouldExternalFail, hm.shouldCacheFail, nil
}
