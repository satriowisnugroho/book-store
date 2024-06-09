package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/satriowisnugroho/book-store/docs"
	"github.com/satriowisnugroho/book-store/internal/config"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Book Store API
// @description An API Documentation
// @version     1.0
// @host        localhost:9999
// @BasePath    /v1
func NewRouter(
	handler *gin.Engine,
	l logger.LoggerInterface,
	cfg *config.Config,
	bu usecase.BookUsecaseInterface,
	ou usecase.OrderUsecaseInterface,
	uu usecase.UserUsecaseInterface,
) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	h := handler.Group("/v1")
	{
		newBookHandler(h, l, bu)
		newOrderHandler(h, l, cfg, ou)
		newUserHandler(h, l, uu)
	}
}
