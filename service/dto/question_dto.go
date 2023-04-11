package dto

import (
	"time"

	"github.com/lightsoft/interview-knowledge-base/model"
)

type QuestionDTO struct {
	Uid            string    `json:"uid"`
	Language       string    `json:"language" form:"language" binding:"required" message:"语言不能为空"`
	KeyWord        string    `json:"keyWord" form:"keyWord"`
	Priority       int       `json:"priority" form:"priority"`
	QuestionDesc   string    `json:"questionDesc" form:"questionDesc" binding:"required" message:"问题不能为空"`
	AnswerDesc     string    `json:"answerDesc" form:"answerDesc" binding:"required" message:"答案不能为空"`
	CreateDate     time.Time `bson:"createDate" json:"createDate"`
	CreateBy       string    `json:"createBy"`
	UpdateDate     time.Time `bson:"updateDate" json:"updateDate"`
	UpdateBy       string    `json:"updateBy"`
	LastUpdateTime int64     `json:"lastUpdateTime"`
}

func (m *QuestionDTO) ConvertToModel(question *model.QuestionEntity) {
	question.Uid = m.Uid
	question.Language = m.Language
	question.AnswerDesc = m.AnswerDesc
	question.KeyWord = m.KeyWord
	question.QuestionDesc = m.QuestionDesc
	question.Priority = m.Priority
	question.CreateBy = m.CreateBy
	question.CreateDate = m.CreateDate
	question.UpdateBy = m.UpdateBy
	question.UpdateDate = m.UpdateDate
	question.LastUpdateTime = m.LastUpdateTime
}

// ===============================================================================
// = 更新用户相关DTO
type QuestionUpdateDTO struct {
	Uid          string    `json:"uid" form:"uid"`
	Language     string    `json:"language" form:"language" binding:"required" message:"语言不能为空"`
	KeyWord      string    `json:"keyWord" form:"keyWord"`
	Priority     int       `json:"priority" form:"priority"`
	QuestionDesc string    `json:"questionDesc" form:"questionDesc" binding:"required" message:"问题不能为空"`
	AnswerDesc   string    `json:"answerDesc" form:"answerDesc" binding:"required" message:"答案不能为空"`
	UpdateBy     string    `json:"updateBy"`
	UpdateDate   time.Time `json:"updateDate"`
}

func (m *QuestionUpdateDTO) ConvertToModel(q *model.QuestionEntity) {
	q.Language = m.Language
	q.AnswerDesc = m.AnswerDesc
	q.KeyWord = m.KeyWord
	q.QuestionDesc = m.QuestionDesc
	q.Priority = m.Priority
	q.Uid = m.Uid
	q.UpdateBy = m.UpdateBy
	q.UpdateDate = m.UpdateDate
}

// ===============================================================================
// = 用户列表相关DTO
type QuestionListDTO struct {
	Paginate
}
