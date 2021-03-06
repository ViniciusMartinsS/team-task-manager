package service

import (
	"fmt"
	"log"

	"github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type notificationService struct {
	userRepository contract.UserRepository
}

func NewNotificationService(userRepository contract.UserRepository) contract.NotificationService {
	return notificationService{userRepository}
}

func (n notificationService) Notify(task model.Task) bool {
	row, err := n.userRepository.FindBydId(task.UserId)
	if err != nil {
		log.Println("[ERROR] Gathering user to notify")
		return false
	}

	date := common.DateToStr(task.Performed)
	fmt.Printf("\n[NOTIFICATION] The tech %s performed the task %s on date %s\n", row.Name, task.Name, date)

	return true
}
