package internal

import (
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database"
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/database/repository"
	"github.com/ViniciusMartinsS/manager/internal/service"
	"github.com/ViniciusMartinsS/manager/internal/usecases"
	"github.com/golobby/container/v3"
)

func InitializeDIContainers() {
	conn := database.Connection()

	err := container.Singleton(func() contract.UserRepository {
		return repository.NewUserRepository(conn)
	})
	if err != nil {
		panic("[DI_CONTAINER] failed to setup user repository")
	}

	err = container.Singleton(func() contract.TaskRepository {
		return repository.NewTaskRepository(conn)
	})
	if err != nil {
		panic("[DI_CONTAINER] failed to setup task repository")
	}

	err = container.Singleton(func() contract.EncryptionService {
		encryptionKey := infrastructure.GetConfig("encryption_key")
		return service.NewEncryptionService(encryptionKey)
	})
	if err != nil {
		panic("[DI_CONTAINER] failed to setup encryption_key env")
	}

	err = container.Singleton(service.NewAuthService)
	if err != nil {
		panic("[DI_CONTAINER] failed to setup auth service")
	}

	err = container.Singleton(service.NewNotificationService)
	if err != nil {
		panic("[DI_CONTAINER] failed to setup notification service")
	}

	err = container.Singleton(service.NewTaskService)
	if err != nil {
		panic("[DI_CONTAINER] failed to setup task service")
	}

	err = container.Singleton(usecases.NewAuthUseCases)
	if err != nil {
		panic("[DI_CONTAINER] failed to setup auth usecases")
	}

	err = container.Singleton(usecases.NewTaskUseCases)
	if err != nil {
		panic("[DI_CONTAINER] failed to setup task usecases")
	}
}
