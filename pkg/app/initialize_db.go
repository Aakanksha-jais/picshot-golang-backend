package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"io/ioutil"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/gchaincl/dotsql"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeDB(db *mongo.Database, sqlDB *sql.DB, logger log.Logger) {
	ctx := context.TODO()

	bytes, err := ioutil.ReadFile("./db/blogs.json")
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

	bytes, err = ioutil.ReadFile("./db/tags.json")
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

	dot, err := dotsql.LoadFromFile("./db/schema.sql")
	if err != nil {
		logger.Errorf("cannot read file schema.sql: %v", err.Error())
	}

	dot.Exec(sqlDB, "drop")
	dot.Exec(sqlDB, "create")
	dot.Exec(sqlDB, "use")
	dot.Exec(sqlDB, "create-table")

	dot, err = dotsql.LoadFromFile("./db/test_data.sql")
	if err != nil {
		logger.Errorf("cannot read file schema.sql: %v", err.Error())
	}

	dot.Exec(sqlDB, "use")
	dot.Exec(sqlDB, "insert-aakanksha")
	dot.Exec(sqlDB, "insert-mainak")
	dot.Exec(sqlDB, "insert-divij")
}
