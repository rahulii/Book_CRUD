package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{
		DB: db,
	}
}

func (h *handler) RegisterRoutes(router *gin.Engine) {
	// Add a book to the library.
	router.POST("/", h.AddBook)
	// Get all books in the library.
	router.GET("/books", h.GetBooks)
	// Get a book by ID.
	router.GET("/book/:book_id", h.GetBook)
	// Delete a book by ID.
	router.DELETE("/book/:book_id", h.DeleteBook)
}