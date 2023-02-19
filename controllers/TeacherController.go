package controllers

import (
	"api-school/database"
	"api-school/models"
	"api-school/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingTeacher models.Teacher

	database.DB.Where("email=?", teacher.Email).First(&existingTeacher)

	if existingTeacher.ID != 0 {
		c.JSON(400, gin.H{
			"error": "Teacher already exists",
		})
		return
	}

	if err := utils.ValidateTeacher(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var errHash error
	teacher.Password, errHash = utils.GenerateHashPassword(teacher.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	teacher.Role = "teacher"

	database.DB.Create(&teacher)
	c.JSON(http.StatusOK, gin.H{
		"Sucess": "Teacher created",
	})
}

func EditGrades(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	var testStudent models.Student

	c.BindJSON(&testStudent)

	if validErrs := (*models.Student).Validate(&testStudent); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if !testStudent.GradeModified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Set GradeFlag to true to edit grades."})
		return
	}

	student.Grade = testStudent.Grade
	student.Frequency = testStudent.Frequency

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)

}

func GetAllClasses(c *gin.Context) {
	var teacher models.Teacher
	err := database.DB.Preload("Classroom").First(&teacher, teacher.ID).Error
	if err != nil {
		fmt.Println(err)
	}

	// The classrooms associated with the teacher are now available in the teacher.Classroom slice
	fmt.Printf("Teacher %s has %d classrooms:\n", teacher.Name, len(teacher.Classroom))
	for _, c := range teacher.Classroom {
		fmt.Printf("- %s\n", c.Name)
	}
}

func GetTeacherById(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	err := database.DB.Where("id = ?", id).First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}
