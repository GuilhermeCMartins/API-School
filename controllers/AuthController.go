package controllers

import (
	"api-school/database"
	"api-school/models"
	"api-school/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingStudent models.Student

	database.DB.Where("email = ?", user.Email).First(&existingStudent)

	if existingStudent.ID != 0 {
		errHash := utils.CompareHashPassword(user.Password, existingStudent.Password)

		if !errHash {
			c.JSON(400, gin.H{"error": "invalid password"})
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
			Role: existingStudent.Role,
			StandardClaims: jwt.StandardClaims{
				Subject:   existingStudent.Email,
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			c.JSON(500, gin.H{"error": "could not generate token"})
			return
		}

		c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
		c.JSON(200, gin.H{"success": "Student logged in"})
		return
	}

	var existingTeacher models.Teacher

	database.DB.Where("email = ?", user.Email).First(&existingTeacher)

	if existingTeacher.ID != 0 {

		errHash := utils.CompareHashPassword(user.Password, existingTeacher.Password)

		if !errHash {
			c.JSON(400, gin.H{"error": "invalid password"})
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
			Role: existingTeacher.Role,
			StandardClaims: jwt.StandardClaims{
				Subject:   existingTeacher.Email,
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			c.JSON(500, gin.H{"error": "could not generate token"})
			return
		}

		c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
		c.JSON(200, gin.H{"success": "Teacher logged in"})
		return
	}

	c.JSON(400, gin.H{"error": "User does not exists"})
}

func Home(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "student" && claims.Role != "teacher" {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}

func Premium(c *gin.Context) {
	c.JSON(200, gin.H{"success": "premium page"})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}