package db

import (
	"KaiJi-API-Server/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
)

// TODO: set index
func (c *client) CreateGambler(gambler *collection.Gambler) error {
	log.Debug("create gambler: ", gambler.String())

	if _, err := c.Gambler.InsertOne(nil, gambler); err != nil {
		log.Error("fail to insert document: ", err.Error())
		return err
	}

	log.Debug("gambler created: ", gambler.String())
	return nil
}
