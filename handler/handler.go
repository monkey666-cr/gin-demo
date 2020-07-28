package handler

import (
	"github.com/chenrun666/gin_demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
