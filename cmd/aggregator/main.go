// @title           Aggregator API
// @version         1.0
// @host            localhost:8080
// @BasePath        /

package main

import (
	"aggregator-go-test/internal/infra/db"
	"aggregator-go-test/internal/infra/db/repo"
	logic "aggregator-go-test/internal/logic/subscriptions"
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
	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	pool, err := db.NewPool(ctx, dsn)
	if err != nil {
		log.Fatalf("init db pool: %v", err)
	}
	defer pool.Close()

	subRepo := repo.NewSubRepoAdapter(pool)
	subSvc := logic.NewService(subRepo)
	subMgr := subscriptions.NewSubscriptionRouterManager(subSvc)

	root := chi.NewRouter()
	root.Mount("/", subMgr.Routes())
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
