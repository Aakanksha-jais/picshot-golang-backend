package models

import "go.mongodb.org/mongo-driver/bson"

// Blog is filterable by blog id, account id and title ONLY.
type Blog struct {
	BlogID    string   `bson:"_id" json:"-"`                 // Unique Blog ID
	AccountID int64    `bson:"account_id" json:"account_id"` // ID of Account associated with the Blog
	Title     string   `bson:"title" json:"title"`           // Title of Blog
	Summary   string   `bson:"summary" json:"summary"`       // Summary by-line
	Content   string   `bson:"content" json:"content"`       // Detailed Content of Blog
	Tags      []string `bson:"tags" json:"tags"`             // List of Tags associated with the Blog
	CreatedOn string   `bson:"created_on" json:"created_on"` // Date of Creation of Blog
	Images    []string `bson:"images" json:"images"`         // URL of images stored in cloud
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
