package data

import "github.com/valouir/goquiz/packages/models"

var Questions = []models.Question{
	{
		ID:   1,
		Text: "What is the capital of Italy?",
		Answers: []models.Answer{
			{ID: 1, Text: "Prague"},
			{ID: 2, Text: "Rome"},
			{ID: 3, Text: "London"},
		},
		CorrectAnswer: 2,
	},
	{
		ID:   2,
		Text: "What is the currency of Japan?",
		Answers: []models.Answer{
			{ID: 1, Text: "Euro"},
			{ID: 2, Text: "Crypto"},
			{ID: 3, Text: "Yen"},
		},
		CorrectAnswer: 3,
	},
	{
		ID:   3,
		Text: "What is the smallest country in the world?",
		Answers: []models.Answer{
			{ID: 1, Text: "Malta"},
			{ID: 2, Text: "Vatican City"},
			{ID: 3, Text: "Sardinia"},
		},
		CorrectAnswer: 2,
	},
	{
		ID:   4,
		Text: "What was the first feature-length animated movie ever released?",
		Answers: []models.Answer{
			{ID: 1, Text: "Dumbo"},
			{ID: 2, Text: "Bambi"},
			{ID: 3, Text: "Snow White and the Seven Dwarfs"},
		},
		CorrectAnswer: 3,
	},
}
