package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DashboardStrategy(c *gin.Context) {
	fakeData := map[string]float64{
		"strategy1": 0.3,
		"strategy2": 0.25,
		"strategy3": 0.33,
	}

	c.JSON(http.StatusOK, fakeData)
	return
}
