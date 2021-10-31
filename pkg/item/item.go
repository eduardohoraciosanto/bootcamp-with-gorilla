package item

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/viewmodels"
	"github.com/go-kit/kit/log"
)

const (
	articlesEndpoint = "http://challenge.getsandbox.com/articles"
)

type ExternalService interface {
	GetItem(id string) (models.Item, error)
	GetAllItems() ([]models.Item, error)
}

type externalService struct {
	client ItemClient
	logger log.Logger
}

type ItemClient interface {
	Get(url string) (resp *http.Response, err error)
}

func NewExternalService(logger log.Logger, client ItemClient) ExternalService {

	return &externalService{
		logger: logger,
		client: client,
	}
}

func (e *externalService) GetItem(id string) (models.Item, error) {
	res, err := e.client.Get(articlesEndpoint + "/" + id)
	if err != nil {
		return models.Item{}, err
	}
	eItem := viewmodels.ExternalItem{}

	err = json.NewDecoder(res.Body).Decode(&eItem)
	if err != nil {
		return models.Item{}, err
	}
	price, err := strconv.ParseFloat(eItem.Price, 32)
	if err != nil {
		return models.Item{}, err
	}

	mItem := models.Item{
		ID:    eItem.ID,
		Name:  eItem.Title,
		Price: float32(price),
	}

	return mItem, nil
}
func (e *externalService) GetAllItems() ([]models.Item, error) {
	res, err := e.client.Get(articlesEndpoint)
	if err != nil {
		return []models.Item{}, err
	}
	eItems := []viewmodels.ExternalItem{}

	err = json.NewDecoder(res.Body).Decode(&eItems)
	if err != nil {
		return []models.Item{}, err
	}

	mItems := []models.Item{}
	for _, eItem := range eItems {
		price, err := strconv.ParseFloat(eItem.Price, 32)
		if err != nil {
			return []models.Item{}, err
		}
		mItems = append(mItems, models.Item{
			ID:    eItem.ID,
			Name:  eItem.Title,
			Price: float32(price),
		})
	}

	return mItems, nil
}
