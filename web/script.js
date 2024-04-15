// ## QUIZZES ##

let currentQuizId;

// Create quiz form
document.getElementById('createQuizForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const quizName = document.getElementById('quizName').value;
    const quizDescription = document.getElementById('quizDescription').value;
    createQuiz(quizName, quizDescription);
});

// Function to create a quiz
function createQuiz(name, description) {
    fetch('https://localhost/quizzes', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: name, description: description })
    })
    .then(response => response.json())
    .then(data => {
        console.log('Quiz created:', data);
        alert('Quiz created successfully!');
        listQuizzes();  // Refresh the list after adding
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Failed to create quiz');
    });
}

// Function to list quizzes
function listQuizzes() {
    fetch('https://localhost/quizzes')
    .then(response => response.json())
    .then(data => {
        const list = document.getElementById('quizList');
        list.innerHTML = '';  // Clear existing list
        data.forEach(quiz => {
            const item = document.createElement('li');
            item.className = 'flex justify-between items-center mb-4';
            item.innerHTML = `
                <span class="flex-1">[${quiz.id}] ${quiz.name} - ${quiz.description}</span>
                <div class="flex justify-end items-center gap-2">
                    <button onclick="showQuestions(${quiz.id})" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Questions</button>
                    <button onclick="editQuiz(${quiz.id})" class="bg-orange-500 hover:bg-orange-700 text-white font-bold py-1 px-2 rounded">Edit</button>
                    <button onclick="deleteQuiz(${quiz.id})" class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded">Delete</button>
                </div>
            `;
            list.appendChild(item);
        });
        // Hide question section if it's visible
        document.getElementById('questionSection').style.display = 'none';
        document.getElementById('quizSection').style.display = 'block';
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Failed to fetch quizzes');
    });
}
listQuizzes();  // List quizzes when the page loads

// Function to edit a quiz
function editQuiz(quizId) {
    const newName = prompt('Enter new name for the quiz:');
    const newDescription = prompt('Enter new description for the quiz:');
    if (!newName || !newDescription) {
        alert('Name and description are required');
        return;
    }
    fetch(`https://localhost/quizzes/${quizId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: newName, description: newDescription })
    })
    .then(response => {
        if (response.ok) {
            listQuizzes();
        } else {
            alert('Failed to update quiz');
        }
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Error updating quiz');
    });
}

// Function to delete a quiz
function deleteQuiz(quizId) {
    if (confirm('Are you sure you want to delete this quiz?')) {
        fetch(`https://localhost/quizzes/${quizId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                alert('Quiz deleted successfully!');
                listQuizzes();
            } else {
                alert('Failed to delete quiz');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('Error deleting quiz');
        });
    }
}

// #### QUESTIONS ####

// Function to show questions for a selected quiz
function showQuestions(quizId) {
    currentQuizId = quizId;
    console.log("Setting current quiz id to", quizId);

    fetch(`https://localhost/quizzes/${quizId}/questions`)
    .then(response => response.json())
    .then(data => {
        const questionList = document.getElementById('questionList');
        questionList.innerHTML = '';  // Clear existing list

        if (!data) {
            // Handle case when there are no questions
            const noQuestionsMessage = document.createElement('p');
            noQuestionsMessage.textContent = 'No questions available for this quiz.';
            questionList.appendChild(noQuestionsMessage);
        } else {
            // Render questions if available
            data.forEach(question => {
                const item = document.createElement('li');
                item.className = 'flex justify-between items-center mb-4';
                item.innerHTML = `
                    <span class="flex-1">${question.question_text}</span>
                    <div class="flex justify-end items-center gap-2">
                        <button onclick="editQuestion(${quizId}, ${question.id})" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded">Edit</button>
                        <button onclick="deleteQuestion(${quizId}, ${question.id})" class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded">Delete</button>
                    </div>
                `;
                questionList.appendChild(item);
            });
        }

        // Show question section and hide quiz section
        document.getElementById('quizSection').style.display = 'none';
        document.getElementById('questionSection').style.display = 'block';
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Failed to fetch questions');
    });
}




// Function to go back to quiz list
function goBackToQuizList() {
    // Reload the quiz list
    listQuizzes();
}

// #### QUESTIONS ####

// Submit listener for create question form
document.getElementById('createQuestionForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const quizId = currentQuizId;
    const questionText = document.getElementById('questionText').value;
    const option1 = document.getElementById('option1').value;
    const option2 = document.getElementById('option2').value;
    const option3 = document.getElementById('option3').value;
    const option4 = document.getElementById('option4').value;
    const correctOption = document.querySelector('input[name="correctOption"]:checked').value;

    console.log(questionText, option1, option2, option3, option4, correctOption)

    // Validate that at least one option is selected
    if (!quizId || !questionText || !option1 || !option2 || !option3 || !option4 || !correctOption) {
        alert('All fields are required');
        return;
    }

    // Create the question object
    const questionData = {
        question_text: questionText,
        option1: option1,
        option2: option2,
        option3: option3,
        option4: option4,
        correct_option: parseInt(correctOption)
    };

    // Call the createQuestion function with the question data
    createQuestion(quizId, questionData);
});

// Function to create a question for a quiz
function createQuestion(quizId, questionData) {
    fetch(`https://localhost/quizzes/${quizId}/questions`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(questionData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Question created:', data);
        alert('Question created successfully!');
        showQuestions(quizId);  // Refresh the questions list after adding
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Failed to create question');
    });
}



// Function to edit a question
function editQuestion(quizId, questionId) {
    fetch(`https://localhost/questions/${questionId}`)
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to fetch question data');
        }
        return response.json();
    })
    .then(data => {
        // Update form fields with fetched question data
        console.log(JSON.stringify(data))
        document.getElementById('questionText').value = data.question_text;
        document.getElementById('option1').value = data.option1;
        document.getElementById('option2').value = data.option2;
        document.getElementById('option3').value = data.option3;
        document.getElementById('option4').value = data.option4;
        document.querySelector(`input[name="correctOption"][value="${data.correct_option}"]`).checked = true;
        
        // Change button text to 'Update Question'
        const addButton = document.querySelector('#createQuestionForm button[type="submit"]');
        addButton.textContent = 'Update Question';
        addButton.onclick = function(event) {
            updateQuestion(quizId, questionId);
            event.preventDefault(); // Prevent form submission
        };

        // Show the question section
        document.getElementById('quizSection').style.display = 'none';
        document.getElementById('questionSection').style.display = 'block';
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Failed to fetch question data');
    });
}



// Function to delete a question
function deleteQuestion(quizId, questionId) {
    if (confirm('Are you sure you want to delete this question?')) {
        fetch(`https://localhost/questions/${questionId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                alert('Question deleted successfully!');
                showQuestions(quizId);
            } else {
                alert('Failed to delete question');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('Error deleting question');
        });
    }
}

