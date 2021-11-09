package application

import "github.com/ViniciusMartinsS/manager/internal/domain/contract"

type encryptionRepositoryMock struct {
	keyString string
}

func NewEncryptionServiceMock(keyString string) contract.EncryptionService {
	return encryptionRepositoryMock{keyString}
}

func (e encryptionRepositoryMock) Encrypt(content string) string {
	return ""
}

func (e encryptionRepositoryMock) Decrypt(contentEncrypted string) string {
	return ""
}
