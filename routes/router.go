package routes

import (
	"github.com/gin-gonic/gin"
	"upvote-crypto/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("upvote")
	{
		currencies := main.Group("currencies")
		{
			currencies.GET("/ranking", controllers.ShowRanking)
			currencies.GET("/", controllers.ShowCurrencies)
			currencies.POST("/", controllers.CreateVote)
			currencies.PUT("/:id", controllers.EditCurrency)
			currencies.DELETE("/:id", controllers.DeleteCurrency)
		}
	}
	return router
}