package main

import (
	"test/auth"
	"test/models"
	"test/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	// config := cors.DefaultConfig()/reader/search-book/reader@gmail.com

	// r.Use(cors.New(config))
	r.Use(cors.Default())

	r.Use(auth.Middleware())

	routes.LoadRoutes(r)

	r.Run(":3001")

}
