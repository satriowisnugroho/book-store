package v1

import (
	"encoding/json"
	"fmt"

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
		h.POST("/login", r.Login)
	}
}

// @Summary     Register
// @Description An API to register
// @ID          create user
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Param       request		body		entity.RegisterPayload		true		"payload"
// @Success     200 {object} response.SuccessBody{data=entity.User,meta=response.MetaInfo}
// @Failure     422 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	msg := "http - v1 - User - Register"

	var payload entity.RegisterPayload
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		h.Logger.Error(err, fmt.Sprintf("%s: Decode payload", msg))
		response.Error(c, err)

		return
	}

	user, err := h.UserUsecase.CreateUser(c.Request.Context(), &payload)
	if err != nil {
		h.Logger.Error(err, fmt.Sprintf("%s: CreateUser", msg))
		response.Error(c, err)

		return
	}

	response.OK(c, user, "Successfully create an user")
}

// @Summary     Login
// @Description An API to login
// @ID          login
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Param       request		body		entity.LoginPayload		true		"payload"
// @Success     200 {object} response.SuccessBody{data=entity.User,meta=response.MetaInfo}
// @Failure     401 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	msg := "http - v1 - User - Login"

	var payload entity.LoginPayload
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		h.Logger.Error(err, fmt.Sprintf("%s: Decode payload", msg))
		response.Error(c, err)

		return
	}

	res, err := h.UserUsecase.Login(c.Request.Context(), &payload)
	if err != nil {
		h.Logger.Error(err, fmt.Sprintf("%s: Login", msg))
		response.Error(c, err)

		return
	}

	response.OK(c, res, "")
}
