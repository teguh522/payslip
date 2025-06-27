package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/teguh522/payslip/cmd/internal/container"
	"github.com/teguh522/payslip/cmd/internal/infrastucture/http/router"
	"github.com/teguh522/payslip/cmd/internal/infrastucture/persistence/database"
	"github.com/teguh522/payslip/cmd/internal/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	dbGORM, err := database.NewPostgreSQLGORM(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL with GORM: %v", err)
	}
	defer database.ClosePostgreSQLGORM(dbGORM)

	repos := container.NewRepositories(dbGORM)

	usecases := container.NewUseCases(repos, cfg)

	userHandler := container.NewHandlers(usecases)

	middlewares := container.NewMiddlewares(cfg)

	r := router.NewRouter(userHandler, cfg, middlewares)

	server := &http.Server{
		Addr:    ":" + cfg.App.AppPort,
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")

}
