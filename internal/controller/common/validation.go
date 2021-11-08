package common

import (
	"encoding/json"
	"fmt"
	"regexp"

	constants "github.com/ViniciusMartinsS/manager/internal/common"
	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/go-playground/validator/v10"
)

type TaskCreate struct {
	Name      string  `validate:"required"`
	Summary   string  `validate:"required,max=2500"`
	Performed *string `json:",omitempty"`
}

type TaskUpdate struct {
	Name      string  `json:",omitempty"`
	Summary   string  `json:",omitempty" validate:"max=2500"`
	Performed *string `json:",omitempty"`
}

var validate *validator.Validate

func ValidateLoginSchema(body []byte) error {
	var login domain.LoginPayload
	_ = json.Unmarshal(body, &login)

	validate = validator.New()

	err := validate.Struct(login)
	if err == nil {
		return nil
	}

	return err
}

func ValidateTaskCreateSchema(body []byte) error {
	var task TaskCreate
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
	var task TaskUpdate
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

	matched, _ := regexp.MatchString(constants.DATE_REGEX, *performed)
	if matched {
		return nil
	}

	err := fmt.Errorf(constants.DATE_BAD_REQUEST)
	return err
}
