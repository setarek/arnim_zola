package router

import (
	"github.com/gin-gonic/gin"
	"github.com/setarek/arnim_zola/internal/spot"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	v1 := router.Group("api/v1")
	{
		spotGroup := v1.Group("spot")
		{
			s := new(spot.SpotHandler)
			spotGroup.POST("/kline", s.CalculateOptPoint)
		}

	}

	return router
}
