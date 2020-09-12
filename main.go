package main

import (
	"KaiJi-API-Server/internal/pkg/configs"
	"KaiJi-API-Server/internal/pkg/db"
	"KaiJi-API-Server/internal/pkg/route"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.New()
	db.New()
	r := gin.Default()
	gin.SetMode(configs.New().Mode)
	route.Setup(r)

	// start http server and listen on default port 8080
	_ = r.Run()
}
