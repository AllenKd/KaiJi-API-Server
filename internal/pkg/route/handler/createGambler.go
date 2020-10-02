package handler

import (
	"KaiJi-API-Server/internal/pkg/db"
	"KaiJi-API-Server/internal/pkg/db/collection"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateGambler(c *gin.Context) {
	gambler := new(collection.Gambler)
	if err := c.Bind(gambler); err != nil {
		log.Error("invalid payload: ", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	log.Debug("create gambler: ", gambler.String())

	if err := db.New().CreateGambler(gambler); err != nil {
		log.Error("fail to create gambler: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Debug("create gambler successful: ", gambler.String())
	c.JSON(http.StatusOK, nil)
	return
}
