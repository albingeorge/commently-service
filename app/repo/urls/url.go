package urls

import (
	commentsmodel "github.com/albingeorge/commently-service/app/models/comments"
	urlmodel "github.com/albingeorge/commently-service/app/models/urls"
)

type Repository interface {
	Create(string) *urlmodel.Url
	Fetch(string) (*urlmodel.Url, error)
	Get() []*urlmodel.Url
	GetComments(string) ([](*commentsmodel.Comment), error)
}

func GetRepo() Repository {
	return &DbImpl{}
}
