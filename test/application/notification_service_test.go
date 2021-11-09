package application

import (
	"testing"
	"time"

	"github.com/ViniciusMartinsS/manager/internal/application/service"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/ViniciusMartinsS/manager/test/mock/infrastructure"
	. "github.com/onsi/gomega"
)

func TestNotificationService(t *testing.T) {
	g := NewGomegaWithT(t)

	performed := time.Now()
	params := model.Task{UserId: 1, Performed: &performed}

	t.Run("Expect notify successfully", func(it *testing.T) {
		userRepository := infrastructure.NewuserRepositoryMock(false, false)
		notified := service.NewNotificationService(userRepository).
			Notify(params)

		g.Expect(notified).To(BeTrue())
	})

	t.Run("Expect not notify", func(it *testing.T) {
		userRepository := infrastructure.NewuserRepositoryMock(true, true)
		notified := service.NewNotificationService(userRepository).
			Notify(params)

		g.Expect(notified).To(BeFalse())
	})
}
