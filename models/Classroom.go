package models

type Classroom struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	TeacherID uint       `json:"teacher_id"`
	Teacher   *Teacher   `json:"teacher" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Students  []*Student `json:"students" gorm:"many2many:classroom_students"`
}
