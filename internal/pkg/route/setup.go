package route

import (
	"dataProvider/internal/pkg/route/handler"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	v1 := r.Group("/kaiji/v1/games")
	v1.GET("", handler.GetGames)
}
