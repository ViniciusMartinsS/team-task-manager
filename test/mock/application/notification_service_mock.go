package application

import (
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type notificationRepositoryMock struct {
	userRepository contract.UserRepository
}

func NewNotificationServiceMock(userRepository contract.UserRepository) contract.NotificationService {
	return notificationRepositoryMock{userRepository}
}

func (n notificationRepositoryMock) Notify(task model.Task) bool {
	return true
}
