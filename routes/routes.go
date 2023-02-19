package routes

import (
	"api-school/controllers"
	"api-school/middlewares"

	"github.com/gin-gonic/gin"
)

var IsAuthorized = middlewares.IsAuthorized()

func HandleRequests(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	r.POST("/students/register", controllers.CreateNewStudent)
	r.POST("/teacher/register", controllers.CreateNewTeacher)

	r.POST("/classrooms", controllers.CreateClassroom)
	r.GET("/classrooms/:id", controllers.GetStudentsByClassroomID)

	r.PATCH("/teacher/editgrades/:id", IsAuthorized, controllers.EditGrades)

	r.GET("/home", controllers.Home)
	r.GET("/premium", IsAuthorized, controllers.Premium)

	r.GET("/logout", controllers.Logout)

	r.Run()
}
