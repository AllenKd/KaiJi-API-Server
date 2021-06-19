package util

import (
	"crypto/sha256"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
)

func HashedPassword(password string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(password)); err != nil {
		log.Fatal("fail to hash password: ", err.Error())
	}
	return hex.EncodeToString(h.Sum(nil))
}
