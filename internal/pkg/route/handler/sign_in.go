package handler

import (
	"KaiJi-Admin/internal/pkg/db/collection"
	"KaiJi-Admin/internal/pkg/kjError"
	"KaiJi-Admin/internal/pkg/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

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

	c.JSON(http.StatusOK, fmt.Sprintf(`"token": "%s"`, token))
	log.Debug("user sign in successfully")
	return
}
