import { displayScore } from './displayScore.js';

let answers = []; // Initialize an array to store the answers

// Function to fetch quiz details and questions from the backend API
export async function fetchQuizDetails(quizId) {
    return fetch(`/api/user/quizzes/${quizId}`)
        .then(response => response.json())
        .catch(error => {
            console.error('Error fetching quiz details:', error);
            throw new Error('Failed to fetch quiz details');
        });
}

function fetchQuizQuestions(quizId) {
    return fetch(`/api/user/questions/${quizId}`)
        .then(response => response.json())
        .catch(error => {
            console.error('Error fetching quiz questions:', error);
            throw new Error('Failed to fetch quiz questions');
        });
}


// Function to fetch quiz ID from the join code
export async function fetchQuizIdFromJoinCode(joinCode) {
    return fetch(`/join/${joinCode}`)
        .then(response => response.json())
        .catch(error => {
            console.error('Error fetching quiz ID from join code:', error);
            throw new Error('Failed to fetch quiz ID');
        });
}

// Function to display quiz details and questions
export async function showQuizDetails(quiz_id) {
    try {
        // Extract the join code from the URL
        displayQuizDetails(quiz_id);

    } catch (error) {
        console.error(error.message);
        alert('Failed to load quiz details. Please try again later.');
    }
}

// Function to extract quiz ID from the URL
function extractQuizIdFromUrl() {
    // Extract the quiz ID from the URL pathname
    const pathname = window.location.pathname;
    const match = pathname.match(/^\/quiz\/(\d+)$/);
    return match ? match[1] : null;
}

// Function to display quiz details using the given quiz ID
async function displayQuizDetails(quizId) {
    try {
        // Reset the answers array when displaying quiz details
        answers = [];
        const quiz = await fetchQuizDetails(quizId);
        console.log('Quiz details:', quiz);

        // Update quiz details section with quiz title and description
        const quizTitle = document.getElementById('quizTitle');
        const quizDescription = document.getElementById('quizDescription');
        quizTitle.textContent = quiz.name;
        quizDescription.innerHTML = `🧑‍💻 Join code: <span id="joinCode">${quiz.join_code}</span> <button onclick="navigator.clipboard.writeText('${quiz.join_code}')"> 📋</button><br>ℹ️ <i>${quiz.description}</i>`;

        // Clear existing form elements
        const quizForm = document.getElementById('quizForm');
        quizForm.innerHTML = '';

        const quizQuestions = await fetchQuizQuestions(quizId);
        console.log('Quiz questions:', quizQuestions); // Log the value of quizQuestions

        if (quizQuestions !== null && quizQuestions.length > 0) {
            let currentQuestionIndex = 0; // Index of the current question being displayed

            const displayNextQuestion = () => {
                // Remove the current question from the form
                quizForm.innerHTML = '';

                const question = quizQuestions[currentQuestionIndex];

                const questionLabel = document.createElement('label');
                questionLabel.textContent = question.question_text;
                quizForm.appendChild(questionLabel);

                // Create a group for radio buttons
                const optionGroup = document.createElement('div');
                optionGroup.classList.add('option-group');

                // Iterate over each option field and add radio buttons for them
                for (let i = 1; i <= 4; i++) {
                    const option = question[`option${i}`]; // Access each option field dynamically
                    const optionLabel = document.createElement('label');
                    const optionInput = document.createElement('input');

                    optionInput.type = 'radio';
                    optionInput.name = `question${question.id}`;
                    optionInput.value = i;
                    optionInput.addEventListener('change', () => {
                        // Store the selected option in the question object
                        question.selectedOption = i;

                        // Push the selected answer to the answers array
                        answers.push({ question_id: question.id, answer: i });

                        // If this is the last question, display the score
                        if (currentQuestionIndex === quizQuestions.length - 1) {
                            displayScore(quiz.id, quizQuestions, answers);
                        } else {
                            // Move to the next question
                            currentQuestionIndex++;
                            displayNextQuestion();
                        }
                    });

                    optionLabel.appendChild(optionInput);
                    optionLabel.appendChild(document.createTextNode(option));
                    optionGroup.appendChild(optionLabel);
                }

                // Append the option group to the form
                quizForm.appendChild(optionGroup);
            };

            // Display the first question
            displayNextQuestion();

        } else {
            alert('No questions are present for this quiz. It might be new?');
            return;
        }

        // Display the quiz details section
        document.getElementById('quizDetails').style.display = 'block';
    } catch (error) {
        console.error(error.message);
        alert('Failed to load quiz details. Please try again later.');
    }
}
