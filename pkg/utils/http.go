package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishError(c *gin.Context, err error, code int) {
	if err != nil {
		_ = c.Error(err)
	}

	c.JSON(code, gin.H{
		"error": err.Error(),
	})
}

func PublishData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
