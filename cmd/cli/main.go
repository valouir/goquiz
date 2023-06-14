package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/valouir/goquiz/packages/models"

	"github.com/spf13/cobra"
)

func main() {
	command := &cobra.Command{
		Use: "Go Quiz test",
		Run: func(cmd *cobra.Command, args []string) {

			// retrieve questions
			questions, err := retrieveQuestions()
			if err != nil {
				log.Fatal(err)
			}

			// display  questions
			for _, quest := range questions {
				fmt.Println("\n" + quest.Text)
				for _, ans := range quest.Answers {
					fmt.Printf("\n%d. %s", ans.ID, ans.Text)
				}

				fmt.Println() // skip a line
			}

			// user prompt
			fmt.Println() // skip a line
			var submittedAnswers []models.SubmittedAnswers
			for _, q := range questions {
				var answerID int
				fmt.Printf("Enter the ID of the answer for question %d: ", q.ID)
				_, err := fmt.Scanln(&answerID)
				if err != nil {
					log.Fatal("Failed to read answer:", err)
				}

				submittedAnswers = append(submittedAnswers, models.SubmittedAnswers{QuestionID: q.ID, AnswerID: answerID})
			}

			// submit answers
			resp, err := submitAnswers(submittedAnswers)
			if err != nil {
				log.Fatal(err)
			}

			// display response
			fmt.Println("\nResponse: " + resp)
		},
	}

	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}

func retrieveQuestions() ([]models.Question, error) {

	// fetch questions
	resp, err := http.Get(fmt.Sprintf("http://localhost:%v/questions", 6379))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch questions: %v", err)
	}
	defer resp.Body.Close()

	// status code check
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve questions with unexpected status code %d", resp.StatusCode)
	}

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading body: %v", err)
	}

	// unmarshal data
	var questions []models.Question
	err = json.Unmarshal(body, &questions)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, fmt.Errorf("failed unmarshalling data: %v", err)
	}

	return questions, nil
}

func submitAnswers(answers []models.SubmittedAnswers) (string, error) {

	// marshal data
	jsonData, err := json.Marshal(answers)
	if err != nil {
		return "", fmt.Errorf("failed to marshal answers: %v", err)
	}

	// send answers request
	body := bytes.NewBuffer(jsonData)
	resp, err := http.Post(fmt.Sprintf("http://localhost:%v/answers", 6379), "application/json", body)
	if err != nil {
		return "", fmt.Errorf("failed to submit answers: %v", err)
	}
	defer resp.Body.Close()

	// display the response
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(response), nil
}
