package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.Run()
}
