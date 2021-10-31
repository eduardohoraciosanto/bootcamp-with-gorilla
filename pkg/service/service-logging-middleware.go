package service

import (
	"context"
	"fmt"
	"time"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/models"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type loggingMiddleware struct {
	logger log.Logger
	next   CartService
}

func NewServiceWithLogger(service CartService, logger log.Logger) CartService {
	return &loggingMiddleware{
		logger: log.With(logger, "service", "Cart Service"),
		next:   service,
	}
}

func (lm *loggingMiddleware) Health(ctx context.Context) (out []models.Health) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "Health",
			"service_output", fmt.Sprintf("%+v", out),
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.Health(ctx)
}

func (lm *loggingMiddleware) CreateCart(ctx context.Context) (out models.Cart, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "CreateCart",
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.CreateCart(ctx)
}
func (lm *loggingMiddleware) GetCart(ctx context.Context, cartID string) (out models.Cart, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "GetCart",
			"service_inputs", []interface{}{cartID},
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.GetCart(ctx, cartID)
}
func (lm *loggingMiddleware) GetAvailableItems(ctx context.Context) (out []models.Item, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "GetAvailableItems",
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.GetAvailableItems(ctx)
}
func (lm *loggingMiddleware) GetItem(ctx context.Context, id string) (out models.Item, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "GetAvailableItems",
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.GetItem(ctx, id)
}
func (lm *loggingMiddleware) AddItemToCart(ctx context.Context, cartID, itemID string, quantity int) (out models.Cart, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "AddItemToCart",
			"service_inputs", []interface{}{cartID, itemID, quantity},
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.AddItemToCart(ctx, cartID, itemID, quantity)
}
func (lm *loggingMiddleware) ModifyItemInCart(ctx context.Context, cartID, itemID string, newQuantity int) (out models.Cart, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "ModifyItemInCart",
			"service_inputs", []interface{}{cartID, itemID, newQuantity},
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.ModifyItemInCart(ctx, cartID, itemID, newQuantity)
}
func (lm *loggingMiddleware) DeleteItemInCart(ctx context.Context, cartID, itemID string) (out models.Cart, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "DeleteItemInCart",
			"service_inputs", []interface{}{cartID, itemID},
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.DeleteItemInCart(ctx, cartID, itemID)
}
func (lm *loggingMiddleware) DeleteAllItemsInCart(ctx context.Context, cartID string) (out models.Cart, err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "DeleteAllItemsInCart",
			"service_inputs", []interface{}{cartID},
			"service_output", fmt.Sprintf("%+v", out),
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.DeleteAllItemsInCart(ctx, cartID)

}
func (lm *loggingMiddleware) DeleteCart(ctx context.Context, cartID string) (err error) {
	defer func(begin time.Time) {
		lm.logger.Log(
			"service_method", "DeleteCart",
			"service_inputs", []interface{}{cartID},
			"service_err", err,
			"service_took", time.Since(begin),
			"correlation_id", getCorrelationID(ctx),
		)
	}(time.Now())
	return lm.next.DeleteCart(ctx, cartID)
}

func getCorrelationID(ctx context.Context) string {
	cid, ok := ctx.Value(httptransport.ContextKeyRequestXRequestID).(string)
	if !ok {
		return ""
	}
	return cid
}
