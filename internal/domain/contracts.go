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
	Login([]byte) LoginResponse
}

type AuthService interface {
	Login(email, password string) LoginResponse
}

type TaskController interface {
	List(HandleTaskRequest) TaskResponse
	Create(HandleTaskRequest) TaskResponse
	Update(HandleTaskRequest) TaskResponse
	Delete(HandleTaskRequest) TaskResponse
}

type TaskService interface {
	List(userId int) TaskResponse
	Create(userId int, payload TaskPayload) TaskResponse
	Update(id int, userId int, payload TaskPayload) TaskResponse
	Delete(id int, userId int) TaskResponse
}

type NotificationService interface {
	Notify(task Task)
}

type EncryptionService interface {
	Encrypt(content string) string
	Decrypt(contentEncrypted string) string
}
