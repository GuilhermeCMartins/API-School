package utils

import (
	"api-school/models"

	"gopkg.in/validator.v2"
)

func ValidateStudent(student *models.Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}

func ValidateTeacher(teacher *models.Teacher) error {
	if err := validator.Validate(teacher); err != nil {
		return err
	}
	return nil
}
