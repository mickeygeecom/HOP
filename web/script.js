// Function to fetch quiz data from the backend API
function fetchQuizzes() {
    fetch('/api/user/quizzes')
        .then(response => response.json())
        .then(data => {
            // Call function to display quiz items
            displayQuizzes(data);
        })
        .catch(error => {
            console.error('Error fetching quizzes:', error);
            alert('Failed to fetch quizzes');
        });
}

// Function to display quiz items in the HTML document
function displayQuizzes(quizzes) {
    const quizList = document.getElementById('quizList');

    // Clear existing quiz items
    quizList.innerHTML = '';

    // Loop through each quiz and create HTML elements to represent them
    quizzes.forEach(quiz => {
        const quizItem = document.createElement('div');
        quizItem.classList.add('quiz-item');

        const title = document.createElement('h2');
        title.textContent = quiz.name; // Display quiz name instead of title

        const description = document.createElement('p');
        description.textContent = quiz.description;

        const startButton = document.createElement('button');
        startButton.textContent = 'Start Quiz';
        startButton.addEventListener('click', () => showQuizDetails(quiz.id));

        // Append elements to quizItem
        quizItem.appendChild(title);
        quizItem.appendChild(description);
        quizItem.appendChild(startButton);

        // Append quizItem to quizList
        quizList.appendChild(quizItem);
    });
}

// Function to fetch quiz details and questions from the backend API
function fetchQuizDetails(quizId) {
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

// Function to display quiz details and questions
async function showQuizDetails(quizId) {
    try {
        const quiz = await fetchQuizDetails(quizId);
        console.log('Quiz details:', quiz)

        // Update quiz details section with quiz title and description
        const quizTitle = document.getElementById('quizTitle');
        const quizDescription = document.getElementById('quizDescription');
        quizTitle.textContent = quiz.name;
        quizDescription.textContent = quiz.description;

        // Clear existing form elements
        const quizForm = document.getElementById('quizForm');
        quizForm.innerHTML = '';

        const quizQuestions = await fetchQuizQuestions(quizId);
        console.log('Quiz questions:', quizQuestions); // Log the value of quizQuestions
        console.log('Type of quizQuestions:', typeof quizQuestions); // Log the type of quizQuestions

        // Check if quizQuestions is an array before calling forEach
        if (Array.isArray(quizQuestions) && quizQuestions.length > 0) {
            quizQuestions.forEach(question => {
                // Your existing code to create form elements for each question
            });
        } else {
            console.error('No questions found in the quiz');
            alert('Failed to load quiz details. Please try again later.');
            return;
        }


        // Create form elements for each question in the quiz if questions array exists
        console.log(JSON.stringify(quizQuestions))
        if (quizQuestions.length > 0) {
            quizQuestions.forEach(question => {
                const questionLabel = document.createElement('label');
                questionLabel.textContent = question.question_text;
                quizForm.appendChild(questionLabel);

                question.options.forEach((option, index) => {
                    const optionLabel = document.createElement('label');
                    const optionInput = document.createElement('input');
                    optionInput.type = 'radio';
                    optionInput.name = `question${question.id}`;
                    optionInput.value = index + 1;
                    optionLabel.appendChild(optionInput);
                    optionLabel.appendChild(document.createTextNode(option));
                    quizForm.appendChild(optionLabel);
                });
            });
        } else {
            console.error('No questions found in the quiz');
            alert('Failed to load quiz details. Please try again later.');
            return;
        }

        // Display the quiz details section
        document.getElementById('quizDetails').style.display = 'block';
    } catch (error) {
        console.error(error.message);
        alert('Failed to load quiz details. Please try again later.');
    }

    // Create a submit button for the quiz
    const submitButton = document.createElement('button');
    submitButton.textContent = 'Submit Quiz';
    submitButton.addEventListener('click', () => {
        const answers = [];

        // Collect selected answers from the form
        quiz.questions.forEach(question => {
            const selectedAnswer = document.querySelector(`input[name="question${question.id}"]:checked`);
            if (selectedAnswer) {
                answers.push({
                    question_id: question.id,
                    answer: parseInt(selectedAnswer.value)
                });
            }
        });

        // Submit the answers
        if (answers.length > 0) {
            submitAnswers(quizId, answers);
        } else {
            alert('Please answer all questions before submitting the quiz');
        }
    });

    // Append submit button to quizForm
    quizForm.appendChild(submitButton);

    // Display the quiz details section
    document.getElementById('quizDetails').style.display = 'block';
}

// Call fetchQuizzes function when the page loads
window.onload = fetchQuizzes;

// Function to submit quiz answers to the backend API
function submitAnswers(quizId, answers) {
    fetch(`/api/user/quizzes/${quizId}/submit`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(answers)
    })
    .then(response => {
        if (response.ok) {
            alert('Quiz answers submitted successfully');
            // You can optionally redirect the user to another page or perform other actions
        } else {
            throw new Error('Failed to submit quiz answers');
        }
    })
    .catch(error => {
        console.error('Error submitting quiz answers:', error);
        alert('Failed to submit quiz answers. Please try again later.');
    });
}