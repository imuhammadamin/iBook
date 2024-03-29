package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"iBook/api/models"
	"iBook/storage"
)

type Handler struct {
	Store storage.IStorage
}

func New(store storage.IStorage) Handler {
	return Handler{
		Store: store,
	}
}

func handleResponse(c *gin.Context, msg string, statusCode int, data interface{}) {
	response := models.Response{}

	if statusCode >= 100 && statusCode < 200 {
		response.Message = "100-200 ERROR"
	} else if statusCode >= 200 && statusCode < 300 {
		response.Message = "SUCCESS"
	} else if statusCode >= 300 && statusCode < 400 {
		response.Message = "300-400 ERROR"
	} else if statusCode >= 400 && statusCode < 500 {
		response.Message = "SIZNI AYBINGIZ BILAN BO'LDI HAMMASI"
	} else {
		response.Message = "ENDI BU BIZNI AYB"
	}

	if msg != "" {
		response.Message = msg
	}

	response.StatusCode = statusCode
	response.Data = data

	c.JSON(statusCode, response)
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page := cast.ToUint64(pageStr)

	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}

	limit := cast.ToUint64(limitStr)

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}
