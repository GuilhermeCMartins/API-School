package models

type Teacher struct {
	User
	Role      string
	Classroom []*Classroom `json:"classroom"`
}
