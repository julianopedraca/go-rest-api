package main

import (
	"julianopedraca/go-rest-api/db"
	"julianopedraca/go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
