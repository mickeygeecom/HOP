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
    fetch('/api/admin/quizzes', {
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
        alert('Failed to create quiz', 'error');
    });
}

// Function to list quizzes
function listQuizzes() {
    // Change button text to 'Add Question' under questions section
    const addButton = document.querySelector('#createQuestionForm button[type="submit"]');
    addButton.textContent = 'Add Question';

    fetch('/api/admin/quizzes')
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
                    <button onclick="showQuestions(${quiz.id})" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 rounded">Questions</button>
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
        alert('Failed to fetch quizzes', 'error');
    });
}
listQuizzes();  // List quizzes when the page loads

// Function to edit a quiz
function editQuiz(quizId) {
    const newName = prompt('Enter new name for the quiz:');
    const newDescription = prompt('Enter new description for the quiz:');
    if (!newName || !newDescription) {
        alert('⚠️ Name and description are required', 'error');
        return;
    }
    fetch(`/api/admin/quizzes/${quizId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: newName, description: newDescription })
    })
    .then(response => {
        if (response.ok) {
            alert('Quiz updated successfully');
            listQuizzes();
        } else {
            alert('Failed to update quiz', 'error');
        }
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Error updating quiz', 'error');
    });
}

// Function to delete a quiz
function deleteQuiz(quizId) {
    if (confirm('Are you sure you want to delete this quiz?')) {
        fetch(`/api/admin/quizzes/${quizId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                alert('Quiz deleted successfully!');
                listQuizzes();
            } else {
                alert('Failed to delete quiz', 'error');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('Error deleting quiz', 'error');
        });
    }
}

// #### QUESTIONS ####

// Function to show questions for a selected quiz
function showQuestions(quizId) {
    currentQuizId = quizId;
    console.log("Setting current quiz id to", quizId);

    fetch(`/api/admin/quizzes/${quizId}/questions`)
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
            let i = 0;
            data.forEach(question => {
                i++;
                const item = document.createElement('li');
                item.className = 'flex justify-between items-center mb-4';
                item.innerHTML = `
                    <span class="flex-1">${i}: ${question.question_text}</span>
                    <div class="flex justify-end items-center gap-2">
                        <button onclick="editQuestion(${quizId}, ${question.id})" class="bg-orange-500 hover:bg-orange-700 text-white font-bold py-1 px-2 rounded">Edit</button>
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
        alert('Failed to fetch questions', 'error');
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
        alert('All fields are required', 'error');
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
    fetch(`/api/admin/quizzes/${quizId}/questions`, {
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
        alert('Failed to create question', 'error');
    });
}



// Function to edit a question
function editQuestion(quizId, questionId) {
    fetch(`/api/admin/questions/${questionId}`)
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
        alert('Failed to fetch question data', 'error');
    });
}

function updateQuestion(quizId, questionId) {
    // Get the updated question data from the form
    const updatedQuestion = {
        question_text: document.getElementById('questionText').value,
        option1: document.getElementById('option1').value,
        option2: document.getElementById('option2').value,
        option3: document.getElementById('option3').value,
        option4: document.getElementById('option4').value,
        correct_option: parseInt(document.querySelector('input[name="correctOption"]:checked').value)
    };

    // Send the updated question data to the backend to update the question
    fetch(`/api/admin/questions/${questionId}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(updatedQuestion)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to update question');
        }
        return response.json();
    })
    .then(data => {
        // Handle successful update (e.g., show success message, navigate back to quiz section, etc.)
        console.log('Question updated successfully:', data);
        // Update the question list after updating
        showQuestions(quizId);
    })
    .catch(error => {
        console.error('Error updating question:', error);
        alert('Failed to update question', 'error');
    });
}




// Function to delete a question
function deleteQuestion(quizId, questionId) {
    if (confirm('Are you sure you want to delete this question?')) {
        fetch(`/api/admin/questions/${questionId}`, {
            method: 'DELETE'
        })
        .then(response => {
            if (response.ok) {
                //alert('Question deleted successfully!');
                showQuestions(quizId);
            } else {
                alert('Failed to delete question', 'error');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert('Error deleting question', 'error');
        });
    }
}

