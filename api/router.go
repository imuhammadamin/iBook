package api

import (
	"iBook/api/handlers"
	"iBook/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter(store storage.IStorage) *gin.Engine {
	handler := handlers.New(store)

	router := gin.Default()

	router.POST("/authors", handler.CreateAuthor)
	router.GET("/authors", handler.GetAllAuthors)
	router.GET("/authors/:id", handler.GetAuthorById)
	router.PUT("/authors", handler.UpdateAuthor)
	router.DELETE("/authors/:id", handler.DeleteAuthor)

	router.POST("/books", handler.CreateBook)
	router.GET("/books", handler.GetAllBooks)
	router.GET("/books/:id", handler.GetBookById)
	router.PUT("/books", handler.UpdateBook)
	router.DELETE("/books/:id", handler.DeleteBook)

	return router
}
