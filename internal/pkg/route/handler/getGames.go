package handler

import (
	"KaiJi-API-Server/internal/pkg/db"
	"KaiJi-API-Server/internal/pkg/db/collection"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

type GetGamesFilter struct {
	GameType string               `bson:"game_type,omitempty"`
	GameTime map[string]time.Time `bson:"game_time,omitempty"`
}

type GetGamesOption struct {
	Limit  int
	Offset int
	SortBy string
	Desc   bool
}

func GetGames(c *gin.Context) {
	log.Debug("getting games")

	// parse query parameters
	gameType := c.Query("game-type")
	begin, _ := strconv.ParseInt(c.Query("begin"), 10, 64)
	end, _ := strconv.ParseInt(c.Query("end"), 10, 64)

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	sortBy := c.DefaultQuery("sort-by", "_id")
	desc, _ := strconv.ParseBool(c.Query("desc"))

	// construct query filter
	f := new(GetGamesFilter)
	if gameType != "" {
		f.GameType = gameType
	}
	f.GameTime = map[string]time.Time{}
	if begin > 0 {
		f.GameTime["$gte"] = time.Unix(begin, 0)
	}
	if end > 0 {
		f.GameTime["$lte"] = time.Unix(end, 0)
	}
	var filter bson.M
	b, _ := bson.Marshal(f)
	if err := bson.Unmarshal(b, &filter); err != nil {
		log.Error("fail to unmarshal filter: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// construct query options
	option := &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
		Sort: bson.M{
			sortBy: parseOrder(desc),
		},
	}

	documents, err := db.New().GetGames(filter, option)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct {
		Count int64                   `json:"count"`
		Games []collection.SportsData `json:"games"`
	}{
		Count: db.New().CountGames(filter),
		Games: documents,
	})
}
