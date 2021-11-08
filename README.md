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

  <p><b>Handle API Authentication</b></p>

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

  <p><b>Show all task of a technician</b></p>

  #### URL
  `/tasks`

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

  #### Try it out
  ```bash
  curl --location --request GET 'localhost:3000/tasks' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
  ```

</details>

<details>
  <summary><b>Create Task</b></summary>

  <p><b>Creates a task for a specific technician</b></p>

  #### URL
  `/tasks`

  #### Method
  `POST`

  #### Data Params
  ```json
  {
    "name": "Hello World",
    "summary": "Hello World! This is my new task",
    "performed": "07/11/2021"
  }
  ```

  * `name` **Required**
  * `summary` **Required**
  * `performed` **Optional - DD/MM/YYYY**

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  ```json
  {
    "status": 0,
    "result": [
        {
            "id": 2,
            "name": "Hello World",
            "summary": "Hello World! This is my new task",
            "Performed": "07/11/2021"
        }
     ]
  }
  ```

  #### Try it out
  ```bash
  curl --location --request POST 'localhost:3000/tasks' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "name": "Hello World",
      "summary": "Hello World! This is my new task",
      "performed": "07/11/2021"
  }'
  ```

</details>

<details>
  <summary><b>Update Task</b></summary>

  <p><b>Update a task of technician</b></p>

  #### URL
  `/tasks/:id`

  #### Method
  `PUT`

  #### Data Params
  ```json
  {
    "name": "Hello World",
    "performed": "07/11/2021"
  }
  ```

  * `name` **Optional**
  * `summary` **Optional**
  * `performed` **Optional - DD/MM/YYYY**

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
            "name": "Hello World",
            "Performed": "07/11/2021"
        }
     ]
  }
  ```

  #### Try it out
  ```bash
  curl --location --request PUT 'localhost:3000/tasks/1' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "performed": "18/11/2020"
  }'
  ```

</details>

<details>
  <summary><b>Delete Task</b></summary>

  <p><b>[MANAGER ONLY] Delete task of a technician</b></p>

  #### URL
  `/tasks/:id`

  #### Method
  `DELETE`

  #### Authorization
  `Bearer Token`

  * `token` **Required**

  #### Success Response
  ```json
  {
      "status": 0,
      "message": "Register with the following ID: '1' was deleted successfully!"
  }
  ```

  #### Error Response
  ```json
  {
      "status": 103,
      "message": "Hummmm... It seems you are not allowed to do such a thing. Ask for your manager help!"
  }
  ```

  #### Try it out
  ```bash
  curl --location --request DELETE 'localhost:3000/tasks/1' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
  ```

</details>



<!--
#### Postman Collection:
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/671b82a64b22a20a683e) -->
