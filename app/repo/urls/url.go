package urls

import (
	urlmodel "github.com/albingeorge/commently-service/models/urls"
)

type Repository interface {
	Create(string) urlmodel.Url
	Get() []urlmodel.Url
}

func GetRepo() Repository {
	return &DbImpl{}
}
