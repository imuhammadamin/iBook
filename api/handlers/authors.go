package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"iBook/api/models"
	"log"
	"net/http"
)

func (h Handler) CreateAuthor(c *gin.Context) {
	author := models.Author{}

	err := c.ShouldBindJSON(&author)
	if err != nil {
		log.Println("error on [11 context shouldbindjson(author)]")
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Author().Create(author)
	if err != nil {
		handleResponse(c, "error while creating new author", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, nil)
}

func (h Handler) UpdateAuthor(c *gin.Context) {
	author := models.Author{}

	err := c.ShouldBindJSON(&author)
	if err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	err = h.Store.Author().Update(author)
	if err != nil {
		handleResponse(c, "error while updating exist author", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, author.ID)
}

func (h Handler) GetAuthorById(c *gin.Context) {
	strId := c.Param("id")

	id := cast.ToInt(strId)

	author, err := h.Store.Author().GetById(id)
	if err != nil {
		handleResponse(c, "error: author not found", http.StatusNotFound, err.Error())
		return
	}

	handleResponse(c, "Found", http.StatusOK, author)
}

func (h Handler) GetAllAuthors(c *gin.Context) {
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

	fmt.Println(paginationMetadata)

	authors, err := h.Store.Author().GetAll(paginationMetadata)
	if err != nil {
		handleResponse(c, "any authors not found", http.StatusNotFound, err.Error())
		return
	}

	c.Header("count_of_authors", string(len(authors)))

	handleResponse(c, "Found", http.StatusOK, authors)
}

func (h Handler) DeleteAuthor(c *gin.Context) {
	strId := c.Param("id")

	id := cast.ToInt(strId)

	err := h.Store.Author().Delete(id)
	if err != nil {
		handleResponse(c, "error while deleting author", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Deleted successfully", http.StatusOK, nil)
}
