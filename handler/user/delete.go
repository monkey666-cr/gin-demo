package user

import (
	"github.com/chenrun666/gin_demo/handler"
	"github.com/chenrun666/gin_demo/model"
	"github.com/chenrun666/gin_demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)

}
