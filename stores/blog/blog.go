package blog

import (
	"context"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type blog struct {
	db *mongo.Database
}

func New(db *mongo.Database) blog {
	return blog{db: db}
}

// GetAll is used to retrieve all blogs that match the filter.
// BLogs can be filtered by account_id, blog_id and title.
func (b blog) GetAll(ctx context.Context, filter models.Blog) ([]*models.Blog, error) {
	collection := b.db.Collection("blogs")

	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}}) // retrieve the blogs in reverse chronological order

	cursor, err := collection.Find(ctx, filter.GetFilter(), opts)
	if err != nil {
		return nil, err
	}

	blogs := make([]*models.Blog, 0)

	for cursor.Next(ctx) {
		var blog models.Blog

		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blog)
	}

	return blogs, cursor.Close(ctx)
}

// GetByIDs retrieves all blogs whose IDs have been provided as parameter.
func (b blog) GetByIDs(ctx context.Context, idList []string) ([]*models.Blog, error) {
	collection := b.db.Collection("blogs")

	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": idList}})
	if err != nil {
		return nil, err
	}

	blogs := make([]*models.Blog, 0)

	for cursor.Next(ctx) {
		var blog models.Blog

		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blog)
	}

	return blogs, cursor.Close(ctx)
}

// Get is used to retrieve a SINGLE blog that matches the filter.
// A blog can be filtered by account_id, blog_id and title.
func (b blog) Get(ctx context.Context, filter models.Blog) (*models.Blog, error) {
	var blog models.Blog

	collection := b.db.Collection("blogs")

	res := collection.FindOne(ctx, filter.GetFilter())

	err := res.Err()
	switch err {
	case mongo.ErrNoDocuments:
		return nil, err
	case nil:
		err = res.Decode(&blog)
		if err != nil {
			return nil, err
		}

		return &blog, err
	default:
		return nil, err
	}
}

// Create is used to create a new blog.
func (b blog) Create(ctx context.Context, model models.Blog) (*models.Blog, error) {
	collection := b.db.Collection("blogs")

	res, err := collection.InsertOne(ctx, model) // nil is returned if InsertOne operation is successful
	if err != nil {
		return nil, err
	}

	id := res.InsertedID

	return b.Get(ctx, models.Blog{BlogID: id.(string)})
}

// Update updates the blog by its ID.
func (b blog) Update(ctx context.Context, model models.Blog) (*models.Blog, error) {
	collection := b.db.Collection("blogs")

	res := collection.FindOneAndUpdate(ctx, bson.M{"_id": model.BlogID}, generateFilter(model))

	if err := res.Err(); err != nil {
		return nil, err
	}

	if !reflect.DeepEqual(model.Images, []string{}) {
		r := collection.FindOneAndUpdate(ctx, bson.M{"_id": model.BlogID}, bson.M{"$push": bson.M{"images": bson.M{"$each": model.Images}}})
		if err := r.Err(); err != nil {
			return nil, err
		}
	}

	if !reflect.DeepEqual(model.Tags, []string{}) {
		r := collection.FindOneAndUpdate(ctx, bson.M{"_id": model.BlogID}, bson.M{"$push": bson.M{"tags": bson.M{"$each": model.Tags}}})
		if err := r.Err(); err != nil {
			return nil, err
		}
	}

	return b.Get(ctx, models.Blog{BlogID: model.BlogID})
}

func generateFilter(model models.Blog) bson.M {
	update := bson.M{}
	if model.Title != "" {
		update["title"] = model.Title
	}

	if model.Summary != "" {
		update["summary"] = model.Summary
	}

	if model.Content != "" {
		update["content"] = model.Content
	}

	return bson.M{"$set": update}
}

// Delete deletes a blog by its ID.
func (b blog) Delete(ctx context.Context, blogID string) error {
	collection := b.db.Collection("blogs")

	res := collection.FindOneAndDelete(ctx, bson.D{bson.E{Key: "_id", Value: blogID}})

	return res.Err()
}
