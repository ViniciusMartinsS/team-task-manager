package taskservice

import (
	"testing"

	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/ViniciusMartinsS/manager/internal/service"
	"github.com/ViniciusMartinsS/manager/test/mock/application"
	"github.com/ViniciusMartinsS/manager/test/mock/infrastructure"
	. "github.com/onsi/gomega"
)

const (
	TECHNICIAN_ID = iota + 1
	MANAGER_ID
)

func TestTaskServiceListMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	encryptionService := application.NewEncryptionServiceMock("")

	t.Run("Expect list task successfully to the technician role", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			List(TECHNICIAN_ID)

		g.Expect(task.Code).To(Equal(constant.SUCCESS_CODE))
		g.Expect(task.Result).To(BeAssignableToTypeOf(resultType))
		g.Expect(task.Message).To(BeEmpty())
	})

	t.Run("Expect list task successfully to the manager role", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			List(MANAGER_ID)

		g.Expect(task.Code).To(Equal(constant.SUCCESS_CODE))
		g.Expect(task.Result).To(BeAssignableToTypeOf(resultType))
		g.Expect(task.Message).To(BeEmpty())
	})

	t.Run("Expect not found message when there is no tasks", func(it *testing.T) {
		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, true)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			List(TECHNICIAN_ID)

		g.Expect(task.Result).To(BeEmpty())
		g.Expect(task.Code).To(Equal(constant.RECORD_NOT_FOUND_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.RECORD_NOT_FOUND_LIST_MESSAGE))
	})

	t.Run("Expect internal server error from task repository", func(it *testing.T) {
		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(true, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			List(MANAGER_ID)

		g.Expect(task.Result).To(BeEmpty())
		g.Expect(task.Code).To(Equal(constant.INTERNAL_SERVER_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.INTERNAL_SERVER_ERROR_MESSAGE))
	})

	t.Run("Expect internal server error from user repository", func(it *testing.T) {
		userRepository := infrastructure.NewuserRepositoryMock(true, true)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			List(TECHNICIAN_ID)

		g.Expect(task.Result).To(BeEmpty())
		g.Expect(task.Code).To(Equal(constant.INTERNAL_SERVER_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.INTERNAL_SERVER_ERROR_MESSAGE))
	})
}
