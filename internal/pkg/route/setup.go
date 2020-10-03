package route

import (
	"KaiJi-API-Server/internal/pkg/route/handler"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	v1 := r.Group("/kaiji/v1")
	v1.GET("/games", handler.GetGames)
	v1.POST("/gambler", handler.CreateGambler)
}
