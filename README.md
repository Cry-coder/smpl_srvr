## Helpdesk API
Here we've created helpdesk API for IT department where staff will be able to put theirs problems or questions forwarded for IT department. 

### Helpdesk app logic

First time running server we need connect to postgresql database and create needed tables and relations. For this propose run:
`go run main.go -migrate -dsn "postgres://{user}:{Password}@localhost/{databaseName}?sslmode=disable" -path "{Exect path to tables.sql}"`


After launching server you need to create admin role account using post request.
This could be happened once or after cleaning database.  

POST: `http://localhost:8007/v1/admin/create/admin`

 `{
"Fn": "Firstname",
"Ln": "Lastname",
"Email": "example@email.com",
"Password": "strongpassword"
}`

Logging for admin and users processed at: 

PUT: `http://localhost:8007/v1/login`

After creating admin account and logging in we need to create profiles for our staff. Creating this type of accounts could be created only with admin account at: 

POST: `http://localhost:8007/v1/admin/createuser`

`{
"Fn": "Firstname",
"Ln": "Lastname",
"Email": "example1@email.com",
"Password": "strongpassword"
}`

## With admin account we able:
- Login to the system

PUT: `http://localhost:8007/v1/login`

- View account profile

GET: `http://localhost:8007/v1/admin/profile`
- View all users profiles

GET: `http://localhost:8007/v1/admin/all`

- One user profile


GET: `http://localhost:8007/v1/admin/one/{id}`

- All questions


GET: `http://localhost:8007/v1/admin/questions`

- One question

GET: `http://localhost:8007/v1/admin/question/{id}`

- Delete user by ID


DELETE: `http://localhost:8007/v1/admin/delete/{id}`
- Delete question by ID


DELETE: `http://localhost:8007/v1/admin/delete/question/{id}`

- Create user account

POST: `http://localhost:8007/v1/admin/createuser`

- Update question

PUT: `http://localhost:8007/v1/admin/update/question`

`{
"QuestionId": {id},
"Status": {bool}
}`
- Logout from the system

GET: `http://localhost:8007/v1/admin/logout`


## After users accounts were created by admin account they are able:
- Login to the system

PUT: `http://localhost:8007/v1/login`

- View account profile

GET: `http://localhost:8007/v1/user/profile`

- Update firstname, lastname and password in theirs account

PUT: `http://localhost:8007/v1/user/update/profile`

- Create questions

POST: `http://localhost:8007/v1/user/cr`
`{
"Question": "Question string"
}`
- Update questions

PUT: `http://localhost:8007/v1/users/update/question`

`{ "QuestionId": {id},
"Question": "Changed question" }`
- Find questions by id

GET: `http://localhost:8007/v1/users/question/{id}`
- Logout from the system

GET: `http://localhost:8007/v1/user/logout`





