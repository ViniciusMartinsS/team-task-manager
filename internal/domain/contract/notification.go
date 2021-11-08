package contract

import "github.com/ViniciusMartinsS/manager/internal/domain/model"

type NotificationService interface {
	Notify(task model.Task)
}
