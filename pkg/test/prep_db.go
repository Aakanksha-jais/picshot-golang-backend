package test

import (
	"context"
	"encoding/json"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"io/ioutil"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeTestDB(db *mongo.Database, logger log.Logger) {
	ctx := context.TODO()

	bytes, err := ioutil.ReadFile("../../db/test_blogs.json")
	if err != nil {
		logger.Errorf("cannot read file blogs.json: %v", err.Error())
	}

	var (
		data []interface{}
		blogs []models.Blog
	)

	err = json.Unmarshal(bytes, &blogs)
	if err != nil {
		logger.Errorf("cannot unmarshal blogs: %v", err.Error())
	}

	collection := db.Collection("blogs")
	collection.Drop(ctx)

	for _, blog:= range blogs{
		data = append(data, blog)
	}

	_, err = collection.InsertMany(ctx, data)
	if err != nil {
		logger.Errorf("cannot insert data into blogs collection: %v", err.Error())
	}

	bytes, err = ioutil.ReadFile("../../db/test_tags.json")
	if err != nil {
		logger.Errorf("cannot read file tags.json: %v", err.Error())
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		logger.Errorf("cannot unmarshal tags: %v", err.Error())
	}

	collection = db.Collection("tags")
	collection.Drop(ctx)

	_, err = collection.InsertMany(ctx, data)
	if err != nil {
		logger.Errorf("cannot insert data into tags collection: %v", err.Error())
	}
}
