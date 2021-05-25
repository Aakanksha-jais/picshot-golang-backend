package blog

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

func initializeTest() (*app.Context, stores.Blog) {
	app.InitializeTestBlogCollection(a.Mongo.DB(), a.Logger, "../../db")
	return &app.Context{Context: context.TODO(), App: a}, New()
}

func getTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z07:00", date)
	return t
}

func TestBlog_GetAll(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		input       *models.Blog
		output      []*models.Blog
		err         error
	}{
		{
			description: "get all with empty filter",
			input:       &models.Blog{},
			output: []*models.Blog{
				{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: getTime("2021-03-16T15:04:05Z"), Images: []string{"url1"}},
				{BlogID: "id4", AccountID: 2, Title: "title4", Summary: "summary4", Content: "content4", Tags: []string{"tag3", "tag2"}, CreatedOn: getTime("2021-03-15T15:04:05Z"), Images: []string{"url1"}},
				{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: getTime("2021-03-16T15:04:05Z"), Images: []string{"url1"}},
				{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: getTime("2021-03-20T15:04:05Z"), Images: []string{"url1", "url2"}},
				{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: getTime("2021-03-21T15:04:05Z"), Images: []string{"url1"}},
			},
			err: nil,
		},
		{
			description: "get all with account id = 5",
			input:       &models.Blog{AccountID: 5},
			output: []*models.Blog{
				{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: getTime("2021-03-16T15:04:05Z"), Images: []string{"url1"}},
				{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: getTime("2021-03-20T15:04:05Z"), Images: []string{"url1", "url2"}},
				{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: getTime("2021-03-21T15:04:05Z"), Images: []string{"url1"}},
			},
			err: nil,
		},
		{
			description: "get all with title = title5",
			input:       &models.Blog{Title: "title5"},
			output: []*models.Blog{
				{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: getTime("2021-03-16T15:04:05Z"), Images: []string{"url1"}},
			},
			err: nil,
		},
	}

	for i := range tests {
		output, err := blog.GetAll(ctx, tests[i].input)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_GetByIDs(t *testing.T) {
	ctx, blog := initializeTest()

	tests := []struct {
		description string
		input       []string
		output      []*models.Blog
		err         error
	}{
		{
			description: "get blogs with id = 5, 4, 1 (always returns in chronological order of creation)",
			input:       []string{"id5", "id4", "id1"},
			output: []*models.Blog{
				{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: getTime("2021-03-21T15:04:05Z"), Images: []string{"url1"}},
				{BlogID: "id4", AccountID: 2, Title: "title4", Summary: "summary4", Content: "content4", Tags: []string{"tag3", "tag2"}, CreatedOn: getTime("2021-03-15T15:04:05Z"), Images: []string{"url1"}},
				{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: getTime("2021-03-16T15:04:05Z"), Images: []string{"url1"}},
			},
			err: nil,
		},
		{
			description: "get blogs with non-existing id",
			input:       []string{"id0"},
			output:      []*models.Blog(nil),
			err:         nil,
		},
	}

	for i := range tests {
		output, err := blog.GetByIDs(ctx, tests[i].input)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
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
			description: "get blog with blog id = 5",
			input:       &models.Blog{BlogID: "id5"},
			output:      &models.Blog{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: getTime("2021-03-16T15:04:05Z"), Images: []string{"url1"}},
		},
		{
			description: "get blog with account id = 5",
			input:       &models.Blog{AccountID: 5},
			output:      &models.Blog{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: getTime("2021-03-21T15:04:05Z"), Images: []string{"url1"}},
		},
	}

	for i := range tests {
		output, err := blog.Get(ctx, tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_Create(t *testing.T) {
	ctx, blog := initializeTest()

	model := &models.Blog{BlogID: "TEST_ID", AccountID: 5, Title: "title", Summary: "summary", Content: "content", Tags: []string{"tag1"}, CreatedOn: getTime("2020-09-21T15:04:05Z"), Images: []string{"url8"}}

	tests := []struct {
		description string
		input       *models.Blog
		output      *models.Blog
		err         error
	}{
		{
			description: "Create Blog with Valid Details.",
			input:       model,
			output:      model,
		},
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
			description: "Valid Update on Title, Tags and Images.",
			input:       &models.Blog{BlogID: "id1", Title: "new_title", Tags: []string{"tag1", "tag3"}, Images: []string{"url8"}},
			output:      &models.Blog{BlogID: "id1", AccountID: 5, Title: "new_title", Summary: "summary1", Content: "content1", Tags: []string{"tag2", "tag1", "tag3"}, Images: []string{"url1", "url8"}, CreatedOn: getTime("2021-03-21T15:04:05Z")},
		},
		{
			description: "Valid Update on Content.",
			input:       &models.Blog{BlogID: "id2", Content: "new_content"},
			output:      &models.Blog{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "new_content", Tags: []string{"tag1", "tag2"}, Images: []string{"url1", "url2"}, CreatedOn: getTime("2021-03-20T15:04:05Z")},
		},
		{
			description: "Valid Update on Summary.",
			input:       &models.Blog{BlogID: "id2", Summary: "new_summary"},
			output:      &models.Blog{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "new_summary", Content: "new_content", Tags: []string{"tag1", "tag2"}, Images: []string{"url1", "url2"}, CreatedOn: getTime("2021-03-20T15:04:05Z")},
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
		{
			description: "Valid Delete.",
			blogID:      "id1",
			err:         nil,
		},
	}

	for i := range tests {
		err := blog.Delete(ctx, tests[i].blogID)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}
