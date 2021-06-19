package user

import (
	"KaiJi-Admin/internal/pkg/configs"
	"KaiJi-Admin/internal/pkg/db"
	"KaiJi-Admin/internal/pkg/db/collection"
	"KaiJi-Admin/internal/pkg/kjError"
	"KaiJi-Admin/internal/pkg/route/middle_ware"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	tokenTtl = 24 * time.Hour
)

func SignIn(userData collection.User) (token string, err *kjError.Error) {
	log.Debug("sign in user: ", userData.Name)

	u, dbErr := db.New().GetUser(userData.Name)
	if dbErr != nil {
		err = kjError.DbOperationFail.WithDetailAndStatus(dbErr.Error(), http.StatusInternalServerError)
		return
	}

	h := sha256.New()
	if _, sErr := h.Write([]byte(userData.Password)); sErr != nil {
		log.Error("fail to hash password: ", sErr.Error())
		err = kjError.TokenSignFail.WithDetailAndStatus(sErr.Error(), http.StatusInternalServerError)
		return
	}

	if hex.EncodeToString(h.Sum(nil)) != u.Password {
		log.Warn("invalid password")
		err = kjError.SignInFail.WithStatus(http.StatusForbidden)
		return
	}

	return GenToken(u.Id.Hex())
}

func GenToken(userId string) (token string, err *kjError.Error) {
	log.Debug("gen token for user: ", userId)

	claims := middle_ware.IdTokenClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTtl).Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, signErr := tokenClaims.SignedString(configs.New().JwtSign)
	if signErr != nil {
		log.Error("fail to sign token: ", signErr.Error())
		err = kjError.TokenSignFail.WithDetailAndStatus(signErr.Error(), http.StatusInternalServerError)
	}
	token = tk
	return
}
