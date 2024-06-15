package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []Article{
	{
		Id:      1,
		Title:   "Article 1",
		Content: "Content 1",
	},
}

func GetArticles(c *gin.Context) {
	f := c.NegotiateFormat(gin.MIMEHTML, gin.MIMEJSON)
	if f == gin.MIMEJSON {
		c.JSON(http.StatusOK, articles)
		return
	}
	c.HTML(http.StatusOK, "articles.html", articles)
}

func PostArticle(c *gin.Context) {
	// We get data from the form.
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	id := len(articles) + 1
	articles = append(articles, Article{Id: id, Title: title, Content: content})
	// Let's replace the entire thing.
	c.HTML(http.StatusCreated, "articles.html", articles)
}
