package repository

import (
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) contract.UserRepository {
	return userRepository{conn}
}

func (u userRepository) FindBydId(id int) (domain.User, error) {
	var user domain.User

	result := u.conn.
		Preload("Role").
		First(&user, id)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return user, nil
}

func (u userRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	result := u.conn.
		Preload("Role").
		Select("email", "password", "id").
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return user, nil
}
