package blog

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

func initializeTest(t *testing.T) (*stores.MockBlog, *services.MockTag, *stores.MockImage, *app.Context, services.Blog) {
	ctrl := gomock.NewController(t)

	mockBlogStore := stores.NewMockBlog(ctrl)
	mockImageStore := stores.NewMockImage(ctrl)
	mockTagService := services.NewMockTag(ctrl)
	mockBlogService := New(mockBlogStore, mockTagService, mockImageStore)

	ctx := app.NewContext(nil, &app.App{Logger: log.NewLogger()})

	return mockBlogStore, mockTagService, mockImageStore, ctx, mockBlogService
}

func TestBlog_GetAll(t *testing.T) {
	mockBlogStore, _, _, ctx, mockBlogService := initializeTest(t)

	mockBlogStore.EXPECT().GetAll(gomock.Any(), &models.Blog{}, &models.Page{Limit: 3, PageNo: 1}).Return(getAllOutput(), nil).AnyTimes()

	mockBlogStore.EXPECT().GetAll(gomock.Any(), &models.Blog{}, &models.Page{Limit: 2, PageNo: 1}).Return(nil, errors.DBError{})

	tests := []struct {
		description string
		input       *models.Blog
		page        *models.Page
		output      []*models.Blog
		err         error
	}{
		{description: "get all with empty filter", input: &models.Blog{}, output: getAllOutput(), page: &models.Page{Limit: 3, PageNo: 1}},
		{description: "get all with nil filter", input: nil, output: getAllOutput(), page: &models.Page{Limit: 3, PageNo: 1}},
		{description: "database error", input: nil, page: &models.Page{Limit: 2, PageNo: 1}, err: errors.DBError{}},
	}

	for i, tc := range tests {
		output, err := mockBlogService.GetAll(ctx, tc.input, tc.page)

		if assert.Equal(t, len(tc.output), len(output), "TEST [%v], failed.\n%s", i+1, tc.description) {
			for j := range output {
				assert.Equal(t, tc.output[j], output[j], "TEST [%v], failed.\n%s", i+1, tc.description)
			}
		}

		assert.Equal(t, tc.err, err, "TEST [%v], failed.\n%s", i+1, tc.description)
	}
}

func TestBlog_GetAllByTagName(t *testing.T) {
	mockBlogStore, mockTagService, _, ctx, mockBlogService := initializeTest(t)

	mockTagService.EXPECT().Get(gomock.Any(), "#tag1").Return(&models.Tag{Name: "#tag1", BlogIDList: []string{"MSI8WKNSH9", "9SNVSH8K2M", "ABK7SH2V37"}}, nil)
	mockBlogStore.EXPECT().GetByIDs(gomock.Any(), []string{"MSI8WKNSH9", "9SNVSH8K2M", "ABK7SH2V37"}).Return(getAllOutput(), nil)

	mockTagService.EXPECT().Get(gomock.Any(), "#tag2").Return(nil, errors.DBError{})

	mockTagService.EXPECT().Get(gomock.Any(), "#tag3").Return(&models.Tag{Name: "#tag3", BlogIDList: []string{"MSI8WKNSH9", "9SNVSH8K2M", "ABK7SH2V37"}}, nil)
	mockBlogStore.EXPECT().GetByIDs(gomock.Any(), []string{"MSI8WKNSH9", "9SNVSH8K2M", "ABK7SH2V37"}).Return(nil, errors.DBError{})

	tests := []struct {
		description string
		input       string
		output      []*models.Blog
		err         error
	}{
		{description: "success case", input: "#tag1", output: getAllOutput(), err: nil},
		{description: "db error in call to tagService.Get", input: "#tag2", output: nil, err: errors.DBError{}},
		{description: "db error in call to tagService.Get", input: "#tag3", output: nil, err: errors.DBError{}},
	}

	for i, tc := range tests {
		output, err := mockBlogService.GetAllByTagName(ctx, tc.input)

		if assert.Equal(t, len(tc.output), len(output), "TEST [%v], failed.\n%s", i+1, tc.description) {
			for j := range output {
				assert.Equal(t, tc.output[j], output[j], "TEST [%v], failed.\n%s", i+1, tc.description)
			}
		}

		assert.Equal(t, tc.err, err, "TEST [%v], failed.\n%s", i+1, tc.description)
	}
}

func TestBlog_GetByID(t *testing.T) {
	mockBlogStore, _, _, ctx, mockBlogService := initializeTest(t)

	res := &models.Blog{
		BlogID:    "MSI8WKNSH9",
		AccountID: 2,
		Title:     "music",
		Summary:   "a blog on music",
		Content:   "avicii left :(",
		Tags:      []string{},
		CreatedOn: getTime("2021-05-23T15:04:05Z"),
		Images:    []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"},
	}

	mockBlogStore.EXPECT().Get(gomock.Any(), &models.Blog{BlogID: "MSI8WKNSH9"}).
		Return(res, nil)

	tests := []struct {
		description string
		id          string
		output      *models.Blog
		err         error
	}{
		{id: "", output: nil, err: errors.MissingParam{Param: "blog_id"}},
		{id: "MSI8WKNSH9", output: res, err: nil},
	}

	for i, tc := range tests {
		output, err := mockBlogService.GetByID(ctx, tc.id)

		assert.Equal(t, tc.output, output, "TEST [%v], failed.\n%s", i+1, tc.description)

		assert.Equal(t, tc.err, err, "TEST [%v], failed.\n%s", i+1, tc.description)
	}
}

//nolint:lll // hampers readability
func getAllOutput() []*models.Blog {
	return []*models.Blog{
		{
			BlogID:    "MSI8WKNSH9",
			AccountID: 2,
			Title:     "music",
			Summary:   "a blog on music",
			Content:   "avicii left :(",
			Tags:      []string{},
			CreatedOn: getTime("2021-05-23T15:04:05Z"),
			Images:    []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/avicii.jpg"},
		},
		{
			BlogID:    "9SNVSH8K2M",
			AccountID: 3,
			Title:     "flowers",
			Summary:   "a blog on flowers",
			Content:   "blue orchids are the most beautiful",
			Tags:      []string{"#love", "#nature", "#life"},
			CreatedOn: getTime("2021-04-16T15:04:05Z"),
			Images:    []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/orchids.jpg"},
		},
		{
			BlogID:    "ABK7SH2V37",
			AccountID: 2,
			Title:     "chocolate",
			Summary:   "a blog on chocolate",
			Content:   "bournville is the best chocolate!",
			Tags:      []string{"#cocoa", "#sweet", "#chocolate", "#trending"},
			CreatedOn: getTime("2021-04-10T15:04:05Z"),
			Images:    []string{"https://picshot-images.s3.ap-south-1.amazonaws.com/chocolate-gettyimages-473741340.jpg", "https://picshot-images.s3.ap-south-1.amazonaws.com/chocolate.jpg"},
		},
	}
}

func getTime(date string) time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05Z07:00", date)
	return t
}
