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

type GetGamblersFilter struct {
	Strategy    string `bson:"strategy.name,omitempty"`
	PutStrategy string `bson:"strategy.put_strategy.name,omitempty"`
}

func GetGamblers(c *gin.Context) {
	log.Debug("getting gamblers")

	// parse query parameters
	strategy := c.Query("strategy")
	putStrategy := c.Query("put-strategy")

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	sortBy := c.DefaultQuery("sort-by", "_id")
	desc, _ := strconv.ParseBool(c.Query("desc"))

	// construct query filter
	f := GetGamblersFilter{}
	if strategy != "" {
		f.Strategy = strategy
	}
	if putStrategy != "" {
		f.PutStrategy = putStrategy
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

	gamblers, err := db.New().GetGamblers(filter, option)
	if err != nil {
		log.Error("fail to get gamblers: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, struct {
		Count    int64                `json:"count"`
		Gamblers []collection.Gambler `json:"gamblers"`
	}{
		Count:    db.New().CountGamblers(filter),
		Gamblers: gamblers,
	})
}
