package api

import (
	"encoding/json"
	"lootlocker-quiz/internal/db"
	"net/http"
	"time"
)

func userListQuizzes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, description, created_at FROM quiz")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	quizzes := []map[string]interface{}{}
	for rows.Next() {
		var id int
		var name, description string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &description, &createdAt); err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		quizzes = append(quizzes, map[string]interface{}{
			"id":          id,
			"name":        name,
			"description": description,
			"created_at":  createdAt.Format(time.RFC3339),
		})
	}

	json.NewEncoder(w).Encode(quizzes)
}
