package api

import (
	"encoding/json"
	"lootlocker-quiz/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)

// ListQuestions lists all questions associated with the specified quiz
func ListQuestions(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID := params["quizId"]

	// Define a struct to represent a question
	var question struct {
		ID            int    `json:"id"`
		QuizID        int    `json:"quiz_id"`
		Question      string `json:"question_text"`
		Option1       string `json:"option1"`
		Option2       string `json:"option2"`
		Option3       string `json:"option3"`
		Option4       string `json:"option4"`
		CorrectOption bool   `json:"correct_option"`
	}

	// Query the database to get questions and options
	rows, err := db.DB.Query("SELECT id, question_text FROM questions WHERE quiz_id = ?", quizID)
	if err != nil {
		// Handle error
		http.Error(w, "Server error - unable to find questions for quiz ID "+quizID, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Encode the questions slice into JSON and write it to the response
	json.NewEncoder(w).Encode(question)
}
