package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulii/Library/model"
)

// AddBook 	godoc
//	@Summary		Adds a new book to the library.
//	@Description	Adds a new book to the library.
//	@Tags			AddBook
//	@Accept			json
//	@Produce		json
//	@Param			request body AddBookRequest true "Add Book Request"
//	@Success		200			{object}	SuccessResponse
//	@Failure		400			{object}	ErrResponse
//	@Failure		500			{object}	ErrResponse
//	@Router			/ [post]
func (h *handler) AddBook(c *gin.Context) {
	var book AddBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrResponse{
			Error: err.Error(),
		})
	}
	if book.Title == "" || book.Author == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrResponse{
			Error: "title and author are required fields",
		})
	
	}
	var newBook = model.NewBook(book.Title, book.Author, book.Description)
	
	if result := h.DB.Create(&newBook); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrResponse{
			Error: result.Error.Error(),
		})
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Book added successfully",
	})
}