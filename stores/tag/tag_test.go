package tag

import (
	"context"
	"testing"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
	"github.com/stretchr/testify/assert"
)

func initializeTest() (*app.Context, stores.Tag) {
	app.InitializeTestTagsCollection(a.Mongo.DB(), a.Logger, "../../db")
	return &app.Context{Context: context.TODO(), App: a}, New()
}

func TestTag_Get(t *testing.T) {
	ctx, tag := initializeTest()

	tests := []struct {
		description string
		input       string
		output      *models.Tag
		err         error
	}{
		{
			description: "get tag with valid name",
			input:       "#trending",
			output:      &models.Tag{Name: "#trending", BlogIDList: []string{"UMS672XR8J", "ABK7SH2V37", "MSI8NS2909", "POQA7B2J7X"}},
		},
		{
			description: "get tag with invalid name",
			input:       "#abc",
			err:         errors.DBError{Err: mongo.ErrNoDocuments},
		},
	}

	for i := range tests {
		output, err := tag.Get(ctx, tests[i].input)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestTag_Get_Error(t *testing.T) {
	ctx, tag := initializeTest()

	type demo struct {
		Name       string `bson:"_id"`
		BlogIDList []int  `bson:"blog_id_list"`
	}

	collection := ctx.Mongo.DB().Collection("tags")

	_, _ = collection.InsertOne(ctx, demo{Name: "dummy", BlogIDList: []int{1, 2, 3, 4}})

	tc := struct {
		description string
		input       string
		output      *models.Tag
		err         error
	}{
		description: "get tag: failure case (decode error)",
		input:       "dummy",
		output:      nil,
		err:         errors.DBError{},
	}

	output, err := tag.Get(ctx, tc.input)

	assert.Equal(t, tc.output, output, "TEST, failed.\n%s", tc.description)

	assert.IsType(t, tc.err, err, "TEST, failed.\n%s", tc.description)
}

func TestTag_Update(t *testing.T) {
	ctx, tag := initializeTest()

	tests := []struct {
		description string
		blogID      string
		tag         string
		operation   constants.Operation
		output      *models.Tag
		err         error
	}{
		{
			description: "update (add) tag valid operation",
			blogID:      "TEST_ID",
			tag:         "#trending",
			operation:   constants.Add,
			output:      &models.Tag{Name: "#trending", BlogIDList: []string{"UMS672XR8J", "ABK7SH2V37", "MSI8NS2909", "POQA7B2J7X", "TEST_ID"}},
		},
		{
			description: "update (remove) tag valid operation",
			blogID:      "TEST_ID",
			tag:         "#trending",
			operation:   constants.Remove,
			output:      &models.Tag{Name: "#trending", BlogIDList: []string{"UMS672XR8J", "ABK7SH2V37", "MSI8NS2909", "POQA7B2J7X"}},
		},
		{
			description: "update (add) tag on non-existing tag",
			blogID:      "TEST_ID",
			tag:         "#invalidtag",
			operation:   constants.Add,
			err:         errors.DBError{Err: mongo.ErrNoDocuments},
		},
		{
			description: "update (remove) tag on non-existing tag",
			blogID:      "TEST_ID",
			tag:         "#invalidtag",
			operation:   constants.Remove,
			err:         errors.DBError{Err: mongo.ErrNoDocuments},
		},
	}

	for i := range tests {
		err := tag.Update(ctx, tests[i].blogID, tests[i].tag, tests[i].operation)

		output, _ := tag.Get(ctx, tests[i].tag)

		assert.Equal(t, tests[i].output, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestTag_Create(t *testing.T) {
	ctx, tag := initializeTest()

	tests := []struct {
		description string
		input       *models.Tag
		err         error
	}{
		{
			description: "create tag: success case",
			input:       &models.Tag{Name: "#test_tag", BlogIDList: []string{"TEST_ID"}},
		},
		{
			description: "create tag: failure case (redundant input)",
			input:       &models.Tag{Name: "#test_tag", BlogIDList: []string{"TEST_ID"}},
			err:         errors.DBError{},
		},
		{
			description: "create tag: failure case (nil input)",
			input:       nil,
			err:         errors.DBError{},
		},
		{
			description: "create empty tag",
			input:       &models.Tag{},
			err:         nil,
		},
	}

	for i := range tests {
		err := tag.Create(ctx, tests[i].input)

		var output *models.Tag

		if t := tests[i].input; t != nil {
			output, _ = tag.Get(ctx, t.Name)
		}

		assert.Equal(t, tests[i].input, output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.IsType(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}

func TestTag_Delete(t *testing.T) {
	ctx, tag := initializeTest()

	tests := []struct {
		description string
		input       string
		err         error
	}{
		{
			description: "delete tag with valid name",
			input:       "#trending",
		},
		{
			description: "delete tag with invalid name",
			input:       "#abc",
			err:         errors.DBError{Err: mongo.ErrNoDocuments},
		},
	}

	for i := range tests {
		err := tag.Delete(ctx, tests[i].input)

		output, _ := tag.Get(ctx, tests[i].input)

		assert.Equal(t, (*models.Tag)(nil), output, "TEST [%v], failed.\n%s", i+1, tests[i].description)

		assert.Equal(t, tests[i].err, err, "TEST [%v], failed.\n%s", i+1, tests[i].description)
	}
}
