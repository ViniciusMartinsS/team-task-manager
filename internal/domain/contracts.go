package domain

type Database interface {
	CreateDatabase() error
	CreateTables() error
	SeedTables()
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

type NotificationService interface {
	Notify(task Task)
}

type EncryptionService interface {
	Encrypt(content string) string
	Decrypt(contentEncrypted string) string
}
