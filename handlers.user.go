package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func showLoginPage(context *gin.Context) {
	render(context, gin.H{"title": "Login"}, "login.html")
}

func performLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	// var sameSiteCookie http.SameSite

	if isUserValid(username, password) {
		token := generateSessionToken()
		context.SetCookie("token", token, 3600, "", "", false, true)
		context.Set("is_logged_in", true)

		render(context, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		context.HTML(http.StatusBadGateway, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(context *gin.Context) {
	// var sameSiteCookie http.SameSite;

	context.SetCookie("token", "", -1, "", "", false, true)

	context.Redirect(http.StatusTemporaryRedirect, "/")
}

func showRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// var sameSiteCookie http.SameSite;

	if _, err := registerNewUser(username, password); err == nil {

		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}
