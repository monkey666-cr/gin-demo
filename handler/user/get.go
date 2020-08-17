package user

import (
	"github.com/chenrun666/gin_demo/handler"
	"github.com/chenrun666/gin_demo/model"
	"github.com/chenrun666/gin_demo/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, user)
}