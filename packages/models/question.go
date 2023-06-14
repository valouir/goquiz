package models

type Question struct {
	ID            int
	Text          string
	CorrectAnswer int
	Answers       []Answer
}
