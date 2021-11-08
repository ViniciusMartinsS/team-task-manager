package contract

import "github.com/ViniciusMartinsS/manager/internal/domain"

type UserRepository interface {
	FindBydId(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}
