package comments

import commentmodel "github.com/albingeorge/commently-service/models/comments"

type Repository interface {
	Create(string) commentmodel.Comment
}
