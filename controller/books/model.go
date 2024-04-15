package books

import "github.com/rahulii/Library/model"

type AddBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type GetBooksResponse struct {
	Books []model.Book `json:"books"`
}

type GetBookResponse struct {
	Book model.Book `json:"book"`
}