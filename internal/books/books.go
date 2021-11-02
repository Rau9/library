package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/Rau9/library/internal/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TODO: replace models.Book with something to serialise

type ReadResponse struct {
	Book  models.Book `json:"book,omitempty"`
	Error string      `json:"error,omitempty"`
}

type BookRequest struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      Author `json:"author"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type BookUpdateRequest struct {
	BookRequest
	AuthorID uuid.UUID `json:"author_id"`
}

type Author struct {
	Name        string `json:"name"`
	Nick        string `json:"nick,omitempty"`
	DateOfBirth string `json:"date_of_birth"`
	DateOfDeath string `json:"date_of_death,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type DeleteResponse struct {
	Ok    bool   `json:"deleted,omitempty"`
	Error string `json:"error,omitempty"`
}

type ListResponse struct {
	Books []models.Book `json:"books"`
	Error string        `json:"error,omitempty"`
}

func NewListResponse() *ListResponse {
	return &ListResponse{}
}

func NewReadResponse() *ReadResponse {
	return &ReadResponse{}
}

func NewBookRequest() *BookRequest {
	return &BookRequest{}
}

func NewBookUpdateRequest() *BookUpdateRequest {
	return &BookUpdateRequest{}
}

func NewDeleteResponse() *DeleteResponse {
	return &DeleteResponse{}
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func List(dbclient *gorm.DB, logger *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := NewListResponse()
		booksList, err := models.ListBooks(dbclient)
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.Books = booksList
		return c.JSON(http.StatusOK, res)
	}
}

func Read(dbclient *gorm.DB, logger *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := NewReadResponse()
		id, err := uuid.Parse(path.Base(c.Request().URL.Path))
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusUnprocessableEntity, res)
		}
		book, err := models.GetBook(dbclient, id)
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			if err.Error() == "record not found" {
				return c.JSON(http.StatusNotFound, res)
			}
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.Book = *book
		return c.JSON(http.StatusOK, res)
	}
}

func Create(dbclient *gorm.DB, logger *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := NewErrorResponse()
		book := NewBookRequest()
		bodyBytes, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		if err := json.Unmarshal(bodyBytes, book); err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusUnprocessableEntity, res)
		}

		// TODO: validate format of ISBN, if failed it should return 422

		author := models.NewAuthor()
		author.Name = book.Author.Name
		author.Nick = book.Author.Nick
		author.DateOfBirth = book.Author.DateOfBirth
		author.DateOfDeath = book.Author.DateOfDeath
		bookID, err := models.CreateBook(dbclient, book.ISBN, book.Title, &author, book.Description, book.Category)
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			if ok := strings.Contains(err.Error(), "duplicate key value"); ok {
				return c.JSON(http.StatusConflict, res)
			}
			return c.JSON(http.StatusInternalServerError, res)
		}
		c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books/%s", bookID))
		return c.NoContent(http.StatusNoContent)
	}
}

func Update(dbclient *gorm.DB, logger *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := NewErrorResponse()

		book := NewBookUpdateRequest()
		bodyBytes, err := ioutil.ReadAll(c.Request().Body)

		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusBadRequest, res)
		}
		if err := json.Unmarshal(bodyBytes, book); err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusUnprocessableEntity, res)
		}

		id, err := uuid.Parse(path.Base(c.Request().URL.Path))
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusUnprocessableEntity, res)
		}

		// TODO: validate format of ISBN, if failed it should return 422

		author := models.NewAuthor()
		author.Name = book.Author.Name
		author.Nick = book.Author.Nick
		author.DateOfBirth = book.Author.DateOfBirth
		author.DateOfDeath = book.Author.DateOfDeath
		updatedBook, err := models.UpdateBook(dbclient, id, book.ISBN, book.Title, &author, book.Description, book.Category)

		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusInternalServerError, res)
		}
		c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books/%s", updatedBook.ID))
		return c.NoContent(http.StatusNoContent)
	}
}

func Delete(dbclient *gorm.DB, logger *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := NewDeleteResponse()
		id, err := uuid.Parse(path.Base(c.Request().URL.Path))
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			return c.JSON(http.StatusUnprocessableEntity, res)
		}
		err = models.DeleteBook(dbclient, id)
		if err != nil {
			logger.Error(err.Error())
			res.Error = err.Error()
			if err.Error() == "unable to delete, record not found" {
				return c.JSON(http.StatusNotFound, res)
			}
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.Ok = true
		return c.JSON(http.StatusOK, res)
	}
}
