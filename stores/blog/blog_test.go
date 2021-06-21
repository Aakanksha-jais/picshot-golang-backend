package blog

import (
	"context"
	"testing"
	"time"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/test"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

func initializeTest() (*app.Context, stores.Blog) {
	test.InitializeTestBlogsCollection(a.Mongo.Database, a.Logger, "../../db")
	return &app.Context{Context: context.TODO(), App: a}, New()
}

func getTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z07:00", date)
	return t
}

//nolint:lll // test cases need to be readable
func TestBlog_GetAll(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		input       *models.Blog
		page        *models.Page
		output      []*models.Blog
		err         error
	}{
		{
			description: "get all with empty filter and page = nil",
			input:       &models.Blog{},
			output:      getAllOutput(),
			err:         nil,
		},
		{
			description: "get all with empty filter and page limit = 2",
			input:       &models.Blog{},
			output: []*models.Blog{
				{BlogID: "MSI8WKNSH9", AccountID: 2, Title: "music", Summary: "a blog on music", Content: "avicii left :(", Tags: []string{}, CreatedOn: getTime("2021-05-23T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"}},
				{BlogID: "9SNVSH8K2M", AccountID: 3, Title: "flowers", Summary: "a blog on flowers", Content: "blue orchids are the most beautiful", Tags: []string{"#love", "#nature", "#life"}, CreatedOn: getTime("2021-04-16T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/orchids.jpg"}},
			},
			page: &models.Page{Limit: 2, PageNo: 1},
			err:  nil,
		},
		{
			description: "get all (page number = 1, limit = 3)",
			input:       &models.Blog{},
			output:      getAllOutput()[:3],
			page:        &models.Page{Limit: 3, PageNo: 1},
			err:         nil,
		},
		{description: "get all with empty filter and invalid page offset", input: &models.Blog{}, output: []*models.Blog(nil), page: &models.Page{Limit: 2, PageNo: 100}, err: nil},
		{description: "get all with empty filter and zero page limit", input: &models.Blog{}, output: getAllOutput(), page: &models.Page{Limit: 0, PageNo: 1}, err: nil},
		{description: "get all with empty filter and empty page", input: &models.Blog{}, output: getAllOutput(), page: &models.Page{}, err: nil},
		{
			description: "get all with account id = 2",
			input:       &models.Blog{AccountID: 2},
			output: []*models.Blog{
				{BlogID: "MSI8WKNSH9", AccountID: 2, Title: "music", Summary: "a blog on music", Content: "avicii left :(", Tags: []string{}, CreatedOn: getTime("2021-05-23T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"}},
				{BlogID: "ABK7SH2V37", AccountID: 2, Title: "chocolate", Summary: "a blog on chocolate", Content: "bournville is the best chocolate!", Tags: []string{"#cocoa", "#sweet", "#chocolate", "#trending"}, CreatedOn: getTime("2021-04-10T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/chocolate-gettyimages-473741340.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/chocolate.jpg"}},
				{BlogID: "9SH7SH2V37", AccountID: 2, Title: "memories", Summary: "a blog on memories", Content: "the best of childhood days!", Tags: []string{"#memories", "#life"}, CreatedOn: getTime("2021-04-08T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/girl.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/kid.jpg"}},
				{BlogID: "KN78FH8K2M", AccountID: 2, Title: "life", Summary: "a blog on life", Content: "life is a journey..", Tags: []string{"#life"}, CreatedOn: getTime("2021-01-06T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/691169.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/411820.jpg"}},
			},
			err: nil,
		},
		{
			description: "get all with account id = 2",
			input:       &models.Blog{AccountID: 2},
			output: []*models.Blog{
				{BlogID: "9SH7SH2V37", AccountID: 2, Title: "memories", Summary: "a blog on memories", Content: "the best of childhood days!", Tags: []string{"#memories", "#life"}, CreatedOn: getTime("2021-04-08T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/girl.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/kid.jpg"}},
				{BlogID: "KN78FH8K2M", AccountID: 2, Title: "life", Summary: "a blog on life", Content: "life is a journey..", Tags: []string{"#life"}, CreatedOn: getTime("2021-01-06T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/691169.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/411820.jpg"}},
			},
			page: &models.Page{PageNo: 2, Limit: 2},
			err:  nil,
		},
	}

	for i := range tests {
		output, err := blog.GetAll(ctx, tests[i].input, tests[i].page)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_GetAll_Error(t *testing.T) {
	ctx, blog := initializeTest()

	type demo struct {
		BlogID string `bson:"_id"`
		Title  []int  `bson:"title"`
	}

	collection := ctx.Mongo.Collection("blogs")

	_, _ = collection.InsertOne(ctx, demo{BlogID: "TEST_ID", Title: []int{1, 2, 3, 4}})

	tc := struct {
		description string
		input       *models.Blog
		page        *models.Page
		output      []*models.Blog
		err         error
	}{description: "get blog with blog id = TEST_ID", input: &models.Blog{BlogID: "TEST_ID"}, page: &models.Page{Limit: 2, PageNo: 1}, err: errors.DBError{}}

	output, err := blog.GetAll(ctx, tc.input, tc.page)

	assert.Equal(t, tc.output, output, "TEST, failed.\n%s", tc.description)

	assert.IsType(t, tc.err, err, "TEST, failed.\n%s", tc.description)
}

//nolint:lll // test cases need to be readable
func TestBlog_GetByIDs(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		input       []string
		output      []*models.Blog
		page        *models.Page
		err         error
	}{
		{
			description: "get blogs with id = MSI8WKNSH9, 9SH7SH2V37, MSO8WB2J7X",
			input:       []string{"MSI8WKNSH9", "9SH7SH2V37", "MSO8WB2J7X"},
			output: []*models.Blog{
				{BlogID: "MSI8WKNSH9", AccountID: 2, Title: "music", Summary: "a blog on music", Content: "avicii left :(", Tags: []string{}, CreatedOn: getTime("2021-05-23T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"}},
				{BlogID: "9SH7SH2V37", AccountID: 2, Title: "memories", Summary: "a blog on memories", Content: "the best of childhood days!", Tags: []string{"#memories", "#life"}, CreatedOn: getTime("2021-04-08T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/girl.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/kid.jpg"}},
				{BlogID: "MSO8WB2J7X", AccountID: 3, Title: "books", Summary: "a blog on books", Content: "the subtle art of not giving a fuck- an award winner", Tags: []string{"#markmanson"}, CreatedOn: getTime("2021-02-12T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/book.jpg"}},
			},
			page: &models.Page{Limit: 3, PageNo: 1},
			err:  nil,
		},
		{
			description: "get blogs (in reverse chronological order always)",
			input:       []string{"POQA7B2J7X", "MSI8WKNSH9", "9SH7SH2V37", "MSO8WB2J7X"},
			output: []*models.Blog{
				{BlogID: "MSI8WKNSH9", AccountID: 2, Title: "music", Summary: "a blog on music", Content: "avicii left :(", Tags: []string{}, CreatedOn: getTime("2021-05-23T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"}},
				{BlogID: "9SH7SH2V37", AccountID: 2, Title: "memories", Summary: "a blog on memories", Content: "the best of childhood days!", Tags: []string{"#memories", "#life"}, CreatedOn: getTime("2021-04-08T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/girl.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/kid.jpg"}},
				{BlogID: "MSO8WB2J7X", AccountID: 3, Title: "books", Summary: "a blog on books", Content: "the subtle art of not giving a fuck- an award winner", Tags: []string{"#markmanson"}, CreatedOn: getTime("2021-02-12T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/book.jpg"}},
			},
			page: &models.Page{Limit: 3, PageNo: 1},
			err:  nil,
		},
		{description: "get blogs with non-existing id", input: []string{"DUMMY_ID_123"}, output: []*models.Blog(nil), err: nil},
	}

	for i := range tests {
		output, err := blog.GetByIDs(ctx, tests[i].input, tests[i].page)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_GetByIDs_Error(t *testing.T) {
	ctx, blog := initializeTest()

	type demo struct {
		BlogID string `bson:"_id"`
		Title  []int  `bson:"title"`
	}

	collection := ctx.Mongo.Collection("blogs")

	_, _ = collection.InsertOne(ctx, demo{BlogID: "TEST_ID", Title: []int{1, 2, 3, 4}})

	tc := struct {
		description string
		input       []string
		output      []*models.Blog
		err         error
	}{description: "get blog with blog id = TEST_ID", input: []string{"TEST_ID"}, err: errors.DBError{}}

	output, err := blog.GetByIDs(ctx, tc.input, &models.Page{Limit: 3, PageNo: 1})

	assert.Equal(t, tc.output, output, "TEST, failed.\n%s", tc.description)

	assert.IsType(t, tc.err, err, "TEST, failed.\n%s", tc.description)
}

func TestBlog_Get(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		input       *models.Blog
		output      *models.Blog
		err         error
	}{
		{
			description: "get blog with blog id = MSO8WB2J7X",
			input:       &models.Blog{BlogID: "MSO8WB2J7X"},
			output:      &models.Blog{BlogID: "MSO8WB2J7X", AccountID: 3, Title: "books", Summary: "a blog on books", Content: "the subtle art of not giving a fuck- an award winner", Tags: []string{"#markmanson"}, CreatedOn: getTime("2021-02-12T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/book.jpg"}},
		},
		{description: "get blog with non-existing blog id", input: &models.Blog{BlogID: "DUMMY_ID_123"}, err: errors.DBError{Err: mongo.ErrNoDocuments}},
	}

	for i := range tests {
		output, err := blog.Get(ctx, tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_Get_Error(t *testing.T) {
	ctx, blog := initializeTest()

	type demo struct {
		BlogID string `bson:"_id"`
		Title  []int  `bson:"title"`
	}

	collection := ctx.Mongo.Collection("blogs")

	_, _ = collection.InsertOne(ctx, demo{BlogID: "TEST_ID", Title: []int{1, 2, 3, 4}})

	tc := struct {
		description string
		input       *models.Blog
		output      *models.Blog
		err         error
	}{description: "get blog with blog id = TEST_ID", input: &models.Blog{BlogID: "TEST_ID"}, err: errors.DBError{}}

	output, err := blog.Get(ctx, tc.input)

	assert.Equal(t, tc.output, output, "TEST, failed.\n%s", tc.description)

	assert.IsType(t, tc.err, err, "TEST, failed.\n%s", tc.description)
}

func TestBlog_Create(t *testing.T) {
	ctx, blog := initializeTest()

	model := &models.Blog{
		BlogID:    "TEST_ID",
		AccountID: 5,
		Title:     "title",
		Summary:   "summary",
		Content:   "content",
		Tags:      []string{"tag"},
		CreatedOn: getTime("2020-09-21T15:04:05Z"),
		Images:    []string{"url"},
	}

	tests := []struct {
		description string
		input       *models.Blog
		output      *models.Blog
		err         error
	}{
		{description: "Create Blog with Valid Details.", input: model, output: model},
	}

	for i := range tests {
		output, err := blog.Create(ctx, tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_Update(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		input       *models.Blog
		output      *models.Blog
		err         error
	}{
		{
			description: "valid update on title, tags and images.",
			input:       &models.Blog{BlogID: "MSO8WB2J7X", Title: "new_title", Tags: []string{"tag1", "tag3"}, Images: []string{"url8"}},
			output:      &models.Blog{BlogID: "MSO8WB2J7X", AccountID: 3, Title: "new_title", Summary: "a blog on books", Content: "the subtle art of not giving a fuck- an award winner", Tags: []string{"tag1", "tag3"}, CreatedOn: getTime("2021-02-12T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/book.jpg", "url8"}},
		},
		{
			description: "valid update on content and tags.",
			input:       &models.Blog{BlogID: "POQA7B2J7X", Content: "new_content", Tags: []string{"#love", "life"}},
			output:      &models.Blog{BlogID: "POQA7B2J7X", AccountID: 1, Title: "love", Summary: "a blog on love", Content: "new_content", Tags: []string{"#love", "life"}, CreatedOn: getTime("2019-03-16T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/3b6c399e8f5d9d4f54f4d91c6db7cfde.jpg"}},
		},
		{
			description: "valid update on summary.",
			input:       &models.Blog{BlogID: "9SH7SH2V37", Summary: "new_summary", Tags: []string{"#memories", "#life"}},
			output:      &models.Blog{BlogID: "9SH7SH2V37", AccountID: 2, Title: "memories", Summary: "new_summary", Content: "the best of childhood days!", Tags: []string{"#memories", "#life"}, CreatedOn: getTime("2021-04-08T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/girl.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/kid.jpg"}},
		},
	}

	for i := range tests {
		output, err := blog.Update(ctx, tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_DeleteBlog(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		blogID      string
		err         error
	}{
		{description: "valid delete.", blogID: "9SH7SH2V37", err: nil},
		{description: "invalid delete.", blogID: "9SH7SH2V37", err: errors.DBError{Err: mongo.ErrNoDocuments}},
	}

	for i := range tests {
		err := blog.Delete(ctx, tests[i].blogID)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

//nolint:lll // hampers readability
func getAllOutput() []*models.Blog {
	return []*models.Blog{
		{BlogID: "MSI8WKNSH9", AccountID: 2, Title: "music", Summary: "a blog on music", Content: "avicii left :(", Tags: []string{}, CreatedOn: getTime("2021-05-23T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"}},
		{BlogID: "9SNVSH8K2M", AccountID: 3, Title: "flowers", Summary: "a blog on flowers", Content: "blue orchids are the most beautiful", Tags: []string{"#love", "#nature", "#life"}, CreatedOn: getTime("2021-04-16T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/orchids.jpg"}},
		{BlogID: "ABK7SH2V37", AccountID: 2, Title: "chocolate", Summary: "a blog on chocolate", Content: "bournville is the best chocolate!", Tags: []string{"#cocoa", "#sweet", "#chocolate", "#trending"}, CreatedOn: getTime("2021-04-10T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/chocolate-gettyimages-473741340.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/chocolate.jpg"}},
		{BlogID: "9SH7SH2V37", AccountID: 2, Title: "memories", Summary: "a blog on memories", Content: "the best of childhood days!", Tags: []string{"#memories", "#life"}, CreatedOn: getTime("2021-04-08T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/girl.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/kid.jpg"}},
		{BlogID: "UMS672XR8J", AccountID: 1, Title: "coffee", Summary: "a blog on coffee", Content: "visit starbucks for best coffee!", Tags: []string{"#caffiene", "#coffee", "#trending"}, CreatedOn: getTime("2021-03-25T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/812231.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/809031.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/761638.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/653598.jpg"}},
		{BlogID: "U72B72XR8J", AccountID: 1, Title: "songs", Summary: "a blog on songs", Content: "billie eilish is love<3", Tags: []string{"#love", "#eilish"}, CreatedOn: getTime("2021-03-14T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/billieeilish.jpeg", "https://picshot-images.s3.ap-south-1.amazonaws.com/billie.jpg"}},
		{BlogID: "MSO8WB2J7X", AccountID: 3, Title: "books", Summary: "a blog on books", Content: "the subtle art of not giving a fuck- an award winner", Tags: []string{"#markmanson"}, CreatedOn: getTime("2021-02-12T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/book.jpg"}},
		{BlogID: "KN78FH8K2M", AccountID: 2, Title: "life", Summary: "a blog on life", Content: "life is a journey..", Tags: []string{"#life"}, CreatedOn: getTime("2021-01-06T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/691169.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/411820.jpg"}},
		{BlogID: "MSI8NS2909", AccountID: 3, Title: "movies", Summary: "a blog on movies", Content: "oculus is terrific!!!", Tags: []string{"#movie", "#movies", "#trending"}, CreatedOn: getTime("2020-12-15T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/oculus.jpg"}},
		{BlogID: "POQA7B2J7X", AccountID: 1, Title: "love", Summary: "a blog on love", Content: "<3", Tags: []string{"#love", "life", "#trending"}, CreatedOn: getTime("2019-03-16T15:04:05Z"), Images: []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/3b6c399e8f5d9d4f54f4d91c6db7cfde.jpg"}},
	}
}
