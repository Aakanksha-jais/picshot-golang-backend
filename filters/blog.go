package filters

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Blog struct {
	BlogID    string
	AccountID int64
	Title     string
}

func (b Blog) GetFilter() bson.D {
	var filter = bson.D{}

	if b.BlogID != "" {
		filter = append(filter, bson.E{Key: "_id", Value: b.BlogID})
	}

	if b.AccountID != 0 {
		filter = append(filter, bson.E{Key: "account_id", Value: b.AccountID})
	}

	if b.Title != "" {
		filter = append(filter, bson.E{Key: "title", Value: b.Title})
	}

	return filter
}
