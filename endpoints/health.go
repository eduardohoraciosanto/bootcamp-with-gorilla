package endpoints

import (
	"context"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func Health(svc service.CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.Health(ctx), nil
	}
}
