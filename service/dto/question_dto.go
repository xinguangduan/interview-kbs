package dto

type QuestionDTO struct {
	QuestionDesc string `json:"question,omitempty"`
	AnswerDesc   string `json:"answerDesc"`
	Language     string `json:"language"`
}
