package tag

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type tag struct {
	db *mongo.Database
}

func New(db *mongo.Database) tag {
	return tag{db}
}

// GetByName retrieves a tag by its name.
// A tag name uniquely identifies a tag entity.
// The tag entity has tag name and list of blog_id's associated with the tag.
func (t tag) GetByName(ctx context.Context, name string) (*models.Tag, error) {
	collection := t.db.Collection("tags")

	var tag models.Tag

	res := collection.FindOne(ctx, bson.D{{Key: "_id", Value: name}})

	err := res.Err()
	switch err {
	case mongo.ErrNoDocuments:
		return nil, errors.EntityNotFound{Entity: "tag", ID: name}
	}
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	err = res.Decode(&tag)
	if err != nil {
		return nil, errors.DBError{Err: err}
	}

	return &tag, err
}

// AddBlogID adds blog_id to given list of tags.
// Tags are created if they do not exist already.
func (t tag) AddBlogID(ctx context.Context, blogID string, tags []string) ([]*models.Tag, error) {
	collection := t.db.Collection("tags")

	var res []*models.Tag

	for i := range tags {
		_, err := t.GetByName(ctx, tags[i])

		switch err.(type) {
		case errors.EntityNotFound:
			tag := models.Tag{Name: tags[i], BlogIDList: []string{blogID}}

			_, err = collection.InsertOne(ctx, tag)
			if err != nil {
				return nil, err
			}
		case nil:
			collection.FindOneAndUpdate(ctx, bson.D{{Key: "_id", Value: tags[i]}}, bson.M{"$push": bson.M{"blog_id_list": blogID}})
		default:
			return nil, err
		}

		tag, _ := t.GetByName(ctx, tags[i])
		res = append(res, tag)
	}

	return res, nil
}

// RemoveBlogID removes blog_id from given list of tags.
func (t tag) RemoveBlogID(ctx context.Context, blogID string, tags []string) ([]*models.Tag, error) {
	collection := t.db.Collection("tags")

	var res []*models.Tag

	for i := range tags {
		_, err := t.GetByName(ctx, tags[i])
		if err != nil {
			return nil, err
		}

		collection.FindOneAndUpdate(ctx, bson.D{{Key: "_id", Value: tags[i]}}, bson.M{"$pull": bson.M{"blog_id_list": blogID}})

		tag, _ := t.GetByName(ctx, tags[i])
		res = append(res, tag)
	}

	return res, nil
}
