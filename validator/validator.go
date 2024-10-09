package validator

import (
	"errors"
	"github.com/bminh1762000/jwt-auth-go/models"
)

func ValidateCreateTaskInput(task models.TaskInput) error {
	if task.DueDate == nil && task.Title == nil {
		return errors.New("due_date and title are required")
	}

	return nil
}

func ValidateCreateUserInput(user models.UserInput) error {
	if user.Username == nil && user.Password == nil {
		return errors.New("username and password are required")
	}

	return nil
}

func ValidateLoginInput(user models.UserInput) error {
	if user.Username == nil && user.Password == nil {
		return errors.New("username and password are required")
	}

	return nil
}
