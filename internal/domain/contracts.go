package domain

type Database interface {
	CreateDatabase() error
	CreateTables() error
	SeedTables()
}

type EncryptionService interface {
	Encrypt(content string) string
	Decrypt(contentEncrypted string) string
}
