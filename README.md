# Task Manager
This service aims to manage tasks performed during a working day.

## Getting Started

#### Clone Repository
```
$ git clone https://github.com/ViniciusMartinsS/team-task-manager.git
```

#### Running on docker
```
$ make docker
```

#### Running on locally
```
$ make setup
$ make run
```

## API Documentation

<details>
  <summary><b>Authentication</b></summary>

<p><b>Returns JSON data with the API access token</b></p>

#### URL
`/auth/login`

#### Method
`POST`

#### Data Params
```json
{
    "email": "example@example.io",
    "password": "123456"
}
```

* `email` **Required**
* `password` **Required**

#### Success Response
```json
{
    "status": 0,
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
}
```

#### Error Response
```json
{
    "status": 101,
    "message": "Stop right there! You are unauthorized!"
}
```

#### Try it out
```bash
curl --location --request POST 'localhost:3000/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "example@example.io",
    "password": "123456"
}'
```

</details>

<details>
  <summary><b>List Tasks</b></summary>

  <p><b>Returns JSON data with the created tasks</b></p>

  #### URL
  `/auth/login`

  #### Method
  `GET`

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  ```json
  {
      "status": 0,
      "result": [
          {
              "id": 1,
              "name": "Task Hello World",
              "summary": "Hello World! This is my new task"
          }
      ]
  }
  ```

</details>



<!--
#### Postman Collection:
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/671b82a64b22a20a683e) -->
