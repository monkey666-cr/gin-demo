package user

import (
	"fmt"
	"github.com/chenrun666/gin_demo/handler"
	"github.com/chenrun666/gin_demo/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var r CreateRequest

	var err error
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

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

	rsp := CreateResponse{
		Username: r.Username,
	}
	handler.SendResponse(c, nil, rsp)
}
