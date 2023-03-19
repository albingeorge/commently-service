package comments

import commentmodel "github.com/albingeorge/commently-service/models/comments"

type Repository interface {
	Create(urlId string, text string) commentmodel.Comment
	GetAll(urlId string) []*commentmodel.Comment
}

func GetRepo() Repository {
	return &DbImpl{}
}
