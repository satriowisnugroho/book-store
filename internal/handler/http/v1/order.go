package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/pkg/logger"
)

type OrderHandler struct {
	Logger       logger.LoggerInterface
	OrderUsecase usecase.OrderUsecaseInterface
}

func newOrderHandler(handler *gin.RouterGroup, l logger.LoggerInterface, bu usecase.OrderUsecaseInterface) {
	r := &OrderHandler{l, bu}

	h := handler.Group("/orders")
	{
		h.GET("/", r.GetOrderHistory)
	}
}

// @Summary     Show History Order
// @Description An API to show history of orders
// @ID          order list
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Success     200 {object} response.SuccessBody{data=[]entity.Order,meta=response.MetaInfo}
// @Failure     500 {object} response.ErrorBody
// @Router      /orders [get]
func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	orders, err := h.OrderUsecase.GetOrdersByUserID(c.Request.Context())
	if err != nil {
		h.Logger.Error(err, "http - v1 - GetOrderHistory")
		response.Error(c, err)

		return
	}

	response.OK(c, orders, "")
}
