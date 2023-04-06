package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/api"
)

func ConfigRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/questions/all", api.FindQuestion)
		v1.GET("/questions", api.FindQuestion)
		v1.GET("/questions/:id", api.SearchQuestion)
		v1.POST("/questions", api.CreateOneQuestion)
		v1.POST("/questions/multi", api.CreateBatchQuestion)
		v1.PUT("/questions", api.UpdateQuestion)
		v1.DELETE("/questions/:id", api.DeleteQuestion)
		v1.DELETE("/questions/removeall", api.RemoveAll)
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/questions", api.FindQuestion)
		v2.GET("/questions/:id", api.SearchQuestion)
		v2.POST("/questions", api.CreateOneQuestion)
		v2.PUT("/questions", api.UpdateQuestion)
		v2.DELETE("/questions/:id", api.DeleteQuestion)
	}
}
