package taskservice

import (
	"fmt"
	"testing"

	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/ViniciusMartinsS/manager/internal/service"
	"github.com/ViniciusMartinsS/manager/test/mock/application"
	"github.com/ViniciusMartinsS/manager/test/mock/infrastructure"
	. "github.com/onsi/gomega"
)

func TestTaskServiceDeleteMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	encryptionService := application.NewEncryptionServiceMock("")
	// payload := model.TaskPayload{Name: "Example", Summary: "Task Example", Performed: "09/11/2021"}

	t.Run("Expect update a task successfully", func(it *testing.T) {
		var resultType []model.TaskResponseContent
		successMessage := fmt.Sprintf(constant.SUCCESS_DELETE_MESSAGE, 1)

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Delete(1, MANAGER_ID)

		g.Expect(task.Code).To(Equal(constant.SUCCESS_CODE))
		g.Expect(task.Message).To(Equal(successMessage))
		g.Expect(task.Result).To(Equal(resultType))
	})

	t.Run("Expect error when user is not a technician", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Delete(1, TECHNICIAN_ID)

		g.Expect(task.Result).To(Equal(resultType))
		g.Expect(task.Code).To(Equal(constant.FORBIDDEN_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.FORBIDDEN_ERROR_MESSAGE))
	})

	t.Run("Expect not found error when there is no record for the id provided", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, true)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Delete(1, MANAGER_ID)

		g.Expect(task.Result).To(Equal(resultType))
		g.Expect(task.Code).To(Equal(constant.RECORD_NOT_FOUND_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.RECORD_NOT_FOUND_ERROR_MESSAGE))
	})

	t.Run("Expect internal server error", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(true, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Delete(1, MANAGER_ID)

		g.Expect(task.Result).To(Equal(resultType))
		g.Expect(task.Code).To(Equal(constant.INTERNAL_SERVER_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.INTERNAL_SERVER_ERROR_MESSAGE))
	})

	t.Run("Expect internal server error caused by user repository", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(true, true)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Delete(1, MANAGER_ID)

		g.Expect(task.Result).To(Equal(resultType))
		g.Expect(task.Code).To(Equal(constant.INTERNAL_SERVER_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.INTERNAL_SERVER_ERROR_MESSAGE))
	})
}
