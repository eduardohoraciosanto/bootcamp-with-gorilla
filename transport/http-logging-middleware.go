package transport

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type HTTPMiddleware func(endpoint.Endpoint) endpoint.Endpoint

func httpLoggingMiddleware(logger log.Logger) HTTPMiddleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			level.Info(logger).Log(
				"http_action", "calling endpoint",
				"correlation_id", getCorrelationID(ctx),
			)

			defer func(begin time.Time) {
				level.Info(logger).Log(
					"http_action", "endpoint_called",
					"took", time.Since(begin),
					"correlation_id", getCorrelationID(ctx),
				)
			}(time.Now())
			return next(ctx, request)
		}
	}
}
