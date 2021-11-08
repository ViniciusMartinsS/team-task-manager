package contract

type EncryptionService interface {
	Encrypt(content string) string
	Decrypt(contentEncrypted string) string
}
