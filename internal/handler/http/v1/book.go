package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/satriowisnugroho/book-store/internal/entity"
	"github.com/satriowisnugroho/book-store/internal/helper"
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

// @Summary     Show List of Books
// @Description An API to show list of books
// @ID          book list
// @Tags  	    Book
// @Accept      json
// @Produce     json
// @Param       keyword 		query		string 		false 	"title search by keyword"
// @Param       offset 			query 	integer 	false		"offset"
// @Param       limit 			query 	integer 	false 	"limit"
// @Success     200 {object} response.SuccessBody{data=[]entity.Book,meta=response.MetaInfo}
// @Failure     500 {object} response.ErrorBody
// @Router      /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	limit, offset := helper.GetLimitOffsetFromURLQuery(c)
	payload := entity.GetBooksPayload{
		TitleKeyword: c.Query("keyword"),
		Offset:       offset,
		Limit:        limit,
	}
	books, count, err := h.BookUsecase.GetBooks(c.Request.Context(), payload)
	if err != nil {
		h.Logger.Error(err, "http - v1 - book - GetBooks: GetBooks")
		response.Error(c, err)

		return
	}

	response.OKWithPagination(c, books, "", count, offset, limit)
}
