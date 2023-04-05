package main

import (
	"gin_notes/controllers"
	"gin_notes/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.LoadEnv()
	helpers.LoadDatabase()

	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes Application",
		})
	})

	r.GET("/notes", controllers.NotesIndex)
	r.GET("notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)

	log.Println("Server Started")
	r.Run() // default port 8080
}
