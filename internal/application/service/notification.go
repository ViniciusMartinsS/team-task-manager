package service

import (
	"fmt"
	"log"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/helper"
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

	date := helper.DateToStr(task.Performed)
	fmt.Printf("\n[NOTIFICATION] The tech %s performed the task %s on date %s\n", row.Name, task.Name, date)
}
