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

	var questions []struct {
		ID            int    `json:"id"`
		QuizID        int    `json:"quiz_id"`
		QuestionText  string `json:"question_text"`
		Option1       string `json:"option1"`
		Option2       string `json:"option2"`
		Option3       string `json:"option3"`
		Option4       string `json:"option4"`
		CorrectOption int    `json:"correct_option"`
	}

	rows, err := db.DB.Query("SELECT id, quiz_id, question_text, option1, option2, option3, option4, correct_option FROM questions WHERE quiz_id = ?", quizID)
	if err != nil {
		http.Error(w, "Server error - unable to fetch questions for quiz id "+quizID, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var question struct {
			ID            int    `json:"id"`
			QuizID        int    `json:"quiz_id"`
			QuestionText  string `json:"question_text"`
			Option1       string `json:"option1"`
			Option2       string `json:"option2"`
			Option3       string `json:"option3"`
			Option4       string `json:"option4"`
			CorrectOption int    `json:"correct_option"`
		}
		err := rows.Scan(&question.ID, &question.QuizID, &question.QuestionText, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.CorrectOption)
		if err != nil {
			http.Error(w, "Server error - unable to fetch questions for quiz id "+quizID, http.StatusInternalServerError)
			return
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Server error - unable to fetch questions for quiz id "+quizID, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}
