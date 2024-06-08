package v1

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/pkg/logger"
)

type UserHandler struct {
	Logger      logger.LoggerInterface
	UserUsecase usecase.UserUsecaseInterface
}

func newUserHandler(handler *gin.RouterGroup, l logger.LoggerInterface, uu usecase.UserUsecaseInterface) {
	r := &UserHandler{l, uu}

	h := handler.Group("/users")
	{
		h.POST("/register", r.Register)
	}
}

// @Summary     Register an User
// @Description An API to register an user
// @ID          create user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request		body		entity.RegisterPayload		true		"payload"
// @Success     200 {object} response.SuccessBody{data=entity.User,meta=response.MetaInfo}
// @Failure     422 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var payload entity.RegisterPayload
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		h.Logger.Error(err, "http - v1 - Decode payload")
		response.Error(c, err)

		return
	}

	user, err := h.UserUsecase.CreateUser(c.Request.Context(), &payload)
	if err != nil {
		h.Logger.Error(err, "http - v1 - CreateUser")
		response.Error(c, err)

		return
	}

	response.OK(c, user, "Successfully create an user")
}
