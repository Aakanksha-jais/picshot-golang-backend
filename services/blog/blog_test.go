package blog

import (
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/types"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

func initializeTest(t *testing.T) (*stores.MockBlog, *stores.MockTag, *stores.MockImage, *app.Context, blog) {
	ctrl := gomock.NewController(t)
	mockBlogStore := stores.NewMockBlog(ctrl)
	mockTagStore := stores.NewMockTag(ctrl)
	mockImageStore := stores.NewMockImage(ctrl)
	ctx := app.NewContext(nil, &app.App{Logger: log.NewLogger()})
	mockService := New(mockBlogStore, mockTagStore, mockImageStore)

	return mockBlogStore, mockTagStore, mockImageStore, ctx, mockService
}

func TestBlog_GetAll(t *testing.T) {
	mockBlogStore, _, _, ctx, mockService := initializeTest(t)

	tests := []struct {
		description string
		input       *models.Blog
		output      []*models.Blog
	}{
		{
			description: "get all with empty filter",
			input:       &models.Blog{},
			output: []*models.Blog{
				{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
				{BlogID: "id4", AccountID: 2, Title: "title4", Summary: "summary4", Content: "content4", Tags: []string{"tag3", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 15}.String(), Images: []string{"url1"}},
				{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
				{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 20}.String(), Images: []string{"url1", "url2"}},
				{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 21}.String(), Images: []string{"url1"}},
			},
		},
		{
			description: "get all with account id = 5",
			input:       &models.Blog{AccountID: 5},
			output: []*models.Blog{
				{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
				{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 20}.String(), Images: []string{"url1", "url2"}},
				{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 21}.String(), Images: []string{"url1"}},
			},
		},
		{
			description: "get all with title = title5",
			input:       &models.Blog{Title: "title5"},
			output: []*models.Blog{
				{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
			},
		},
	}

	for i := range tests {
		mockBlogStore.EXPECT().GetAll(gomock.Any(), tests[i].input).
			Return(tests[i].output, nil)

		output, err := mockService.GetAll(ctx, tests[i].input)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, nil, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_GetAll_NilFilter(t *testing.T) {
	mockBlogStore, _, _, ctx, mockService := initializeTest(t)

	mockBlogStore.EXPECT().GetAll(gomock.Any(), &models.Blog{}).
		Return([]*models.Blog{
			{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
			{BlogID: "id4", AccountID: 2, Title: "title4", Summary: "summary4", Content: "content4", Tags: []string{"tag3", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 15}.String(), Images: []string{"url1"}},
			{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
			{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 20}.String(), Images: []string{"url1", "url2"}},
			{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 21}.String(), Images: []string{"url1"}},
		}, nil)

	tests := []struct {
		description string
		input       *models.Blog
		output      []*models.Blog
	}{
		{
			description: "get all with nil filter",
			input:       nil,
			output: []*models.Blog{
				{BlogID: "id5", AccountID: 3, Title: "title5", Summary: "summary5", Content: "content5", Tags: []string{"tag1"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
				{BlogID: "id4", AccountID: 2, Title: "title4", Summary: "summary4", Content: "content4", Tags: []string{"tag3", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 15}.String(), Images: []string{"url1"}},
				{BlogID: "id3", AccountID: 5, Title: "title3", Summary: "summary3", Content: "content3", Tags: []string{}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 16}.String(), Images: []string{"url1"}},
				{BlogID: "id2", AccountID: 5, Title: "title2", Summary: "summary2", Content: "content2", Tags: []string{"tag1", "tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 20}.String(), Images: []string{"url1", "url2"}},
				{BlogID: "id1", AccountID: 5, Title: "title1", Summary: "summary1", Content: "content1", Tags: []string{"tag2"}, CreatedOn: types.Date{Year: 2021, Month: 3, Day: 21}.String(), Images: []string{"url1"}},
			},
		},
	}

	for i := range tests {
		output, err := mockService.GetAll(ctx, tests[i].input)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, nil, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestBlog_GetAll_Error(t *testing.T) {
	mockBlogStore, _, _, ctx, mockService := initializeTest(t)

	mockBlogStore.EXPECT().GetAll(gomock.Any(), &models.Blog{Title: "title5"}).
		Return(nil, errors.DBError{})

	tests := []struct {
		description string
		input       *models.Blog
		output      []*models.Blog
		err         error
	}{
		{
			description: "get all with title = title5",
			input:       &models.Blog{Title: "title5"},
			output:      []*models.Blog(nil),
			err:         errors.DBError{},
		},
	}

	for i := range tests {
		output, err := mockService.GetAll(ctx, tests[i].input)

		if assert.Equal(t, len(tests[i].output), len(output), "TEST [%v], failed.\n%s", i+1, tests[i].description) {
			for j := range output {
				assert.Equal(t, tests[i].output[j], output[j], "TEST [%v], failed.\n%s", i+1, tests[i].description)
			}
		}

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}
