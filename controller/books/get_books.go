package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulii/Library/model"
)


//  GetBooks 	godoc
//	@Summary		Lists all books in the library.
//	@Description	Lists all books in the library.
//	@Tags			GetBooks
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	GetBooksResponse
//	@Failure		400			{object}	ErrResponse
//	@Failure		500			{object}	ErrResponse
//	@Router			/books [get]
func (h *handler) GetBooks(ctx *gin.Context) {
	// Get all books from the database
	// TODO: return paginated results.

	var books []model.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, ErrResponse{
			Error: result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, GetBooksResponse{
		Books: books,
	})
}