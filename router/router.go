package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/api"
)

func ConfigRouter(router *gin.Engine) {
	questionApi := api.NewQuestionApi()
	v1 := router.Group("/api/v1/questions")
	{
		v1.GET("", questionApi.GetQuestionList)
		v1.GET("/:uid", questionApi.GetQuestionByUid)
		v1.POST("", questionApi.AddQuestion)
		v1.POST("/multi", questionApi.BatchAddQuestion)
		v1.PUT("", questionApi.UpdateQuestion)
		v1.DELETE("/:uid", questionApi.DeleteQuestionByUid)
		v1.DELETE("/removeall", questionApi.DeleteAll)
	}

	v2 := router.Group("/api/v2/questions")
	{
		v2.GET("", questionApi.GetQuestionList)
		v2.GET("/:uid", questionApi.GetQuestionByUid)
		v2.POST("", questionApi.AddQuestion)
		v2.PUT("", questionApi.UpdateQuestion)
		v2.DELETE("/:uid", questionApi.DeleteQuestionByUid)
	}
}
