package v1

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/entity"
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
		h.POST("/", r.CreateOrder)
		h.GET("/", r.GetOrderHistory)
	}
}

// @Summary     Create an Order
// @Description An API to create an order
// @ID          create order
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Param       request		body		entity.OrderPayload		true		"payload"
// @Success     200 {object} response.SuccessBody{data=entity.Order,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     422 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var payload entity.OrderPayload
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		h.Logger.Error(err, "http - v1 - Decode payload")
		response.Error(c, err)

		return
	}

	order, err := h.OrderUsecase.CreateOrder(c.Request.Context(), &payload)
	if err != nil {
		h.Logger.Error(err, "http - v1 - CreateOrder")
		response.Error(c, err)

		return
	}

	response.OK(c, order, "Successfully create an order")
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
