package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulii/Library/model"
)

//  GetBook 	godoc
//	@Summary		Get a book by ID.
//	@Description	Get a book by ID.
//	@Tags			GetBook
//	@Accept			json
//	@Produce		json
//	@Param			book_id path  int "Book ID"
//	@Success		200			{object}	GetBookResponse
//	@Failure		400			{object}	ErrResponse
//	@Failure		500			{object}	ErrResponse
//	@Router			/book/{book_id} [get]
func (h *handler) GetBook(ctx *gin.Context) {
	book_id := ctx.Param("book_id")

	var book model.Book

	if result := h.DB.First(&book, book_id); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, ErrResponse{
			Error: result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, GetBookResponse{
		Book: book,
	})
}