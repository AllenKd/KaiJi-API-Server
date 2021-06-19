package handler

import (
	"KaiJi-Admin/internal/pkg/db/collection"
	"KaiJi-Admin/internal/pkg/kjError"
	"KaiJi-Admin/internal/pkg/service/user"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type resp struct {
	Token string `json:"token"`
}

func UserSignIn(c *gin.Context) {
	log.Debug("handler user sign in")

	var userData collection.User
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, kjError.InvalidBody.WithDetail(err.Error()))
		return
	}

	token, err := user.SignIn(userData)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	res := resp{
		Token: token,
	}
	c.JSON(http.StatusOK, res)
	log.Debug("user sign in successfully")
	return
}
