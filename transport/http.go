package transport

import (
	"net/http"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/controller"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/health"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/service"

	"github.com/gorilla/mux"
)

func NewHTTPRouter(svc service.CartService, hsvc health.Service) *mux.Router {
	cc := controller.CartController{
		Service: svc,
	}

	ic := controller.ItemController{
		Service: svc,
	}

	hc := controller.HealthController{
		Service: hsvc,
	}

	r := mux.NewRouter()

	r.HandleFunc("/health", hc.Health).Methods(http.MethodGet)
	r.HandleFunc("/cart", cc.CreateCart).Methods(http.MethodPost)
	r.HandleFunc("/cart/{cart_id}", cc.GetCart).Methods(http.MethodGet)
	r.HandleFunc("/cart/{cart_id}/item", cc.AddItem).Methods(http.MethodPost)
	r.HandleFunc("/cart/{cart_id}/item/{item_id:[0-9]+}", cc.UpdateQuantity).Methods(http.MethodPut)

	r.HandleFunc("/cart/{cart_id}/item/all", cc.RemoveAllItems).Methods(http.MethodDelete)
	r.HandleFunc("/cart/{cart_id}/item/{item_id:[0-9]+}", cc.RemoveItem).Methods(http.MethodDelete)
	r.HandleFunc("/cart/{cart_id}", cc.DeleteCart).Methods(http.MethodDelete)

	r.HandleFunc("/items/available", ic.GetAllItems).Methods(http.MethodGet)
	r.HandleFunc("/items/{item_id}", ic.GetItem).Methods(http.MethodGet)

	r.PathPrefix("/swagger").Handler(http.StripPrefix("/swagger", http.FileServer(http.Dir("./swagger"))))
	return r
}
