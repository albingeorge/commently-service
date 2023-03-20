package urls

import (
	"fmt"

	commentsmodel "github.com/albingeorge/commently-service/app/models/comments"
	urlmodel "github.com/albingeorge/commently-service/app/models/urls"
	commentsrepo "github.com/albingeorge/commently-service/app/repo/comments"
	"github.com/albingeorge/commently-service/app/utils"
)

var urlsTableIndexedByUrl map[string]*urlmodel.Url

type DbImpl struct{}

func init() {
	urlsTableIndexedByUrl = map[string]*urlmodel.Url{}
}
func (db *DbImpl) Create(url string) *urlmodel.Url {
	if v, ok := urlsTableIndexedByUrl[url]; ok {
		return v
	}
	id := utils.GenerateRandomId(8)
	currentUrl := urlmodel.Url{Id: id, Name: url}

	urlsTableIndexedByUrl[url] = &currentUrl

	return &currentUrl
}

func (db *DbImpl) Fetch(url string) (*urlmodel.Url, error) {
	if v, ok := urlsTableIndexedByUrl[url]; ok {
		return v, nil
	}
	return &urlmodel.Url{}, fmt.Errorf("error fetching url: %q", url)
}

func (db *DbImpl) Get() []*urlmodel.Url {
	data := []*urlmodel.Url{}
	for _, v := range urlsTableIndexedByUrl {
		data = append(data, v)
	}

	return data
}

func (db *DbImpl) GetComments(url string) ([]*commentsmodel.Comment, error) {
	urlEntity, err := db.Fetch(url)
	if err != nil {
		return []*commentsmodel.Comment{}, err
	}
	urlId := urlEntity.Id

	return commentsrepo.GetRepo().GetAll(urlId), nil
}
