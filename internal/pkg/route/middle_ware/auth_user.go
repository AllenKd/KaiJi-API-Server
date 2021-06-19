package middle_ware

import (
	"KaiJi-Admin/internal/pkg/configs"
	"KaiJi-Admin/internal/pkg/kjError"
	"KaiJi-Admin/internal/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type IdTokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"userId"`
}

func AuthUser(c *gin.Context) {
	log.Debug("auth user")

	idToken := util.ParseBearerToken(c.GetHeader("Authorization"))
	token, err := jwt.ParseWithClaims(idToken, &IdTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.New().JwtSign), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, kjError.UnAuthorized)
		return
	}

	claims, ok := token.Claims.(*IdTokenClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, kjError.UnAuthorized)
		return
	}
	c.Keys = map[string]interface{}{
		"userId": claims.UserId,
		"claims": claims,
	}
	c.Next()
}
