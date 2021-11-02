package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Base
	ISBN        string `gorm:"size:50;not null;index"`
	Title       string `gorm:"size:100;not null;index"`
	Description string `gorm:"size:250;not null"`
	Category    string `gorm:"size:250;not null"`
	Author      Author `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
}

func NewBook() Book {
	return Book{}
}

// Listbooks returns the list of paginated books
func ListBooks(dbclient *gorm.DB) ([]Book, error) {
	// TODO: handle pagination
	booksList := make([]Book, 2)
	result := dbclient.Find(&booksList)
	if result.Error != nil {
		return nil, result.Error
	}
	for index, book := range booksList {
		author := NewAuthor()
		res := dbclient.First(&author, "book_id = ?", book.ID.String())
		if res.Error != nil {
			return nil, res.Error
		}
		booksList[index].Author = author
	}
	return booksList, result.Error
}

// GetBook returns an book by UUID
func GetBook(dbclient *gorm.DB, ID uuid.UUID) (*Book, error) {
	book := NewBook()
	author := NewAuthor()
	result := dbclient.First(&book, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	result = dbclient.First(&author, "book_id = ?", ID.String())
	book.Author = author
	return &book, result.Error
}

// CreateBook inserts a new book
func CreateBook(dbclient *gorm.DB, isbn string, title string, author *Author, description string, category string) (string, error) {
	book := NewBook()
	book.ISBN = isbn
	book.Title = title
	book.Author = *author
	book.Description = description
	book.Category = category
	result := dbclient.Create(&book)
	return book.ID.String(), result.Error
}

// UpdateBook modifies an book by UUID
func UpdateBook(dbclient *gorm.DB, ID uuid.UUID, isbn string, title string, author *Author, description string, category string) (*Book, error) {
	book := NewBook()
	book.ID = ID
	book.ISBN = isbn
	book.Title = title
	book.Author = *author
	book.Description = description
	book.Category = category
	result := dbclient.Model(&book).Updates(book)
	return &book, result.Error
}

// DeleteBook deletes an book by UUID
func DeleteBook(dbclient *gorm.DB, ID uuid.UUID) error {
	book := NewBook()
	result := dbclient.Delete(&book, ID)
	if result.Error == nil && result.RowsAffected == 0 {
		return errors.New("unable to delete, record not found")
	}
	return result.Error
}
