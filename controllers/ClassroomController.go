package controllers

import (
	"api-school/database"
	"api-school/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClassroom(c *gin.Context) {

	var req struct {
		Name       string `json:"name" binding:"required"`
		TeacherID  uint   `json:"teacher_id" binding:"required"`
		StudentIDs []uint `json:"student_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	classroom := models.Classroom{
		Name:      req.Name,
		TeacherID: req.TeacherID,
		Students:  make([]*models.Student, len(req.StudentIDs)),
	}

	for i, studentID := range req.StudentIDs {
		student, err := GetStudentByID(studentID)
		if err != nil {
			return
		}
		classroom.Students[i] = student
	}

	teacher, err := GetTeacherById(req.TeacherID)
	if err != nil {
		return
	}
	classroom.Teacher = teacher

	if err := database.DB.Create(&classroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create classroom"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"classroom": classroom})
}

func GetStudentsByClassroomID(c *gin.Context) {
	var students []models.Student
	id := c.Param("id")
	result := database.DB.Table("students").
		Joins("JOIN classroom_students ON students.id = classroom_students.student_id").
		Where("classroom_students.classroom_id = ?", id).
		Find(&students)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Classroom not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"classroom": students,
	})
}
