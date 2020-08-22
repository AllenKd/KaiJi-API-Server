package main

import (
	"dataProvider/internal/pkg/configs"
	"dataProvider/internal/pkg/db"
	"dataProvider/internal/pkg/route"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.New()
	db.New()
	r := gin.Default()
	gin.SetMode(configs.New().Mode)
	route.Setup(r)

	_ = r.Run(":8080")
}
