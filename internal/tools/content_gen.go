package tools

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"lootlocker-quiz/internal/db" // Correct import path

	"github.com/go-resty/resty/v2"
)

// QuizQuestion represents a quiz question with options and the correct answer
type QuizQuestion struct {
	QuestionText  string    `json:"question"`
	Options       [4]string `json:"options"`
	CorrectOption int       `json:"correct_option"`
}

// GenerateQuizQuestions generates quiz questions using OpenAI API
func GenerateQuizQuestions(subject string) ([]QuizQuestion, error) {
	client := resty.New()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPENAI_API_KEY environment variable not set")
	}

	response, err := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model": "gpt-3.5-turbo",
			"messages": []map[string]string{
				{"role": "system", "content": "You are a helpful assistant."},
				{"role": "user", "content": fmt.Sprintf("Generate one quiz with 15 questions, where the quiz has a unique subject on %s and the 15 questions are paired to the quiz subject. Provide 4 options for each question and indicate the correct option. Do not prefix questions with numbers, and do not prefix answer options with numbers 1, 2, 3, 4 or letters like A, B, C, D. The output format should be code language readable in JSON, as we need Golang to intercept it and use it. Inside options if we have numberic values these should be strings.", subject)},
			},
			"max_tokens": 2048,
		}).
		Post("https://api.openai.com/v1/chat/completions")

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response.Body(), &result); err != nil {
		return nil, err
	}

	// Debugging: Print the entire response
	//log.Printf("Response from OpenAI: %+v\n", result)

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return nil, errors.New("unexpected response structure or no choices returned")
	}

	questions := []QuizQuestion{}
	for _, choice := range choices {
		choiceMap, ok := choice.(map[string]interface{})
		if !ok {
			return nil, errors.New("expected each choice to be a map")
		}
		message, ok := choiceMap["message"].(map[string]interface{})
		if !ok {
			return nil, errors.New("expected 'message' field to be a map")
		}
		text, ok := message["content"].(string)
		if !ok {
			return nil, errors.New("expected 'content' field to be a string")
		}

		// Remove backticks from the JSON content
		text = strings.TrimSpace(text)
		if strings.HasPrefix(text, "```json") && strings.HasSuffix(text, "```") {
			text = text[7 : len(text)-3]
		}

		// Debugging: Print the cleaned text
		//log.Printf("Cleaned JSON content: %s", text)

		// Parse the JSON content
		var parsedData struct {
			QuizSubject string `json:"subject"`
			Questions   []struct {
				Question      string   `json:"question"`
				Options       []string `json:"options"`
				CorrectOption string   `json:"correct_option"`
			} `json:"questions"`
		}
		if err := json.Unmarshal([]byte(text), &parsedData); err != nil {
			return nil, fmt.Errorf("failed to parse JSON: %v", err)
		}

		// Convert the parsed questions into QuizQuestion structs
		for _, q := range parsedData.Questions {
			if len(q.Options) != 4 {
				log.Printf("Skipping question due to invalid number of options: %v", q)
				continue
			}

			options := [4]string{}
			copy(options[:], q.Options)

			correctOptionIndex := -1
			for i, option := range options {
				if strings.TrimSpace(option) == strings.TrimSpace(q.CorrectOption) {
					correctOptionIndex = i + 1
					break
				}
			}
			if correctOptionIndex == -1 {
				return nil, fmt.Errorf("correct option '%s' not found in options for question '%s'", q.CorrectOption, q.Question)
			}

			questions = append(questions, QuizQuestion{
				QuestionText:  q.Question,
				Options:       options,
				CorrectOption: correctOptionIndex,
			})
		}
	}

	return questions, nil
}

// SaveQuizToDatabase saves generated quiz questions to the database
func SaveQuizToDatabase(database *sql.DB, subject string, questions []QuizQuestion) error {
	res, err := database.Exec("INSERT INTO quiz (name, description, join_code) VALUES (?, ?, ?)", subject, "AI-quiz generated on "+subject, db.GenerateJoinCode())
	if err != nil {
		return err
	}

	quizID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	log.Printf("Quiz ID: %d", quizID) // Log the quiz ID

	for _, question := range questions {
		randomizeAnswers(&question) // Randomize answers just before inserting into the database
		_, err := database.Exec("INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_option) VALUES (?, ?, ?, ?, ?, ?, ?)",
			quizID, question.QuestionText, question.Options[0], question.Options[1], question.Options[2], question.Options[3], question.CorrectOption)
		if err != nil {
			log.Printf("Error inserting question: %v", err)
			return err
		}
		log.Printf("Inserted question: %s", question.QuestionText) // Log the inserted question
	}

	return nil
}

func randomizeAnswers(question *QuizQuestion) {
	rand.Seed(time.Now().UnixNano())
	options := question.Options[:]
	correctAnswer := options[question.CorrectOption-1]

	// Shuffle the options
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	// Find the new position of the correct answer
	for i, option := range options {
		if option == correctAnswer {
			question.CorrectOption = i + 1
			break
		}
	}

	copy(question.Options[:], options)
}

func cleanOption(option string) string {
	// Remove common prefixes like "a) ", "1. ", etc.
	re := regexp.MustCompile(`^[a-dA-D0-9]+\W+`)
	return re.ReplaceAllString(option, "")
}
