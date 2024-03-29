package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"iBook/api/models"
	"log"
	"net/http"
)

func (h Handler) CreateBook(c *gin.Context) {
	book := models.Book{}

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error on [11 context shouldbindjson(book)]")
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Book().Create(book)
	if err != nil {
		handleResponse(c, "error while creating new book", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, nil)
}

func (h Handler) UpdateBook(c *gin.Context) {
	book := models.Book{}

	err := c.ShouldBindJSON(&book)
	if err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Book().Update(book)
	if err != nil {
		handleResponse(c, "error while updating exist book", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, book.ID)
}

func (h Handler) GetBookById(c *gin.Context) {
	strId := c.Param("id")

	id := cast.ToInt(strId)

	book, err := h.Store.Book().GetById(id)
	if err != nil {
		handleResponse(c, "error: book not found", http.StatusNotFound, err.Error())
		return
	}

	handleResponse(c, "Found", http.StatusOK, book)
}

func (h Handler) GetAllBooks(c *gin.Context) {
	paginationMetadata := models.PaginationMetadata{}

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing page query param", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit query param", http.StatusBadRequest, err.Error())
		return
	}

	paginationMetadata.Page = page
	paginationMetadata.Limit = limit

	books, err := h.Store.Book().GetAll(paginationMetadata)
	if err != nil {
		handleResponse(c, "any books not found", http.StatusNotFound, err.Error())
		return
	}

	c.Header("count_of_books", string(len(books)))

	handleResponse(c, "Found", http.StatusOK, books)
}

func (h Handler) DeleteBook(c *gin.Context) {
	strId := c.Param("id")

	id := cast.ToInt(strId)

	err := h.Store.Book().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting book", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, nil)
}
