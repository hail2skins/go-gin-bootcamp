package controllers

import (
	"fmt"
	"gin_notes/helpers"
	"gin_notes/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")

	// Check if email already exists
	available := models.CheckEmailAvailable(email)
	fmt.Println(available)
	if !available {
		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert": "Email already exists",
			},
		)
		return
	}
	if password != confirm_password {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Passwords do not match",
			},
		)
		return
	}
	user := models.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Unable to create user!",
			},
		)
	} else {
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := models.UserFindByEmailAndPassword(email, password)
	if user != nil {
		helpers.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert": "Invalid email and/or password",
			},
		)
	}
}

func Logout(c *gin.Context) {
	helpers.SessionClear(c)
	c.HTML(
		http.StatusOK,
		"home/index.html",
		gin.H{
			"alert": "Successfully logged out",
		},
	)
}
