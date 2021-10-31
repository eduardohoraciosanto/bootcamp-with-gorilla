package transport

import (
	"context"
	"net/http"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/endpoints"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/service"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/viewmodels"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func NewHTTPRouter(svc service.CartService, logger log.Logger) *mux.Router {
	r := mux.NewRouter()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(viewmodels.ErrorEncoder),
		httptransport.ServerBefore(jwt.HTTPToContext(), populateContext()),
	}

	healthHandler := httptransport.NewServer(
		httpLoggingMiddleware(log.With(logger, "endpoint_name", "health"))(endpoints.Health(svc)),
		viewmodels.DecodeHealthRequest,
		viewmodels.EncodeHealthResponse,
		opts...,
	)

	createCartHandler := httptransport.NewServer(
		endpoints.CreateCart(svc),
		viewmodels.DecodeCreateCartRequest,
		viewmodels.EncodeCartResponse,
		opts...,
	)

	getCartHandler := httptransport.NewServer(
		endpoints.GetCart(svc),
		viewmodels.DecodeGetCartRequest,
		viewmodels.EncodeCartResponse,
		opts...,
	)

	addItemHandler := httptransport.NewServer(
		endpoints.AddItemToCart(svc),
		viewmodels.DecodeAddItemToCartRequest,
		viewmodels.EncodeCartResponse,
		opts...,
	)

	modifyItemQuantityHandler := httptransport.NewServer(
		endpoints.ModifyItemQuantity(svc),
		viewmodels.DecodeModifyItemQuantityRequest,
		viewmodels.EncodeCartResponse,
		opts...,
	)

	deleteItemHandler := httptransport.NewServer(
		endpoints.DeleteItem(svc),
		viewmodels.DecodeDeleteItemRequest,
		viewmodels.EncodeCartResponse,
		opts...,
	)

	deleteAllItemsHandler := httptransport.NewServer(
		endpoints.DeleteAllItems(svc),
		viewmodels.DecodeDeleteAllItemsRequest,
		viewmodels.EncodeCartResponse,
		opts...,
	)

	deleteCartHandler := httptransport.NewServer(
		endpoints.DeleteCart(svc),
		viewmodels.DecodeDeleteCartRequest,
		viewmodels.EncodeDeleteCartResponse,
		opts...,
	)

	getAvailableItemsHandler := httptransport.NewServer(
		endpoints.GetAvailableItems(svc),
		viewmodels.DecodeGetAvailableItemsRequest,
		viewmodels.EncodeGetAvailableItemsResponse,
		opts...,
	)

	getItemHandler := httptransport.NewServer(
		endpoints.GetItem(svc),
		viewmodels.DecodeGetItemRequest,
		viewmodels.EncodeGetItemResponse,
		opts...,
	)

	r.HandleFunc("/health", healthHandler.ServeHTTP).Methods(http.MethodGet)
	r.HandleFunc("/cart", createCartHandler.ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/cart/{cart_id}", getCartHandler.ServeHTTP).Methods(http.MethodGet)
	r.HandleFunc("/cart/{cart_id}/item", addItemHandler.ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/cart/{cart_id}/item/{item_id:[0-9]+}", modifyItemQuantityHandler.ServeHTTP).Methods(http.MethodPut)

	r.HandleFunc("/cart/{cart_id}/item/all", deleteAllItemsHandler.ServeHTTP).Methods(http.MethodDelete)
	r.HandleFunc("/cart/{cart_id}/item/{item_id:[0-9]+}", deleteItemHandler.ServeHTTP).Methods(http.MethodDelete)
	r.HandleFunc("/cart/{cart_id}", deleteCartHandler.ServeHTTP).Methods(http.MethodDelete)

	r.HandleFunc("/items/available", getAvailableItemsHandler.ServeHTTP).Methods(http.MethodGet)
	r.HandleFunc("/items/{item_id}", getItemHandler.ServeHTTP).Methods(http.MethodGet)

	r.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("./swagger"))))
	return r
}

func NewHTTPLogger(logger log.Logger) log.Logger {
	return log.With(logger, "transport_layer", "HTTP")
}

func populateContext() httptransport.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		// use go-kit function to populate
		// context with common http headers
		ctx = httptransport.PopulateRequestContext(ctx, req)

		refKey := httptransport.ContextKeyRequestXRequestID
		refId := ctx.Value(refKey).(string)
		if refId == "" {
			refId = uuid.New().String()
			// create a reference id for log tracing
			ctx = context.WithValue(ctx, refKey, refId)
		}

		return ctx
	}
}

func getCorrelationID(ctx context.Context) string {
	return ctx.Value(httptransport.ContextKeyRequestXRequestID).(string)
}
