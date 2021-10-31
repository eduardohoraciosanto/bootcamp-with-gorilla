package controller

import (
	"net/http"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/health"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"
)

type HealthController struct {
	Service health.Service
}

//Health is the handler for the health endpoint
func (c *HealthController) Health(w http.ResponseWriter, r *http.Request) {

	//using lower level pkg to do the logic
	service, external, db, err := c.Service.HealthCheck()
	if err != nil {
		viewmodels.RespondWithError(w, viewmodels.StandardInternalServerError)
		return
	}
	hr := viewmodels.HealthResponse{
		Services: []viewmodels.Health{
			{
				Name:  "service",
				Alive: service,
			},
			{
				Name:  "external api",
				Alive: external,
			},
			{
				Name:  "Cache",
				Alive: db,
			},
		},
	}
	viewmodels.RespondWithData(w, http.StatusOK, hr)
}
