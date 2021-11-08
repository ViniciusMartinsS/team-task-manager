package contract

import "github.com/ViniciusMartinsS/manager/internal/domain/model"

type UserRepository interface {
	FindBydId(id int) (model.User, error)
	FindByEmail(email string) (model.User, error)
}
