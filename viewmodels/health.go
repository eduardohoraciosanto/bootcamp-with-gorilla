package viewmodels

import (
	"context"
	"net/http"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
)

type Health struct {
	Name  string `json:"name"`
	Alive bool   `json:"alive"`
}

type HealthRequest struct {
}

func DecodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := HealthRequest{}
	return req, nil
}

type HealthResponse struct {
	Services []Health `json:"services"`
}

func EncodeHealthResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	epRes, ok := response.([]models.Health)
	if !ok {
		return RespondWithError(w, StandardInternalServerError)
	}
	vmHealths := []Health{}

	for _, health := range epRes {
		vmHealths = append(vmHealths, Health{
			Name:  health.Name,
			Alive: health.Alive,
		})
	}

	vmResponse := HealthResponse{
		Services: vmHealths,
	}

	return RespondWithData(w, http.StatusOK, vmResponse)
}
