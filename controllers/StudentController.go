package controllers

import (
	"api-school/database"
	"api-school/models"
	"api-school/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key") // Refactor to .env

func CreateNewStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingStudent models.Student

	database.DB.Where("email=?", student.Email).First(&existingStudent)

	if existingStudent.ID != 0 {
		c.JSON(400, gin.H{
			"error": "Student already exists",
		})
		return
	}

	if err := utils.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var errHash error
	student.Password, errHash = utils.GenerateHashPassword(student.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	student.Role = "student"

	database.DB.Create(&student)
	c.JSON(http.StatusOK, gin.H{
		"Sucess": "Student created",
	})
}

func GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	err := database.DB.Where("id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
