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
        title.style.fontWeight = 'bold'; // Set the font weight of the title

        const description = document.createElement('p');
        description.textContent = quiz.description;

        const startButton = document.createElement('button');
        startButton.textContent = 'Start Quiz';
        startButton.style.fontWeight = 'bold'; // Set the font weight of the title
        startButton.style.backgroundColor = '#1dc26a'; // Set the background color of the button
        startButton.style.borderRadius = '0.5rem'; // Set the border radius of the button
        startButton.style.padding = '1rem'; // Set the margin of the button
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
// Function to display quiz details and questions
async function showQuizDetails(quizId) {
    try {
        const quiz = await fetchQuizDetails(quizId);
        console.log('Quiz details:', quiz);

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
                    // Inside the loop that creates radio buttons
                    optionInput.style.display = 'none'; // Hide the default radio button
                    optionLabel.style.backgroundColor = '#92b0e8'; // Set the background color of the button
        
                    optionLabel.style.display = 'inline-block'; // Display the label as a block element
                    optionLabel.style.margin = '3em'; // Add margin to the label
                    optionLabel.style.padding = '1em'; // Add padding to the label
                    optionLabel.style.marginBottom = '0.5rem'; // Add margin bottom to the label
                    optionLabel.style.border = '2px solid #000'; // Add border to the label
                    optionLabel.style.borderRadius = '0.5rem'; // Add border radius to the label
                    optionLabel.style.cursor = 'pointer'; // Add cursor pointer to the label
                    optionLabel.style.transition = 'all 0.3s ease'; // Add transition effect to the label

                    optionInput.type = 'radio';
                    optionInput.name = `question${question.id}`;
                    optionInput.value = i;
                    optionInput.addEventListener('change', () => {
                        // Store the selected option in the question object
                        question.selectedOption = i;

                        // Move to the next question if not the last one
                        if (currentQuestionIndex < quizQuestions.length - 1) {
                            currentQuestionIndex++;
                            displayNextQuestion();
                        } else {
                            // Notify the user that the quiz is done
                            alert('Congratulations! You have completed the quiz.');
                            // Optionally, you can hide the quiz section
                            document.getElementById('quizDetails').style.display = 'none';
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
            alert('No questions is present for this quiz. It might be new?');
            return;
        }

        // Display the quiz details section
        document.getElementById('quizDetails').style.display = 'block';
    } catch (error) {
        console.error(error.message);
        //alert('Failed to load quiz details. Please try again later.');
    }
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