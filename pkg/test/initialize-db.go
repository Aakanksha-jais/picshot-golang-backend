package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/datastore"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gchaincl/dotsql"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
)

func AddTestData(mongo datastore.MongoDB, sql datastore.SQLClient, awsS3 datastore.AWSS3, logger log.Logger) {
	InitializeTestBlogsCollection(mongo.Database, logger, "./db")

	InitializeTestTagsCollection(mongo.Database, logger, "./db")

	InitializeTestAccountsTable(sql.DB, logger, "./db")

	//InitializeTestAWSBucket(awsS3, os.Getenv("AWS_BUCKET"), logger, "./db")
}

func InitializeTestAWSBucket(awsS3 datastore.AWSS3, bucket string, logger log.Logger, directory string) {
	objects := make([]*s3.ObjectIdentifier, 0)

	svc := awsS3

	res, _ := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	for _, object := range res.Contents {
		objects = append(objects, &s3.ObjectIdentifier{Key: aws.String(*object.Key)})
	}

	if len(objects) != 0 {
		input := &s3.DeleteObjectsInput{
			Bucket: aws.String(bucket),
			Delete: &s3.Delete{Objects: objects},
		}

		_, err := svc.DeleteObjectsWithContext(context.TODO(), input)
		if err != nil {
			logger.Errorf("failed to delete objects: %v", err)
		}
	}

	uploader := s3manager.NewUploader(awsS3.Session)

	imgDir := directory + "/test/images/"

	files, err := ioutil.ReadDir(imgDir)
	if err != nil {
		logger.Errorf("failed to open folder images: %s", err.Error())
		return
	}

	for _, file := range files {
		f, err := os.Open(imgDir + file.Name())
		if err != nil {
			logger.Errorf("failed to open file %v: %v", file.Name(), err)
			return
		}

		_, _ = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(file.Name()),
			Body:   f,
			ACL:    aws.String("public-read"),
		})
	}
}

func InitializeTestAccountsTable(sqlDB *sql.DB, logger log.Logger, directory string) {
	file := directory + "/test/accounts.sql"

	dot, err := dotsql.LoadFromFile(file)
	if err != nil {
		logger.Errorf("cannot read file schema.sql: %v", err.Error())
		return
	}

	_, _ = dot.Exec(sqlDB, "drop")
	_, _ = dot.Exec(sqlDB, "create")
	_, _ = dot.Exec(sqlDB, "use")
	_, _ = dot.Exec(sqlDB, "create-table")
	_, _ = dot.Exec(sqlDB, "insert-aakanksha")
	_, _ = dot.Exec(sqlDB, "insert-mainak")
	_, _ = dot.Exec(sqlDB, "insert-divij")
}

func InitializeTestTagsCollection(db *mongo.Database, logger log.Logger, directory string) {
	file := directory + "/test/tags.json"

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Errorf("cannot read file tags.json: %v", err.Error())
		return
	}

	var (
		data []interface{}
		tags []models.Tag
	)

	err = json.Unmarshal(bytes, &tags)
	if err != nil {
		logger.Errorf("cannot unmarshal tags: %v", err.Error())
	}

	collection := db.Collection("tags")
	_ = collection.Drop(context.TODO())

	for _, tag := range tags {
		data = append(data, tag)
	}

	_, err = collection.InsertMany(context.TODO(), data)
	if err != nil {
		logger.Errorf("cannot insert data into tags collection: %v", err.Error())
	}
}

//nolint
func InitializeTestBlogsCollection(db *mongo.Database, logger log.Logger, directory string) {
	file := directory + "/test/blogs.json"

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Errorf("cannot read file blogs.json: %v", err.Error())
		return
	}

	var (
		data  []interface{}
		blogs []models.Blog
	)

	err = json.Unmarshal(bytes, &blogs)
	if err != nil {
		logger.Errorf("cannot unmarshal blogs: %v", err.Error())
	}

	collection := db.Collection("blogs")
	_ = collection.Drop(context.TODO())

	for _, blog := range blogs {
		data = append(data, blog)
	}

	_, err = collection.InsertMany(context.TODO(), data)
	if err != nil {
		logger.Errorf("cannot insert data into blogs collection: %v", err.Error())
	}
}
