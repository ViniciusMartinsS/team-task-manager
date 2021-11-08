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
```bash
$ make setup
$ make run
```

## API Documentation

#### Authentication
Returns JSON data with the API access token

#### URL
/auth/login

#### Method
`POST`

#### Data Params
```json
{
    "email": "example@example.io",
    "password": "123456"
}
```

<details>
  <summary><b>Authentication</b></summary>
  > Returns JSON data with the API access token
</details>
