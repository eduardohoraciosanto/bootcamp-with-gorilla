package endpoints

import (
	"context"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/service"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/viewmodels"
	"github.com/go-kit/kit/endpoint"
)

func CreateCart(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.CreateCart(ctx)
	}
}

func GetCart(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		vmReq := request.(viewmodels.GetCartRequest)
		return svc.GetCart(ctx, vmReq.CartID)
	}
}

func AddItemToCart(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		vmReq := request.(viewmodels.AddItemToCartRequest)
		return svc.AddItemToCart(ctx, vmReq.CartID, vmReq.ID, vmReq.Quantity)
	}
}

func ModifyItemQuantity(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		vmReq := request.(viewmodels.ModifyItemQuantityRequest)
		return svc.ModifyItemInCart(ctx, vmReq.CartID, vmReq.ID, vmReq.Quantity)
	}
}

func DeleteItem(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		vmReq := request.(viewmodels.DeleteItemRequest)
		return svc.DeleteItemInCart(ctx, vmReq.CartID, vmReq.ID)
	}
}

func DeleteAllItems(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		vmReq := request.(viewmodels.DeleteAllItemsRequest)
		return svc.DeleteAllItemsInCart(ctx, vmReq.CartID)
	}
}

func DeleteCart(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		vmReq := request.(viewmodels.DeleteCartRequest)
		return nil, svc.DeleteCart(ctx, vmReq.CartID)
	}
}

func GetAvailableItems(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.GetAvailableItems(ctx)
	}
}

func GetItem(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(viewmodels.GetItemRequest)
		return svc.GetItem(ctx, req.ID)
	}
}
