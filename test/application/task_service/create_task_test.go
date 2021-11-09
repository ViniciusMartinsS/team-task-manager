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

func TestTaskServiceCreateMethod(t *testing.T) {
	g := NewGomegaWithT(t)

	encryptionService := application.NewEncryptionServiceMock("")
	payload := model.TaskPayload{Name: "Example", Summary: "Task Example", Performed: "09/11/2021"}

	t.Run("Expect create a task successfully", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(false, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Create(TECHNICIAN_ID, payload)

		g.Expect(task.Code).To(Equal(constant.SUCCESS_CODE))
		g.Expect(task.Result).To(BeAssignableToTypeOf(resultType))
		g.Expect(task.Message).To(BeEmpty())
	})

	t.Run("Expect internal server error", func(it *testing.T) {
		var resultType []model.TaskResponseContent

		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		taskRepository := infrastructure.NewTaskRepositoryMock(true, false)
		notificationService := application.NewNotificationServiceMock(userRepository)

		task := service.
			NewTaskService(taskRepository, userRepository, notificationService, encryptionService).
			Create(TECHNICIAN_ID, payload)

		g.Expect(task.Result).To(Equal(resultType))
		g.Expect(task.Code).To(Equal(constant.INTERNAL_SERVER_ERROR_CODE))
		g.Expect(task.Message).To(Equal(constant.INTERNAL_SERVER_ERROR_MESSAGE))
	})
}
