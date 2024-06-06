package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/response"
	"github.com/satriowisnugroho/book-store/internal/usecase"
	"github.com/satriowisnugroho/book-store/pkg/logger"
)

type BookHandler struct {
	Logger      logger.LoggerInterface
	BookUsecase usecase.BookUsecaseInterface
}

func newBookHandler(handler *gin.RouterGroup, l logger.LoggerInterface, bu usecase.BookUsecaseInterface) {
	r := &BookHandler{l, bu}

	h := handler.Group("/books")
	{
		h.GET("/", r.GetBooks)
	}
}

// @Summary     Show Book List
// @Description An API to show list of books
// @ID          list
// @Tags  	    book
// @Accept      json
// @Produce     json
// @Success     200 {object} response.SuccessBody{data=[]entity.Book,meta=response.MetaInfo}
// @Failure     404 {object} response.ErrorBody
// @Failure     500 {object} response.ErrorBody
// @Router      /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.BookUsecase.GetBooks(c.Request.Context())
	if err != nil {
		h.Logger.Error(err, "http - v1 - GetBooks")
		response.Error(c, err)

		return
	}

	response.OK(c, books, "")
}