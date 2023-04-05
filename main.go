package main

import (
	"fmt"
	"gin_notes/controllers"
	"gin_notes/middlewares"
	"gin_notes/setup"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	setup.LoadEnv()
	setup.LoadDatabase()

	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	// Sessions Init
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	r.LoadHTMLGlob("templates/**/**")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "Notes Application",
			"logged_in": (c.GetUint64("user_id") > 0),
		})
		fmt.Println(c.GetUint64("user_id"))
	})

	r.GET("/notes", controllers.NotesIndex)
	r.GET("notes/new", controllers.NotesNew)
	r.POST("/notes", controllers.NotesCreate)
	r.GET("/notes/:id", controllers.NotesShow)
	r.GET("/notes/edit/:id", controllers.NotesEditPage)
	r.POST("/notes/:id", controllers.NotesUpdate)
	r.DELETE("/notes/:id", controllers.NotesDelete)

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/logout", controllers.Logout)

	log.Println("Server Started")
	r.Run() // default port 8080
}
