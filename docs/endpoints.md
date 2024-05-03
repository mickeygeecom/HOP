# LootLocker Quiz API Endpoints

# Admin Endpoints

### Quizzes

- **GET /api/admin/quizzes**
  - Description: Retrieve a list of all quizzes.
  - Method: `GET`
  - Response:
    - Status: `200 OK`
    - Body: Array of quiz objects.

- **POST /api/admin/quizzes**
  - Description: Create a new quiz.
  - Method: `POST`
  - Request Body: JSON object containing quiz details (name, description).
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing the ID of the newly created quiz.

- **GET /api/admin/quizzes/{id}**
  - Description: Retrieve details of a specific quiz.
  - Method: `GET`
  - Path Parameter: `id` (Quiz ID)
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing quiz details.

- **PUT /api/admin/quizzes/{id}**
  - Description: Update details of a specific quiz.
  - Method: `PUT`
  - Path Parameter: `id` (Quiz ID)
  - Request Body: JSON object containing updated quiz details. (name, description)
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing updated quiz details.

- **DELETE /api/admin/quizzes/{id}**
  - Description: Delete a specific quiz.
  - Method: `DELETE`
  - Path Parameter: `id` (Quiz ID)
  - Response:
    - Status: `200 OK`
    - Body: Empty response body.

### Questions

- **POST /api/admin/quizzes/{quizId}/questions**
  - Description: Create a new question for a specific quiz.
  - Method: `POST`
  - Path Parameter: `quizId` (Quiz ID)
  - Request Body: JSON object containing question details. (question_text, option1, option2, option3, option4, correct_option)
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing the ID of the newly created question.

- **GET /api/admin/questions/{questionId}**
  - Description: Retrieve details of a specific question.
  - Method: `GET`
  - Path Parameter: `questionId` (Question ID)
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing question details. (question_text, option1, option2, option3, option4, correct_option)

- **GET /api/admin/quizzes/{quizId}/questions**
  - Description: Retrieve a list of all questions for a specific quiz.
  - Method: `GET`
  - Path Parameter: `quizId` (Quiz ID)
  - Response:
    - Status: `200 OK`
    - Body: Array of question objects.

- **PUT /api/admin/questions/{questionId}**
  - Description: Update details of a specific question.
  - Method: `PUT`
  - Path Parameter: `questionId` (Question ID)
  - Request Body: JSON object containing updated question details.
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing updated question details.

- **DELETE /api/admin/questions/{questionId}**
  - Description: Delete a specific question.
  - Method: `DELETE`
  - Path Parameter: `questionId` (Question ID)
  - Response:
    - Status: `200 OK`
    - Body: Empty response body.

# User Endpoints

- **GET /api/user/quizzes**
  - Description: Retrieve a list of quizzes available to the end user.
  - Method: `GET`
  - Response:
    - Status: `200 OK`
    - Body: Array of quiz objects.

- **GET /api/user/quizzes/{id}**
  - Description: Retrieve details of a specific quiz available to the end user.
  - Method: `GET`
  - Path Parameter: `id` (Quiz ID)
  - Response:
    - Status: `200 OK`
    - Body: JSON object containing quiz details.

- **GET /api/user/questions/{quizId}**
  - Description: Retrieve a list of questions for a specific quiz available to the end user.
  - Method: `GET`
  - Path Parameter: `quizId` (Quiz ID)
  - Response:
    - Status: `200 OK`
    - Body: Array of question objects.

- **POST /api/user/quizzes/{quizId}/submit**
  - Description: Submit answers for a specific quiz.
  - Method: `POST`
  - Path Parameter: `quizId` (Quiz ID)
  - Request Body: JSON object containing submitted answers.
  - Response:
    - Status: `200 OK`
    - Body: Result of the quiz submission.