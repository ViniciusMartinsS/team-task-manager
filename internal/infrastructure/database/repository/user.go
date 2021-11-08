package repository

import (
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) contract.UserRepository {
	return userRepository{conn}
}

func (u userRepository) FindBydId(id int) (model.User, error) {
	var user model.User

	result := u.conn.
		Preload("Role").
		First(&user, id)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (u userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User

	result := u.conn.
		Preload("Role").
		Select("email", "password", "id").
		Where("email = ?", email).
		First(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
