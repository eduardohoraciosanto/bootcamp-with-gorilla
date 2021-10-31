package health

import (
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/cache"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/item"
)

//Service is the interface for the health
type Service interface {
	HealthCheck() (service bool, externalAPI bool, cache bool, err error)
}
type svc struct {
	cache           cache.Cache
	externalService item.ExternalService
}

//NewService gives a new Service
func NewService(c cache.Cache, es item.ExternalService) Service {
	return &svc{
		cache:           c,
		externalService: es,
	}
}

//HealthCheck returns the status of the API and it's components
func (s *svc) HealthCheck() (service bool, externalAPI bool, cache bool, err error) {
	externalApiHealth := true

	exterr := s.externalService.Health()
	if exterr != nil {
		externalApiHealth = false
	}
	return true, externalApiHealth, s.cache.Alive(), nil
}
