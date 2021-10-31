package health

import (
	"fmt"
	"testing"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
)

func TestHealthCheck(t *testing.T) {
	service := NewService(
		&cacheMocked{cacheShouldFail: false},
		&externalAPIMocked{externalAPIShouldFail: false},
	)

	s, e, d, err := service.HealthCheck()
	if s != true || e != true || d != true || err != nil {
		t.Errorf("Unexpected values from method: service %t, external %t, db %t, error %s", s, e, d, err)
	}
}

func TestHealthCheck_CacheFail(t *testing.T) {
	service := NewService(
		&cacheMocked{cacheShouldFail: true},
		&externalAPIMocked{externalAPIShouldFail: false},
	)

	s, e, d, err := service.HealthCheck()
	if s != true || e != true || d != false || err != nil {
		t.Errorf("Unexpected values from method: service %t, external %t, db %t, error %s", s, e, d, err)
	}
}

func TestHealthCheck_ExternalFail(t *testing.T) {
	service := NewService(
		&cacheMocked{cacheShouldFail: false},
		&externalAPIMocked{externalAPIShouldFail: true},
	)

	s, e, d, err := service.HealthCheck()
	if s != true || e != false || d != true || err != nil {
		t.Errorf("Unexpected values from method: service %t, external %t, db %t, error %s", s, e, d, err)
	}
}

//Cache Mocked

type cacheMocked struct {
	cacheShouldFail bool
}

func (c *cacheMocked) Set(key string, value interface{}) error {
	if c.cacheShouldFail {
		return fmt.Errorf("Mock Cache Asked to Fail")
	}
	return nil
}
func (c *cacheMocked) Get(key string, here interface{}) error {
	if c.cacheShouldFail {
		return fmt.Errorf("Mock Cache Asked to Fail")
	}
	return nil
}
func (c *cacheMocked) Del(key string) error {
	if c.cacheShouldFail {
		return fmt.Errorf("Mock Cache Asked to Fail")
	}
	return nil
}
func (c *cacheMocked) Alive() bool {
	if c.cacheShouldFail {
		return false
	}
	return true
}

//External API Mocked

type externalAPIMocked struct {
	externalAPIShouldFail bool
}

func (e *externalAPIMocked) Health() error {
	if e.externalAPIShouldFail {
		return fmt.Errorf("External API Mock was asked to fail")
	}
	return nil
}

func (e *externalAPIMocked) GetItem(id string) (models.Item, error) {
	if e.externalAPIShouldFail {
		return models.Item{}, fmt.Errorf("External API Mock was asked to fail")
	}
	return models.Item{
		ID:    "mockedItem",
		Name:  "Mocked Item",
		Price: 999.99,
	}, nil
}
func (e *externalAPIMocked) GetAllItems() ([]models.Item, error) {
	if e.externalAPIShouldFail {
		return []models.Item{}, fmt.Errorf("External API Mock was asked to fail")
	}
	return []models.Item{
		{
			ID:    "mockedItem1",
			Name:  "Mocked Item 1",
			Price: 999.99,
		},
		{
			ID:    "mockedItem2",
			Name:  "Mocked Item 2",
			Price: 999.99,
		},
	}, nil
}
