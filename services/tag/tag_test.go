package tag

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

func initializeTest(t *testing.T) (*stores.MockTag, services.Tag, *app.Context) {
	ctrl := gomock.NewController(t)
	mockStore := stores.NewMockTag(ctrl)
	mockService := New(mockStore)
	ctx := &app.Context{Context: context.TODO(), App: &app.App{Logger: log.NewLogger()}}

	return mockStore, mockService, ctx
}

// nolint:lll // hampers readability (needed for test cases)
func TestTag_Get(t *testing.T) {
	mockStore, mockService, ctx := initializeTest(t)

	mockStore.EXPECT().Get(gomock.Any(), "#trending").Return(&models.Tag{Name: "#trending", BlogIDList: []string{"UMS672XR8J", "ABK7SH2V37", "MSI8NS2909", "POQA7B2J7X"}}, nil)
	mockStore.EXPECT().Get(gomock.Any(), "#abc").Return(nil, errors.DBError{Err: mongo.ErrNoDocuments})
	mockStore.EXPECT().Get(gomock.Any(), "#def").Return(nil, errors.DBError{})

	tests := []struct {
		description string
		input       string
		output      *models.Tag
		err         error
	}{
		{description: "get tag (database error)", input: "#def", err: errors.DBError{}},
		{description: "get tag that does not exist", input: "#abc", err: errors.EntityNotFound{Entity: "tag", ID: "#abc"}},
		{description: "get tag with valid name", input: "#trending", output: &models.Tag{Name: "#trending", BlogIDList: []string{"UMS672XR8J", "ABK7SH2V37", "MSI8NS2909", "POQA7B2J7X"}}},
	}

	for i, tc := range tests {
		output, err := mockService.Get(ctx, tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tc.description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tc.description)
	}
}

// nolint:lll,gocognit,dupl // hampers readability (needed for test cases)
func TestTag_AddBlogID(t *testing.T) {
	mockStore, mockService, ctx := initializeTest(t)

	registerAddMockCalls(mockStore)

	tests := []struct {
		description string
		blogID      string
		tags        []string
		output      []string
	}{
		{description: "success case", blogID: "TEST_ID", tags: []string{"#tag5"}},
		{description: "invalid tag", blogID: "TEST_ID", tags: []string{"demo-tag"}, output: []string{"invalid tag demo-tag"}},
		{description: "create error", blogID: "TEST_ID", tags: []string{"#tag1"}, output: []string{"cannot create tag #tag1"}},
		{description: "create and invalid tag error", blogID: "TEST_ID", tags: []string{"#tag1", "tag2"}, output: []string{"cannot create tag #tag1", "invalid tag tag2"}},
		{description: "update, create and invalid tag error", blogID: "TEST_ID", tags: []string{"#tag3", "#tag1", "tag2"}, output: []string{"cannot update tag #tag3", "cannot create tag #tag1", "invalid tag tag2"}},
		{description: "update, create, invalid tag error and find error", blogID: "TEST_ID", tags: []string{"#tag3", "#tag1", "tag2", "#tag4"}, output: []string{"cannot update tag #tag3", "cannot create tag #tag1", "invalid tag tag2", "cannot find tag #tag4"}},
	}

	for i, tc := range tests {
		b := new(bytes.Buffer)
		ctx.Logger = log.NewMockLogger(b)

		mockService.AddBlogID(ctx, tc.blogID, tc.tags)

		if len(tc.output) == 0 {
			if b.String() != "" {
				t.Errorf("TEST [%v], failed.\n%s\nExpected empty log.\nGot:%v", i+1, tc.description, b.String())
			}

			continue
		}

		for _, out := range tc.output {
			if !strings.Contains(b.String(), out) {
				t.Errorf("TEST [%v], failed.\n%s\nExpected: %v (in logs)\nGot:%v", i+1, tc.description, out, b.String())
			}
		}
	}
}

func registerAddMockCalls(mockStore *stores.MockTag) {
	mockStore.EXPECT().Get(gomock.Any(), "#tag1").Return(nil, errors.DBError{Err: mongo.ErrNoDocuments}).AnyTimes()
	mockStore.EXPECT().Create(gomock.Any(), &models.Tag{Name: "#tag1", BlogIDList: []string{"TEST_ID"}}).AnyTimes().Return(errors.DBError{})

	mockStore.EXPECT().Get(gomock.Any(), "#tag3").Return(&models.Tag{Name: "#tag3", BlogIDList: []string{"TAG_ID"}}, nil).AnyTimes()
	mockStore.EXPECT().Update(gomock.Any(), "TEST_ID", "#tag3", constants.Add).AnyTimes().Return(errors.DBError{})

	mockStore.EXPECT().Get(gomock.Any(), "#tag4").Return(nil, errors.DBError{})

	mockStore.EXPECT().Get(gomock.Any(), "#tag5").Return(&models.Tag{Name: "#tag5", BlogIDList: []string{"TAG_ID"}}, nil)
	mockStore.EXPECT().Update(gomock.Any(), "TEST_ID", "#tag5", constants.Add).AnyTimes().Return(nil)
}

