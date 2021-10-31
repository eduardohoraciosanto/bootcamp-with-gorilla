package service

import (
	"context"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/cache"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/errors"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/item"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/models"
	"github.com/google/uuid"
)

type CartService interface {
	CreateCart(ctx context.Context) (models.Cart, error)
	GetCart(ctx context.Context, cartID string) (models.Cart, error)
	GetAvailableItems(ctx context.Context) ([]models.Item, error)
	GetItem(ctx context.Context, id string) (models.Item, error)
	AddItemToCart(ctx context.Context, cartID, itemID string, quantity int) (models.Cart, error)
	ModifyItemInCart(ctx context.Context, cartID, itemID string, newQuantity int) (models.Cart, error)
	DeleteItemInCart(ctx context.Context, cartID, itemID string) (models.Cart, error)
	DeleteAllItemsInCart(ctx context.Context, cartID string) (models.Cart, error)
	DeleteCart(ctx context.Context, cartID string) error
}

type service struct {
	//dependencies of the service
	version         string
	cache           cache.Cache
	externalService item.ExternalService
}

func NewCartService(version string, cache cache.Cache, externalService item.ExternalService) CartService {
	return &service{
		version:         version,
		cache:           cache,
		externalService: externalService,
	}
}

func (s *service) CreateCart(ctx context.Context) (models.Cart, error) {
	cartID := uuid.New().String()
	cart := models.Cart{
		ID: cartID,
	}

	if err := s.cache.Set(cartID, cart); err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (s *service) GetCart(ctx context.Context, cartID string) (models.Cart, error) {
	cart := models.Cart{}
	err := s.cache.Get(cartID, &cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.CartNotFoundCode}
	}

	err = s.fetchItemsForCart(&cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.ExternalApiErrorCode}
	}

	return cart, nil
}

func (s *service) GetAvailableItems(ctx context.Context) ([]models.Item, error) {
	items, err := s.externalService.GetAllItems()
	if err != nil {
		return []models.Item{}, err
	}
	return items, nil
}

func (s *service) GetItem(ctx context.Context, id string) (models.Item, error) {
	item, err := s.externalService.GetItem(id)
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (s *service) AddItemToCart(ctx context.Context, cartID, itemID string, quantity int) (models.Cart, error) {
	cart := models.Cart{}
	err := s.cache.Get(cartID, &cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.CartNotFoundCode}
	}

	for _, item := range cart.Items {
		if item.ID == itemID {
			return models.Cart{}, errors.ServiceError{Code: errors.ItemAlreadyInCartCode}
		}
	}

	cart.Items = append(cart.Items, models.Item{
		ID:       itemID,
		Quantity: quantity,
	})

	if err := s.cache.Set(cartID, cart); err != nil {
		return models.Cart{}, err
	}

	err = s.fetchItemsForCart(&cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.ExternalApiErrorCode}
	}

	return cart, nil
}
func (s *service) ModifyItemInCart(ctx context.Context, cartID, itemID string, newQuantity int) (models.Cart, error) {
	cart := models.Cart{}
	err := s.cache.Get(cartID, &cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.CartNotFoundCode}
	}

	for idx, item := range cart.Items {
		if item.ID == itemID {
			cart.Items[idx].Quantity = newQuantity
			if err := s.cache.Set(cartID, cart); err != nil {
				return models.Cart{}, err
			}
			err = s.fetchItemsForCart(&cart)
			if err != nil {
				return models.Cart{}, errors.ServiceError{Code: errors.ExternalApiErrorCode}
			}
			return cart, nil
		}
	}

	return models.Cart{}, errors.ServiceError{Code: errors.ItemNotFoundCode}
}
func (s *service) DeleteItemInCart(ctx context.Context, cartID, itemID string) (models.Cart, error) {
	cart := models.Cart{}
	err := s.cache.Get(cartID, &cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.CartNotFoundCode}
	}

	for idx, item := range cart.Items {
		if item.ID == itemID {
			//we care about the order, so we perform to sub-slices

			cart.Items = append(cart.Items[:idx], cart.Items[idx+1:]...)

			if err := s.cache.Set(cartID, cart); err != nil {
				return models.Cart{}, err
			}

			err = s.fetchItemsForCart(&cart)
			if err != nil {
				return models.Cart{}, errors.ServiceError{Code: errors.ExternalApiErrorCode}
			}

			return cart, nil
		}
	}

	return models.Cart{}, errors.ServiceError{Code: errors.ItemNotFoundCode}
}
func (s *service) DeleteAllItemsInCart(ctx context.Context, cartID string) (models.Cart, error) {
	cart := models.Cart{}
	err := s.cache.Get(cartID, &cart)
	if err != nil {
		return models.Cart{}, errors.ServiceError{Code: errors.CartNotFoundCode}
	}

	cart.Items = []models.Item{}
	if err := s.cache.Set(cartID, cart); err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}
func (s *service) DeleteCart(ctx context.Context, cartID string) error {
	err := s.cache.Del(cartID)
	if err != nil {
		return errors.ServiceError{Code: errors.CartNotFoundCode}
	}
	return nil
}
func (s *service) fetchItemsForCart(cart *models.Cart) error {
	//We fetch information from the external service to fill in Name and Price
	for idx, item := range cart.Items {
		extItem, err := s.externalService.GetItem(item.ID)
		if err != nil {
			return err
		}
		cart.Items[idx].Price = extItem.Price
		cart.Items[idx].Name = extItem.Name
	}
	return nil
}
