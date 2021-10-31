package item

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/viewmodels"

	"github.com/sirupsen/logrus"
)

const (
	healthEndpoint   = "https://bootcamp-products.getsandbox.com/health"
	articlesEndpoint = "https://bootcamp-products.getsandbox.com/products"

	healthStatusOK = "OK"
)

type ExternalService interface {
	Health() error
	GetItem(id string) (models.Item, error)
	GetAllItems() ([]models.Item, error)
}

type externalService struct {
	client ItemClient
	logger *logrus.Logger
}

type ItemClient interface {
	Get(url string) (resp *http.Response, err error)
}

func NewExternalService(logger *logrus.Logger, client ItemClient) ExternalService {

	return &externalService{
		logger: logger,
		client: client,
	}
}

func (e *externalService) Health() error {
	e.logger.Log(logrus.InfoLevel, "Calling External API Health")
	res, err := e.client.Get(healthEndpoint)
	if err != nil {
		e.logger.WithError(err).Log(logrus.ErrorLevel, "Error Calling External API Health")
		return err
	}
	eHealth := viewmodels.ExternalHealthResponse{}

	err = json.NewDecoder(res.Body).Decode(&eHealth)
	if err != nil {
		e.logger.WithError(err).Log(logrus.ErrorLevel, "Error Decoding External API Health")
		return err
	}

	if eHealth.Data.Status != healthStatusOK {
		e.logger.WithField("external_api_status", eHealth.Data.Status).Log(logrus.ErrorLevel, "External API Not Healthy")
		return fmt.Errorf("external API not Healthy - Status: %s", eHealth.Data.Status)
	}
	return nil
}
func (e *externalService) GetItem(id string) (models.Item, error) {
	res, err := e.client.Get(articlesEndpoint + "/" + id)
	if err != nil {
		return models.Item{}, err
	}
	eItem := viewmodels.ExternalGetItemResponse{}

	err = json.NewDecoder(res.Body).Decode(&eItem)
	if err != nil {
		return models.Item{}, err
	}
	price, err := strconv.ParseFloat(eItem.Data.Price, 32)
	if err != nil {
		return models.Item{}, err
	}

	mItem := models.Item{
		ID:    eItem.Data.ID,
		Name:  eItem.Data.Name,
		Price: float32(price),
	}

	return mItem, nil
}
func (e *externalService) GetAllItems() ([]models.Item, error) {
	res, err := e.client.Get(articlesEndpoint)
	if err != nil {
		return []models.Item{}, err
	}
	eItems := viewmodels.ExternalGetAllItemsResponse{}

	err = json.NewDecoder(res.Body).Decode(&eItems)
	if err != nil {
		return []models.Item{}, err
	}

	mItems := []models.Item{}
	for _, eItem := range eItems.Data {
		price, err := strconv.ParseFloat(eItem.Price, 32)
		if err != nil {
			return []models.Item{}, err
		}
		mItems = append(mItems, models.Item{
			ID:    eItem.ID,
			Name:  eItem.Name,
			Price: float32(price),
		})
	}

	return mItems, nil
}
