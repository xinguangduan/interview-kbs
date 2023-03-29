package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/dao"
	"github.com/lightsoft/interview-knowledge-base/model"
)

func SearchQuestion(c *gin.Context) {

	q := &model.QuestionEntity{
		Id:           "1121111",
		QuestionDesc: "how to do something",
		AnswerDesc:   "you to do something",
		CreateDate:   time.Now().String(),
		CreateBy:     "zhangsan",
		UpdateDate:   time.Now().String(),
		UpdateBy:     "zhangsan",
	}

	c.AbortWithStatusJSON(http.StatusOK, q)
}
func FindQuestion(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)
	q := &model.QuestionEntity{
		Id:           "1121111",
		QuestionDesc: "how to do something",
		AnswerDesc:   "you to do something",
		CreateDate:   time.Now().String(),
		CreateBy:     "zhangsan",
		UpdateDate:   time.Now().String(),
		UpdateBy:     "zhangsan",
	}

	c.AbortWithStatusJSON(http.StatusOK, q)
}
func CreateQuestion(c *gin.Context) {
	dao.Insert(context.TODO())
}

func UpdateQuestion(c *gin.Context) {
	dao.Insert(context.TODO())
}

func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("DeleteQuestion " + id)
	dao.Update(context.TODO())
}

func QueryQuestion(c *gin.Context) {
	dao.Update(context.TODO())
}
