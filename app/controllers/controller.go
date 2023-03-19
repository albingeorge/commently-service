package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var Db map[string][]string

// type Comment string

type CommentInput struct {
	Url  string `json:"url" binding:"required"`
	Text string `json:"text" binding:"required"`
}

func CommentCreate(c *gin.Context) {
	// init Db
	if Db == nil {
		Db = map[string][]string{}
	}
	var commentJsonInput CommentInput

	if c.Bind(&commentJsonInput) == nil {
		if currentComments, ok := Db[commentJsonInput.Url]; ok {
			Db[commentJsonInput.Url] = append(currentComments, commentJsonInput.Text)
		} else {
			Db[commentJsonInput.Url] = []string{commentJsonInput.Text}
		}

		log.Printf("Current Db:\n%+v", Db)

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
}

func CommentFetch(c *gin.Context) {
	inputUrl := c.Query("url")
	log.Printf("Input URL: %q", inputUrl)
	decodedValue, err := url.QueryUnescape(inputUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid urlencoded value": inputUrl})
		log.Fatal(err)
	}

	log.Printf("Decoded URL: %q", decodedValue)

	comments := Db[decodedValue]
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func UrlsFetch(c *gin.Context) {
	urls := make([]string, len(Db))
	i := 0
	for k := range Db {
		urls[i] = k
		i++
	}
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
