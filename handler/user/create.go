package user

import (
	"fmt"
	"github.com/chenrun666/gin_demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"net/http"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": errno.ErrBind,
		})
		return
	}

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can ot found in db: xx.xx.xx.xx")).Add("This is add message")
		log.Errorf(err, "Get an error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
}