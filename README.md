# Task Manager
This service aims to manage tasks performed during a working day.

## API Documentation

<details>
  <summary><b>Authentication</b></summary>
  
  </br>
  
  > **Handle API Authentication**

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
  
  <sub>
  
  **⚠️ Credentials**
  
  ```
  Technician 01 
   Email: john.doe@hotmail.com 
   Password: Sw@rd2021
  
  Technician 02
   Email: john.doe@hotmail.com
   Password: Sw@rd2021
  
  Manager 
    Email: john.doe@hotmail.com 
    Password: Sw@rd2021
  ```
  
  </sup>
    
</details>

<details>
  <summary><b>List Tasks</b></summary>
  
  </br>
  
  > **Show all task of a technician**

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

  </br>
  
  > **Creates a task for a specific technician**

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
  
  </br>
  
  > **Update a task of technician**

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

  </br>
  
  > **[MANAGER ONLY] Delete task of a technician**

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

### Run In Postman
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/955776fb0c44d8d9235f)

## Developer Guideline

- Clone Repository
```
$ git clone https://github.com/ViniciusMartinsS/team-task-manager.git
```

### Useful Commands

- Run Application On Docker
```
$ make docker
```

- Run Application Locally

<sub>⚠️ Before start, on the `config` directory, you must create your `app.json` in order to have the environment variables. 
On the root of the project, execute the following command:</sup>
```
$ cp config/app-dist.json config/app.json
```

<sub> Having the config setup, now you can execute the following commands: <sub>

```
$ make setup
$ make run
```

- Run Database Migration
```
$ make migrate
```

- Run Lint Checker
```
make lint
```

- Run Tests
```bash
make tests
```

- Generate Test Coverage File
```bash
make coverage
```
