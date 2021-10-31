package viewmodels

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/config"
	serviceErrors "github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/errors"
)

type Meta struct {
	Version string `json:"version"`
}

type BaseResponse struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

func newBaseResponseWithData(data interface{}) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Version: config.GetVersion(),
		},
		Data: data,
	}
}

func newBaseResponseWithError(err interface{}) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Version: config.GetVersion(),
		},
		Error: err,
	}
}

func RespondWithData(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(newBaseResponseWithData(data))
}

func RespondWithError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeFromError(err))
	return json.NewEncoder(w).Encode(newBaseResponseWithError(viewModelFromError(err)))
}

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	_ = RespondWithError(w, err)
}

func statusCodeFromError(err error) int {
	mErr, ok := err.(*serviceErrors.ServiceError)

	if ok {
		switch mErr.Code {
		case serviceErrors.CartNotFoundCode, serviceErrors.ItemNotFoundCode:
			return http.StatusNotFound
		case serviceErrors.ItemAlreadyInCartCode:
			return http.StatusUnprocessableEntity
		default:
			return http.StatusInternalServerError
		}
	}
	vErr, ok := err.(*Error)
	if ok {
		switch vErr.Code {
		case ErrCodeBadRequest:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}

	return http.StatusInternalServerError
}

func descriptionFromError(mErr *serviceErrors.ServiceError) string {

	switch mErr.Code {
	case serviceErrors.CartNotFoundCode:
		return ErrDescriptionCartNotFound
	case serviceErrors.ItemAlreadyInCartCode:
		return ErrDescriptionItemAlreadyInCart
	case serviceErrors.ItemNotFoundCode:
		return ErrDescriptionItemNotFound
	}
	return ErrDescriptionInternalServerError
}

func viewModelFromError(err error) Error {
	mErr, ok := err.(*serviceErrors.ServiceError)
	if ok {
		return Error{
			Code:        mErr.Code,
			Description: descriptionFromError(mErr),
		}
	}
	vErr, ok := err.(*Error)
	if ok {
		return *vErr
	}
	return StandardInternalServerError
}
