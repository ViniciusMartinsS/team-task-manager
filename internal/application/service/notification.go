package service

import (
	"fmt"
	"log"

	"github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
)

type notificationService struct {
	userRepository domain.UserRepository
}

func NewNotificationService(userRepository domain.UserRepository) domain.NotificationService {
	return notificationService{userRepository}
}

func (n notificationService) Notify(task domain.Task) {
	row, err := n.userRepository.FindBydId(task.UserId)
	if err != nil {
		log.Println("[ERROR] Gathering user to notify")
		return
	}

	date := common.DateToStr(task.Performed)
	fmt.Printf("\n[NOTIFICATION] The tech %s performed the task %s on date %s\n", row.Name, task.Name, date)
}
