package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulii/Library/model"
)

//  DeleteBook 	godoc
//	@Summary		Delete a book by ID.
//	@Description	Delete a book by ID.
//	@Tags			DeleteBook
//	@Accept			json
//	@Produce		json
//	@Param			book_id path  int "Book ID"
//	@Success		200			{object}	DeleteBookResponse
//	@Failure		400			{object}	ErrResponse
//	@Failure		500			{object}	ErrResponse
//	@Router			/book/{book_id} [delete]
func (h *handler) DeleteBook(ctx *gin.Context) {
	book_id := ctx.Param("book_id")

	var book model.Book

	if result := h.DB.First(&book, book_id); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, ErrResponse{
			Error: result.Error.Error(),
		})
	}

	if result := h.DB.Delete(&book); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse{
			Error: result.Error.Error(),
		})
	}

	ctx.JSON(http.StatusOK, DeleteBookResponse{
		ID:      book.ID,
		Message: "Book deleted successfully",
	})
}