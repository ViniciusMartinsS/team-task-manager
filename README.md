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

  ```json
  {
    "status": 100,
    "message": "Key: 'LoginPayload.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'LoginPayload.Password' Error:Field validation for 'Password' failed on the 'required' tag"
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
   Email: john.doe@example.com
   Password: Sw@rd2021

  Technician 02
   Email: julius.rock@example.com"
   Password: Sw@rd2021

  Manager
    Email: jane.doe@example.com
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

  #### Error Response
  ```json
  {
    "status": 104,
    "message": "You do not have any tasks. Create a new one & let's get to work! ;)"
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

  #### Error Response
  ```json
  {
    "status": 100,
    "message": "Key: 'TaskCreateDTO.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'TaskCreateDTO.Summary' Error:Field validation for 'Summary' failed on the 'required' tag"
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
            "name": "Task Hello World",
            "summary": "Hello World! This is my new task",
            "performed": "20/11/2020"
        }
     ]
  }
  ```

  #### Error Response
  ```json
  {
    "status": 100,
    "message": "key: 'TaskCreate.Summary' Error:Field validation for 'Performed' failed on the 'format' regex"
  }
  ```

  ```json
  {
    "status": 104,
    "message": "Hmmmm... We could not find the requested record. Are you sure it exists? Are you sure it belongs to you?"
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

  ```json
  {
    "status": 104,
    "message": "Hmmmm... We could not find the requested record. Are you sure it exists? Are you sure it belongs to you?"
  }
  ```

  #### Try it out
  ```bash
  curl --location --request DELETE 'localhost:3000/tasks/1' \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
  ```

</details>

<details>
  <summary><b>API Erros Description</b></summary>

  </br>

  > **Here you are going to find out the description of our API errors**

  ```
  Code: 0
  Message: Success - Returns Created Object
  HTTP Status: 200

  Code: 100
  Message: Bad Request - Returns Invalid Fields
  HTTP Status: 400

  Code: 101
  Message: Stop Right There! You Are Unauthorized!
  HTTP Status: 401

  Code: 103
  Message: Hmmmm... It seems you are not allowed to do such a thing. Ask for your manager help!
  HTTP Status: 403

  Code: 104
  Message (UPDATE/DELETE): Hmmmm... We could not find the requested record. Are you sure it exists? Are you sure it belongs to you?
  Message (LIST): You do not have any tasks. Create a new one & let's get to work! ;
  HTTP Status: 404

  Code: 199
  Message: Something is broken on our side :(. Sorry for the inconvenience!
  HTTP Status: 500

  ```

</details>

### Run In Postman
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/955776fb0c44d8d9235f)

## Developer Guideline

**- Clone Repository**
```
$ git clone https://github.com/ViniciusMartinsS/team-task-manager.git
```

### Useful Commands

**- Run Application On Docker**
```
$ make docker
```

**- Run Application Locally**

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

**- Run Database Migration**
```
$ make migrate
```

**- Run Lint Checker**
```
$ make lint
```

**- Run Tests**
```bash
$ make tests
```

**- Generate Test Coverage File**
```bash
$ make coverage
```
