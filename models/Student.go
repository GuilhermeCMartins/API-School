package models

import (
	"net/url"
)

type Student struct {
	User
	Role          string `default:"student" `
	GradeID       int
	Grade         Grade        `json:"notas" `
	GradeModified bool         `json:"GradeFlag"`
	Frequency     float64      `json:"frequencia"`
	Classrooms    []*Classroom `json:"classrooms" gorm:"many2many:classroom_students"`
}

func (s *Student) Validate() url.Values {
	errs := url.Values{}
	if s.Email != "" {
		errs.Add("email", "você não tem permissão")
	}

	if s.Password != "" {
		errs.Add("password", "você não tem permissão")
	}

	if s.Name != "" {
		errs.Add("name", "você não tem permissão")
	}

	if s.Role != "" {
		errs.Add("role", "você não tem permissão")
	}

	return errs
}
