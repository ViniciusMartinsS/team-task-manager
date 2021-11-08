package internal

import (
	"github.com/ViniciusMartinsS/manager/internal/application/service"
	"github.com/ViniciusMartinsS/manager/internal/controller"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database/repository"
	"github.com/golobby/container/v3"
)

func InitializeDIContainers() {
	conn := database.Connection()

	container.Singleton(func() domain.UserRepository {
		return repository.NewUserRepository(conn)
	})
	container.Singleton(func() domain.TaskRepository {
		return repository.NewTaskRepository(conn)
	})

	container.Singleton(service.NewEncryption)
	container.Singleton(service.NewAuthService)
	container.Singleton(service.NewNotificationService)
	container.Singleton(service.NewTaskService)

	container.Singleton(controller.NewAuthController)
	container.Singleton(controller.NewTaskController)
}
