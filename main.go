package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pontakornth/go-htmx/handlers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")
	router.GET("/articles", handlers.GetArticles)
	router.POST("/articles", handlers.PostArticle)
	router.Run(":8080")
}
