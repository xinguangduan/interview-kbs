package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/api"
)

func ConfigRouter(router *gin.Engine) {
	questionApi := api.NewQuestionApi()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/questions/all", questionApi.GetQuestionList)
		v1.GET("/questions", questionApi.GetQuestionList)
		v1.GET("/questions/:uid", questionApi.GetQuestionByUid)
		v1.POST("/questions", questionApi.AddQuestion)
		v1.POST("/questions/multi", questionApi.CreateBatchQuestion)
		v1.PUT("/questions", questionApi.UpdateQuestion)
		v1.DELETE("/questions/:uid", questionApi.DeleteUserByUid)
		v1.DELETE("/questions/removeall", questionApi.DeleteAll)
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/questions", questionApi.GetQuestionList)
		v2.GET("/questions/:uid", questionApi.GetQuestionByUid)
		v2.POST("/questions", questionApi.AddQuestion)
		v2.PUT("/questions", questionApi.UpdateQuestion)
		v2.DELETE("/questions/:uid", questionApi.DeleteUserByUid)
	}
}
