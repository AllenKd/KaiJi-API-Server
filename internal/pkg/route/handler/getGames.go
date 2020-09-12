package handler

import (
	"apiServer/internal/pkg/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

type GetGamesReq struct {
	GameType string
	Begin string
	End string
}

func GetGames(c *gin.Context) {
	log.Debug("getting games")
	req := GetGamesReq{
		GameType: c.Query("game-type"),
		Begin: c.Query("begin"),
		End: c.Query("end"),
	}

	begin, err := time.Parse("20060102", req.Begin)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	end, err := time.Parse("20060102", req.End)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	filter := bson.M{
		"game_type": req.GameType,
		"game_time": bson.M{
			"$gte": begin,
			"$lte": end,
		},
	}
	documents, err := db.New().GetGames(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, documents)
}
