package model

import "gorm.io/gorm"


type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
}

func NewBook(title, author, description string) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Description: description,
	}
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) GetDescription() string {
	return b.Description
}
