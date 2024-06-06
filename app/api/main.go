package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/satriowisnugroho/book-store/internal/config"
	httpv1 "github.com/satriowisnugroho/book-store/internal/handler/http/v1"
	"github.com/satriowisnugroho/book-store/internal/repository/postgres"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/pkg/httpserver"
	"github.com/satriowisnugroho/book-store/pkg/logger"
	pkgpostgres "github.com/satriowisnugroho/book-store/pkg/postgres"
)

func main() {
	cfg := config.NewConfig()

	l := logger.New(cfg.LogLevel)

	// Initialize postgres
	postgresDb, err := pkgpostgres.NewPostgres(&cfg.DatabaseConfig)
	if err != nil {
		l.Fatal(fmt.Errorf("app - api - postgres.NewPostgres: %w", err))
	}
	defer postgresDb.Db.Close()

	// Initialize repositories
	bookRepo := postgres.NewBookRepository(postgresDb.Db)

	// Initialize usecases
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	// HTTP Server
	handler := gin.New()
	httpv1.NewRouter(handler, l, bookUsecase)
	httpServer := httpserver.New(handler, httpserver.Port(fmt.Sprint(cfg.Port)))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - api - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - api - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - api - httpServer.Shutdown: %w", err))
	}
}
