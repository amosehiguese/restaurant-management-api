# Restaurant Management API

The restaurant management API is a comprehensive solution designed to streamline and automate various aspects of restaurant operations. It provides a set of endpoints to manage menus, orders, reservations, tables, users, and more...

### Run the Server

To run the server, execute:

```bash
make start
```

### Libraries and features 
- Command Line flags with flag library
- Structured logging with zap
- Graceful Exits
- Chi as HTTP framework
- Custom Middlewares
- Pagination
- Documentation as Code with http-swagger


### API ENDPOINTS

`GET '/api/v1/menu'`

- returns all menu
- Request Arguments: None
- Returns: json

```json
{
  "1": "Science",
  "2": "Art",
  "3": "Geography",
  "4": "History",
  "5": "Entertainment",
  "6": "Sports"
}
```

`GET '/api/questions'`

- Fetches a list of dictionaries of questions, total number of questions, categories and the current category
- Request Arguments: `page` - integer
- Returns: A jsonified object with 10 paginated questions, total questions, a dictionary of categories and the current category string


```json
{
  "questions": [
    {
      "id": 15,
      "question": "The Taj Mahal is located in which Indian city?",
      "answer": "Agra",
      "difficulty": 2,
      "category": 3
    }
  ],
  "success": true,
  "total_questions": 18,
  "categories": {
    "1": "Science",
    "2": "Art",
    "3": "Geography",
    "4": "History",
    "5": "Entertainment",
    "6": "Sports"
  },
  "current_category": "All"
}
```
---

`GET '/api/categories/<int\:cat_id>questions'`
- Fetches all questions for a given category
- Request Arguments: `cat_id` - integer
- Returns: A jsonify object with questions corresponding to a given category Id, total questions, and current category string

```json
{
  "questions": [
    {
      "id": 18,
      "question": "How many paintings did Van Gogh sell in his lifetime?",
      "answer": "One",
      "difficulty": 4,
      "category": 2
    }
  ],
  "success": true,
  "total_questions": 100,
  "current_category": "Art"
}
```

---

`DELETE '/api/questions/<int\:id>'`

- Deletes a specified question corresponding to the id of the question
- Request Arguments: `id` - integer
- Returns: Does not return anything.

---

`POST '/api/quizzes'`

- Sends a post request to get next question
- Request Body:

```json
{
    'previous_questions': [16, 17, 18],
    'success': 'true'
 }
```

- Returns: a single new question object

```json
{
  "question": {
    "id": 17,
    "question": "La Giaconda is better known as what?",
    "answer": "Mona Lisa",
    "difficulty": 3,
    "category": 2
  }
}
```
---
`POST 'api/questions'`

- Sends a post request to add a new question
- Request Body:

```json
{
  "question": "Your question",
  "answer": "Your answer",
  "difficulty": "The difficulty level",
  "category": "The category"
}
```

- Returns: Does not return any new data

---

`POST 'api/questions/search'`

- Sends a post request for a specified search item
- Request Body:

```json
{
  "searchTerm": "What you searched for"
}
```

- Returns: any array of questions, a number of total questions that tallies with the search term and the current category string

```json
{
  "question": [ 
  {
    "id": 17,
    "question": "La Giaconda is better known as what?",
    "answer": "Mona Lisa",
    "difficulty": 3,
    "category": 2
  }
  ],
  "totalQuestions": 18,
  "currentCategory": "Art"
}
```

