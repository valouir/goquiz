package models

type SubmittedAnswers struct {
	QuestionID int `json:"question_id"`
	AnswerID   int `json:"answer_id"`
}
