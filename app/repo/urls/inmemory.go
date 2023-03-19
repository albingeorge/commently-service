package urls

import (
	urlmodel "github.com/albingeorge/commently-service/models/urls"
	"github.com/albingeorge/commently-service/utils"
)

var urlsTable map[string]urlmodel.Url

type DbImpl struct{}

func (db *DbImpl) Create(url string) urlmodel.Url {
	if v, ok := urlsTable[url]; ok {
		return v
	}
	id := utils.GenerateRandomId(8)
	currentUrl := urlmodel.Url{Id: id, Name: url}

	if urlsTable == nil {
		urlsTable = map[string]urlmodel.Url{url: currentUrl}
	} else {
		urlsTable[url] = currentUrl
	}
	return currentUrl
}

func (db *DbImpl) Get() []urlmodel.Url {
	data := []urlmodel.Url{}
	for _, v := range urlsTable {
		data = append(data, v)
	}

	return data
}
