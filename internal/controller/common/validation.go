package common

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type TaskCreate struct {
	Name      string  `validate:"required"`
	Summary   string  `validate:"required,max=2500"`
	Performed *string `json:",omitempty" validate:"len=10"`
}

type TaskUpdate struct {
	Name      string  `json:",omitempty"`
	Summary   string  `json:",omitempty" validate:"max=2500"`
	Performed *string `json:",omitempty" validate:"len=10"`
}

var validate *validator.Validate

const REGEX = `^([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}$`

func ValidateLoginSchema(body []byte) error {
	var login Login
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
	_ = json.Unmarshal(body, &task)

	validate = validator.New()

	err := validate.Struct(task)
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
	_ = json.Unmarshal(body, &task)

	validate = validator.New()

	err := validate.Struct(task)
	fmt.Println(err)
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

	matched, _ := regexp.MatchString(REGEX, *performed)
	if matched {
		return nil
	}

	err := fmt.Errorf("key: 'TaskCreate.Summary' Error:Field validation for 'Performed' failed on the 'format' regex")
	return err
}
