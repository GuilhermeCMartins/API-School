package main

import (
	"api-school/database"
	"api-school/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectDB()

	routes.HandleRequests(r)
}
