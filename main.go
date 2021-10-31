package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardohoraciosanto/BootcapWithGoKit/config"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/cache"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/item"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/pkg/service"
	"github.com/eduardohoraciosanto/BootcapWithGoKit/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-redis/redis/v8"
)

func main() {
	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowInfo())
	logger = log.With(logger, "caller", log.DefaultCaller)

	level.Info(logger).Log("action", "Application Started", "app_version", config.GetVersion())

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.GetEnvString(config.RedisServerKey, ""),

		Password: config.GetEnvString(config.RedisPasswordKey, ""),
	})

	cacheClient := cache.NewRedisCache(
		logger,
		0,
		redisClient,
	)

	itemsExternalService := item.NewExternalService(logger, &http.Client{
		Timeout: time.Second * 10,
	})

	svc := service.NewCartService(
		config.GetVersion(),
		cacheClient,
		itemsExternalService,
	)

	svc = service.NewServiceWithLogger(svc, logger)

	httpTransportRouter := transport.NewHTTPRouter(svc, transport.NewHTTPLogger(logger))

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", config.GetPort()),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      httpTransportRouter,
	}
	level.Info(logger).Log(
		"action", "Transport Start",
		"transport", "http",
		"port", config.GetPort())

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			level.Error(logger).Log(
				"action", "Transport Stopped",
				"transport", "http",
				"reason", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	level.Info(logger).Log("action", "Service gracefully shutted down")
	os.Exit(0)
}
