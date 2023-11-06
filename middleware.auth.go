package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		loggedInInterface, _ := context.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		loggedInInterface, _ := context.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
