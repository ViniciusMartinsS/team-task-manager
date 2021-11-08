package internal

import (
	"github.com/ViniciusMartinsS/manager/internal/application/service"
	"github.com/ViniciusMartinsS/manager/internal/controller"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database/repository"
	"github.com/golobby/container/v3"
)

func InitializeDIContainers() {
	conn := database.Connection()

	container.Singleton(func() contract.UserRepository {
		return repository.NewUserRepository(conn)
	})
	container.Singleton(func() contract.TaskRepository {
		return repository.NewTaskRepository(conn)
	})

	container.Singleton(func() contract.EncryptionService {
		encryptionKey := infrastructure.GetConfig("encryption_key")
		return service.NewEncryptionService(encryptionKey)
	})
	container.Singleton(service.NewAuthService)
	container.Singleton(service.NewNotificationService)
	container.Singleton(service.NewTaskService)

	container.Singleton(controller.NewAuthController)
	container.Singleton(controller.NewTaskController)
}
