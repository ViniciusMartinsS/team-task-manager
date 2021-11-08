package contract

import "github.com/ViniciusMartinsS/manager/internal/domain"

type NotificationService interface {
	Notify(task domain.Task)
}
