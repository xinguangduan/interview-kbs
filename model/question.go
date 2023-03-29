package model

type QuestionEntity struct {
	Id           string `json:"_id"`
	Lanugage     string `json:"lanugage"`
	KeyWord      string `json:"keyWord"`
	QuestionDesc string `json:"questionDesc"`
	AnswerDesc   string `json:"answerDesc"`
	CreateDate   string `json:"createDate"`
	CreateBy     string `json:"createBy"`
	UpdateDate   string `json:"updateDate"`
	UpdateBy     string `json:"updateBy"`
}