// nolint:lll,gocognit,dupl // hampers readability (needed for test cases)
func TestTag_RemoveBlogID(t *testing.T) {
	mockStore, mockService, ctx := initializeTest(t)

	registerRemoveMockCalls(mockStore)

	tests := []struct {
		description string
		blogID      string
		tags        []string
		output      []string
	}{
		{description: "success case", blogID: "TEST_ID", tags: []string{"#tag4"}, output: []string(nil)},
		{description: "invalid tag", blogID: "TEST_ID", tags: []string{""}, output: []string{"invalid tag "}},
		{description: "get tag after update returns error", blogID: "TEST_ID", tags: []string{"#tag6"}, output: []string(nil)},
		{description: "get returns error", blogID: "TEST_ID", tags: []string{"#tag1"}, output: []string{"cannot find tag #tag1"}},
		{description: "delete returns error", blogID: "TEST_ID", tags: []string{"#tag5"}, output: []string{"cannot remove tag #tag5"}},
		{description: "get tag returns error, invalid tag name", blogID: "TEST_ID", tags: []string{"#tag1", "tag2"}, output: []string{"cannot find tag #tag1", "invalid tag tag2"}},
		{description: "update and find return error, invalid tag name", blogID: "TEST_ID", tags: []string{"#tag3", "#tag1", "tag2"}, output: []string{"cannot update tag #tag3", "cannot find tag #tag1", "invalid tag tag2"}},
	}

	for i, tc := range tests {
		b := new(bytes.Buffer)
		ctx.Logger = log.NewMockLogger(b)

		mockService.RemoveBlogID(ctx, tc.blogID, tc.tags)

		if len(tc.output) == 0 {
			if b.String() != "" {
				t.Errorf("TEST [%v], failed.\n%s\nExpected empty log.\nGot:%v", i+1, tc.description, b.String())
			}

			continue
		}

		for _, out := range tc.output {
			if !strings.Contains(b.String(), out) {
				t.Errorf("TEST [%v], failed.\n%s\nExpected: %v (in logs)\nGot:%v", i+1, tc.description, out, b.String())
			}
		}
	}
}

func registerRemoveMockCalls(mockStore *stores.MockTag) {
	mockStore.EXPECT().Get(gomock.Any(), "#tag1").Return(nil, errors.DBError{Err: mongo.ErrNoDocuments}).AnyTimes()

	mockStore.EXPECT().Get(gomock.Any(), "#tag3").Return(&models.Tag{Name: "#tag3", BlogIDList: []string{"TAG_ID"}}, nil).AnyTimes()
	mockStore.EXPECT().Update(gomock.Any(), "TEST_ID", "#tag3", constants.Remove).AnyTimes().Return(errors.DBError{})

	mockStore.EXPECT().Get(gomock.Any(), "#tag4").Return(&models.Tag{Name: "#tag4", BlogIDList: []string{"TEST_ID"}}, nil)
	mockStore.EXPECT().Update(gomock.Any(), "TEST_ID", "#tag4", constants.Remove).Return(nil)
	mockStore.EXPECT().Get(gomock.Any(), "#tag4").Return(&models.Tag{Name: "#tag4", BlogIDList: []string(nil)}, nil)
	mockStore.EXPECT().Delete(gomock.Any(), "#tag4").Return(nil)

	mockStore.EXPECT().Get(gomock.Any(), "#tag5").Return(&models.Tag{Name: "#tag5", BlogIDList: []string{"TEST_ID"}}, nil)
	mockStore.EXPECT().Update(gomock.Any(), "TEST_ID", "#tag5", constants.Remove).Return(nil)
	mockStore.EXPECT().Get(gomock.Any(), "#tag5").Return(&models.Tag{Name: "#tag5", BlogIDList: []string(nil)}, nil)
	mockStore.EXPECT().Delete(gomock.Any(), "#tag5").Return(errors.DBError{})

	mockStore.EXPECT().Get(gomock.Any(), "#tag6").Return(&models.Tag{Name: "#tag6", BlogIDList: []string{"TEST_ID", "TEST_ID1"}}, nil)
	mockStore.EXPECT().Update(gomock.Any(), "TEST_ID", "#tag6", constants.Remove).Return(nil)
	mockStore.EXPECT().Get(gomock.Any(), "#tag6").Return(nil, errors.DBError{})
}
