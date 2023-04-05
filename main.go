package main

import (
	"fmt"
	"gin_notes/controllers"
	"gin_notes/helpers"
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
			"logged_in": helpers.IsUserLoggedIn(c),
		})
		fmt.Println(c.GetUint64("user_id"))
	})

	// Route Group - /notes
	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("/edit/:id", controllers.NotesEditPage)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/logout", controllers.Logout)

	log.Println("Server Started")
	r.Run() // default port 8080
}
