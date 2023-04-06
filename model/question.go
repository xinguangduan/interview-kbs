package model

import (
	"time"
)

type QuestionEntity struct {
	Id           string    `json:"id"`
	Language     string    `json:"language"`
	KeyWord      string    `json:"keyWord"`
	QuestionDesc string    `json:"questionDesc"`
	AnswerDesc   string    `json:"answerDesc"`
	CreateDate   time.Time `json:"createDate"`
	CreateBy     string    `json:"createBy"`
	UpdateDate   time.Time `json:"updateDate"`
	UpdateBy     string    `json:"updateBy"`
}
