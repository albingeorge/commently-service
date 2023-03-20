package controllers

import (
	"log"
	"net/http"
	"net/url"

	commentrepo "github.com/albingeorge/commently-service/app/repo/comments"
	urlrepo "github.com/albingeorge/commently-service/app/repo/urls"
	"github.com/gin-gonic/gin"
)

var urlRepo urlrepo.Repository
var commentRepo commentrepo.Repository

type CommentInput struct {
	Url  string `json:"url" binding:"required"`
	Text string `json:"text" binding:"required"`
}

func init() {
	urlRepo = urlrepo.GetRepo()
	commentRepo = commentrepo.GetRepo()
}

func CommentCreate(c *gin.Context) {
	var commentJsonInput CommentInput

	if c.Bind(&commentJsonInput) == nil {
		urlEntity := urlRepo.Create(commentJsonInput.Url)
		commentRepo.Create(urlEntity.Id, commentJsonInput.Text)

		urls := urlRepo.Get()
		urlStr := make([]string, len(urls))
		i := 0
		for _, val := range urls {
			urlStr[i] = val.Name
			i++
		}
		log.Printf("Current Db:\n%v", urlStr)

		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
}

func CommentFetch(c *gin.Context) {
	// Parse input
	inputUrl := c.Query("url")
	log.Printf("Input URL: %q", inputUrl)
	decodedValue, err := url.QueryUnescape(inputUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid urlencoded value": inputUrl})
		log.Print(err)
		return
	}

	log.Printf("Decoded URL: %q", decodedValue)

	// Fetch url from db
	urlEntity, err := urlRepo.Fetch(decodedValue)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "failure"})
		log.Print(err)
		return
	}

	// Fetch comments from db
	comments := commentRepo.GetAll(urlEntity.Id)
	commentsStr := make([]string, len(comments))
	i := 0
	for _, val := range comments {
		commentsStr[i] = val.Text
		i++
	}

	c.JSON(http.StatusOK, gin.H{"comments": commentsStr})
}

func UrlsFetch(c *gin.Context) {
	urls := urlRepo.Get()
	urlStr := make([]string, len(urls))
	i := 0
	for _, val := range urls {
		urlStr[i] = val.Name
		i++
	}
	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
