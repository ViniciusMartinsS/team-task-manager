package domain

type Database interface {
	CreateDatabase() error
	CreateTables() error
	SeedTables()
}

type TaskRepository interface {
	FindAll() ([]Task, error)
	FindByUserId(id int) ([]Task, error)
	Create(task Task) (Task, error)
	Update(id int, userId int, task Task) (Task, error)
	Delete(id int) (bool, error)
}

type UserRepository interface {
	FindBydId(id int) (User, error)
	FindByEmail(email string) (User, error)
}

type AuthController interface {
	Login([]byte) (LoginResponse, int)
}

type AuthService interface {
	Login(email, password string) (LoginResponse, int)
}

type TaskController interface {
	List(HandleTaskRequest) (TaskResponse, int)
	Create(HandleTaskRequest) (TaskResponse, int)
	Update(HandleTaskRequest) (TaskResponse, int)
	Delete(HandleTaskRequest) (TaskResponse, int)
}

type TaskService interface {
	List(userId int) (TaskResponse, int)
	Create(userId int, payload TaskPayload) (TaskResponse, int)
	Update(id int, userId int, payload TaskPayload) (TaskResponse, int)
	Delete(id int, userId int) (TaskResponse, int)
}

type NotificationService interface {
	Notify(task Task)
}

type EncryptionService interface {
	Encrypt(content string) string
	Decrypt(contentEncrypted string) string
}
