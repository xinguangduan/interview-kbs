package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lightsoft/interview-knowledge-base/dao"
	"github.com/lightsoft/interview-knowledge-base/model"
)

func SearchQuestion(c *gin.Context) {

	q := &model.QuestionEntity{
		QuestionDesc: "how to do something",
		AnswerDesc:   "you to do something",
		CreateDate:   time.Now(),
		CreateBy:     "zhangsan",
		UpdateDate:   time.Now(),
		UpdateBy:     "zhangsan",
	}

	c.AbortWithStatusJSON(http.StatusOK, q)
}

func CreateBatchQuestion(c *gin.Context) {
	var rawData = c.Request.Body

	var postData []model.QuestionEntity

	json.Marshal(rawData)

	if err := c.ShouldBind(&postData); err != nil {
		fmt.Println(err)
		return
	}
	dao.Insert(context.TODO(), postData)

	// json := struct {
	// 	Array []model.QuestionEntity
	// }{}
	// err := c.Bind(&json)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(json)

	c.AbortWithStatusJSON(http.StatusOK, `{"status":"ok"}`)
}

func CreateOneQuestion(c *gin.Context) {
	// dao.Insert(context.TODO())
	var json model.QuestionEntity
	err := c.BindJSON(&json)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json)

	dao.InsertOne(context.TODO(), json)

	c.AbortWithStatusJSON(http.StatusOK, `{"status":"ok"}`)
}

func CreateMultiple(c *gin.Context) {
	//dao.InsertMany()
}
func UpdateQuestion(c *gin.Context) {
	//dao.Insert(context.TODO())
}

func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("DeleteQuestion " + id)
	dao.Update(context.TODO())
}

func RemoveAll(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("DeleteQuestion " + id)
	dao.DeleteAll(context.TODO())
}

func FindQuestion(c *gin.Context) {
	res := dao.QueryQuestions(context.TODO())
	c.AbortWithStatusJSON(http.StatusOK, res)
}
