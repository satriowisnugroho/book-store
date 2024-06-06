package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/satriowisnugroho/book-store/internal/config"
	httpv1 "github.com/satriowisnugroho/book-store/internal/handler/http/v1"
	"github.com/satriowisnugroho/book-store/pkg/httpserver"
	"github.com/satriowisnugroho/book-store/pkg/logger"
)

func main() {
	cfg := config.NewConfig()

	l := logger.New(cfg.LogLevel)

	// HTTP Server
	handler := gin.New()
	httpv1.NewRouter(handler)
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
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - api - httpServer.Shutdown: %w", err))
	}
}
