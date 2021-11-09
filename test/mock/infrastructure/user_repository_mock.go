package infrastructure

import (
	"fmt"

	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
)

type userRepositoryMock struct {
	shouldFail bool
	internal   bool
}

func NewuserRepositoryMock(shouldFail, internal bool) contract.UserRepository {
	return &userRepositoryMock{shouldFail, internal}
}

func (u userRepositoryMock) FindBydId(id int) (model.User, error) {
	if u.shouldFail && !u.internal {
		return model.User{}, fmt.Errorf("record not found")
	}

	if u.shouldFail && u.internal {
		return model.User{}, fmt.Errorf("internal server error")
	}

	return model.User{}, nil
}

func (u userRepositoryMock) FindByEmail(email string) (model.User, error) {
	if u.shouldFail && !u.internal {
		return model.User{}, fmt.Errorf("record not found")
	}

	if u.shouldFail && u.internal {
		return model.User{}, fmt.Errorf("internal server error")
	}

	return model.User{
		Email:    "visimonemartins@hotmail.com",
		Password: "$2a$12$bvFmmw3X4ctLBf39rF.DqOtp98WPFhm/tztUAXvWpLEgtCpohWzQW",
	}, nil
}
