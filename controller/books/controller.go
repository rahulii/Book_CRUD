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
	router.GET("/books", h.GetBooks)
	router.POST("/", h.AddBook)
	router.GET("/book/:book_id", h.GetBook)
}