package tag

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"go.mongodb.org/mongo-driver/bson"
)

type tag struct{}

func New() stores.Tag {
	return tag{}
}

// Get retrieves a tag by its name.
// A tag name uniquely identifies a tag entity.
// The tag entity has tag name and list of blog_id's associated with the tag.
func (t tag) Get(c *app.Context, name string) (*models.Tag, error) {
	collection := c.Mongo.DB().Collection("tags")

	var tag models.Tag

	res := collection.FindOne(c, bson.D{{Key: "_id", Value: name}})

	if err := res.Err(); err != nil {
		return nil, errors.DBError{Err: err}
	}

	err := res.Decode(&tag)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	return &tag, nil
}

// Update adds blog_id to given list of tags.
// Tags are created if they do not exist already.
func (t tag) Update(c *app.Context, blogID, tag string, operation constants.Operation) error {
	collection := c.Mongo.DB().Collection("tags")

	res := collection.FindOneAndUpdate(c, bson.D{{Key: "_id", Value: tag}}, bson.M{string(operation): bson.M{"blog_id_list": blogID}})
	if err := res.Err(); err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}

// Create inserts a new tag.
func (t tag) Create(c *app.Context, tag *models.Tag) error {
	collection := c.Mongo.DB().Collection("tags")

	_, err := collection.InsertOne(c, tag)
	if err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}

// Delete removes a tag by its name.
func (t tag) Delete(c *app.Context, tag string) error {
	collection := c.Mongo.DB().Collection("tags")

	res := collection.FindOneAndDelete(c, bson.D{bson.E{Key: "_id", Value: tag}})

	if err := res.Err(); err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}
