package comments

import (
	commentsmodel "github.com/albingeorge/commently-service/models/comments"
	"github.com/albingeorge/commently-service/utils"
)

// Indexed by comment id
var commentsTableById map[string]*commentsmodel.Comment

// Indexed by url
var commentsTableByUrl map[string][](*commentsmodel.Comment)

type DbImpl struct{}

func init() {
	commentsTableById = map[string]*commentsmodel.Comment{}
	commentsTableByUrl = map[string][](*commentsmodel.Comment){}
}
func (db *DbImpl) Create(urlId string, text string) commentsmodel.Comment {
	comment := commentsmodel.Comment{
		Id:    utils.GenerateRandomId(8),
		UrlId: urlId,
		Text:  text,
	}

	// It's always a new comment, so we don't need to check if a comment id exists
	commentsTableById[comment.Id] = &comment

	if val, ok := commentsTableByUrl[urlId]; ok {
		val = append(val, &comment)
		commentsTableByUrl[urlId] = val
	} else {
		commentsTableByUrl[urlId] = [](*commentsmodel.Comment){&comment}
	}

	return comment
}

func (db *DbImpl) GetAll(urlId string) []*commentsmodel.Comment {
	comments := []*commentsmodel.Comment{}
	if val, ok := commentsTableByUrl[urlId]; ok {
		return val
	}
	return comments
}
