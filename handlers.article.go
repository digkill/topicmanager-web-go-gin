package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticleCreationPage(context *gin.Context) {
	render(context, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func getArticle(context *gin.Context) {
	if articleID, err := strconv.Atoi(context.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(context, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")
		} else {
			context.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		context.AbortWithStatus(http.StatusNotFound)
	}
}
func createArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
