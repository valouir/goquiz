package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valouir/goquiz/packages/data"
	"github.com/valouir/goquiz/packages/models"

	"github.com/gin-gonic/gin"
)

func SubmitAnswers(ctx *gin.Context) {

	// parse the submitted answers from the request body
	var submittedAnswers []models.SubmittedAnswers
	err := json.NewDecoder(ctx.Request.Body).Decode(&submittedAnswers)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("Failed to parse submitted answers. error: %v", err.Error()))
		return
	}

	// calculate the score
	correctCount := 0
	for _, sa := range submittedAnswers {
		for _, q := range data.Questions {
			if q.ID == sa.QuestionID && q.CorrectAnswer == sa.AnswerID {
				correctCount++
				break
			}
		}
	}

	// generate the response
	score := calculatePercentage(correctCount)

	var response string
	if correctCount == 0 {
		response = fmt.Sprintf("You answered %d out of %d questions correctly and ranked better than %.0f%% of all quizzers. Better Luck Next Time!", correctCount, len(data.Questions), score)
	} else {
		response = fmt.Sprintf("Congratz! You answered %d out of %d questions correctly! ", correctCount, len(data.Questions))
		response += fmt.Sprintf("You ranked better than %.0f%% of all quizzers.", score)
	}

	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

// calculate the percentage compared to others (hardcoded for now)
func calculatePercentage(correctCount int) float64 {
	percentage := (float64(correctCount) / float64(len(data.Questions))) * 100
	return percentage
}
