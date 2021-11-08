package common

import (
	"encoding/json"
	"fmt"
	"regexp"

	constant "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/go-playground/validator/v10"
)

type TaskCreateDTO struct {
	Name      string  `validate:"required"`
	Summary   string  `validate:"required,max=2500"`
	Performed *string `json:",omitempty"`
}

type TaskUpdateDTO struct {
	Name      string  `json:",omitempty"`
	Summary   string  `json:",omitempty" validate:"max=2500"`
	Performed *string `json:",omitempty"`
}

var validate *validator.Validate

func ValidateLoginSchema(body []byte) error {
	var login model.LoginPayload

	err := json.Unmarshal(body, &login)
	if err != nil {
		return err
	}

	validate = validator.New()

	err = validate.Struct(login)
	if err == nil {
		return nil
	}

	return err
}

func ValidateTaskCreateSchema(body []byte) error {
	var task TaskCreateDTO
	err := json.Unmarshal(body, &task)
	if err != nil {
		return err
	}

	validate = validator.New()

	err = validate.Struct(task)
	if err != nil {
		return err
	}

	err = ValidateDateFormat(task.Performed)
	if err != nil {
		return err
	}

	return nil
}

func ValidateTaskUpdateSchema(body []byte) error {
	var task TaskUpdateDTO
	err := json.Unmarshal(body, &task)
	if err != nil {
		return err
	}

	validate = validator.New()

	err = validate.Struct(task)
	if err != nil {
		return err
	}

	err = ValidateDateFormat(task.Performed)
	if err != nil {
		return err
	}

	return nil
}

func ValidateDateFormat(performed *string) error {
	if performed == nil {
		return nil
	}

	matched, _ := regexp.MatchString(constant.DATE_REGEX, *performed)
	if matched {
		return nil
	}

	err := fmt.Errorf(constant.DATE_BAD_REQUEST)
	return err
}
