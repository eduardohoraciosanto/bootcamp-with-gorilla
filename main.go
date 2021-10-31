package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/config"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/cache"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/health"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/item"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/pkg/service"
	"github.com/eduardohoraciosanto/bootcamp-with-gorilla/transport"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	log := logrus.New()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.GetEnvString(config.RedisServerKey, ""),
		Password: config.GetEnvString(config.RedisPasswordKey, ""),
	})

	cacheClient := cache.NewRedisCache(
		log.WithField("owner", "cache").Logger,
		0,
		redisClient,
	)

	itemsExternalService := item.NewExternalService(log.WithField("owner", "external service").Logger, &http.Client{
		Timeout: time.Second * 10,
	})

	svc := service.NewCartService(
		config.GetVersion(),
		cacheClient,
		itemsExternalService,
	)

	hsvc := health.NewService(
		cacheClient,
		itemsExternalService,
	)

	httpTransportRouter := transport.NewHTTPRouter(svc, hsvc)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", config.GetPort()),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      httpTransportRouter,
	}
	log.WithField(
		"transport", "http").
		WithField(
			"port", config.GetPort()).
		Log(logrus.InfoLevel, "Transport Start")
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.WithField(
				"transport", "http").
				WithError(err).
				Log(logrus.ErrorLevel, "Transport Stopped")
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
	log.Log(logrus.InfoLevel, "Service gracefully shutted down")
	os.Exit(0)
}
