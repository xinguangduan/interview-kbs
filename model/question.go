package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type QuestionEntity struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Uid            string        `bson:"uid" json:"uid"`
	Language       string        `bson:"language" json:"language"`
	KeyWord        string        `bson:"keyWord" json:"keyWord"`
	Priority       int           `bson:"priority" json:"priority"`
	QuestionDesc   string        `bson:"questionDesc" json:"questionDesc"`
	AnswerDesc     string        `bson:"answerDesc" json:"answerDesc"`
	CreateDate     time.Time     `bson:"createDate" json:"createDate"`
	CreateBy       string        `bson:"createBy" json:"createBy"`
	UpdateDate     time.Time     `bson:"updateDate" json:"updateDate"`
	UpdateBy       string        `bson:"updateBy" json:"updateBy"`
	LastUpdateTime int64         `bson:"lastUpdateTime" json:"lastUpdateTime"`
}
