// @title           Aggregator API
// @version         1.0
// @host            localhost:8080
// @BasePath        /

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "aggregator-go-test/docs"
	subscriptions "aggregator-go-test/internal/http-server/subscriptions"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	root := chi.NewRouter()
	subscriptionManager := subscriptions.SubscriptionRouterManager{}
	root.Mount("/", subscriptionManager.Init())
	root.Get("/swagger/*", httpSwagger.WrapHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      root,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  50 * time.Second,
	}
	go func() {
		log.Println("Listening on " + server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
