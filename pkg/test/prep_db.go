package test

import (
	"context"
	"log"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func PrepDB(db *mongo.Database) {
	collection := db.Collection("blogs")
	if collection != nil {
		err := collection.Drop(context.TODO())
		if err != nil {
			log.Println("error in dropping collection blogs")
		}
	}

	ctx := context.TODO()

	err := db.CreateCollection(ctx, "blogs")
	if err != nil {
		log.Fatal("error in creating test blog collection")
	}

	collection = db.Collection("blogs")

	_, err = collection.InsertMany(ctx, []interface{}{
		models.Blog{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 21}.String(), Images: []string{"url1"}},
		models.Blog{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 20}.String(), Images: []string{"url1", "url2"}},
		models.Blog{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
		models.Blog{BlogID: "id4", AccountID: 2, Title: "title4", Summary: "summary4", Content: "content4", Tags: []string{"tag3", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 15}.String(), Images: []string{"url1"}},
		models.Blog{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
	})

	if err != nil {
		log.Println("error in inserting documents in blogs")
	}

	collection = db.Collection("tags")
	if collection != nil {
		err := collection.Drop(context.TODO())
		if err != nil {
			log.Println("error in dropping collection tags")
		}
	}

	err = db.CreateCollection(ctx, "tags")
	if err != nil {
		log.Fatal("error in creating test tags collection")
	}

	collection = db.Collection("tags")

	collection.InsertMany(ctx, []interface{}{
		models.Tag{Name: "tag1", BlogIDList: []string{"id2", "id5"}},
		models.Tag{Name: "tag2", BlogIDList: []string{"id1", "id2", "id3"}},
		models.Tag{Name: "tag3", BlogIDList: []string{"id4"}},
		models.Tag{Name: "tag4", BlogIDList: []string{}},
	})

	if err != nil {
		log.Println("error in inserting documents in tags")
	}
}
