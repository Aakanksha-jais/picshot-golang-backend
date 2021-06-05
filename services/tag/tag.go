package tag

import (
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Aakanksha-jais/picshot-golang-backend/models"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/constants"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/Aakanksha-jais/picshot-golang-backend/services"
	"github.com/Aakanksha-jais/picshot-golang-backend/stores"
)

type tag struct {
	store stores.Tag
}

func New(store stores.Tag) services.Tag {
	return tag{store: store}
}

func (t tag) Get(c *app.Context, name string) (*models.Tag, error) {
	tag, err := t.store.Get(c, name)

	switch err := err.(type) {
	case errors.DBError:
		if err.Err == mongo.ErrNoDocuments {
			return nil, errors.EntityNotFound{Entity: "tag", ID: name}
		}

		return nil, err
	default:
		return tag, nil
	}
}

func (t tag) AddBlogID(c *app.Context, blogID string, tags []string) {
	for i := range tags {
		// validate tag name
		if !validateTag(tags[i]) {
			c.Logger.Errorf("invalid tag %s", tags[i])
			continue
		}

		_, err := t.Get(c, tags[i])

		switch err.(type) {
		case errors.EntityNotFound:
			// create tag if it does not exist already
			if err = t.store.Create(c, &models.Tag{Name: tags[i], BlogIDList: []string{blogID}}); err != nil {
				c.Logger.Errorf("cannot create tag %s: %s", tags[i], err.Error())
			}
		case nil:
			// update tag if it exists
			if err = t.store.Update(c, blogID, tags[i], constants.Add); err != nil {
				c.Logger.Errorf("cannot update tag %s: %s", tags[i], err.Error())
			}
		default:
			c.Logger.Errorf("cannot find tag %s: %s", tags[i], err.Error())
		}
	}
}

//nolint:gocognit // hampers readability of code
func (t tag) RemoveBlogID(c *app.Context, blogID string, tags []string) {
	for i := range tags {
		// validate tag name
		if !validateTag(tags[i]) {
			c.Logger.Errorf("invalid tag %s", tags[i])
			continue
		}

		_, err := t.Get(c, tags[i])
		if err != nil {
			c.Logger.Errorf("cannot find tag %s: %s", tags[i], err.Error())
			continue
		}

		// update tag if it exists
		if err = t.store.Update(c, blogID, tags[i], constants.Remove); err != nil {
			c.Logger.Errorf("cannot update tag %s: %s", tags[i], err.Error())
			continue
		}

		tag, err := t.Get(c, tags[i])
		if err != nil {
			continue
		}

		if len(tag.BlogIDList) == 0 {
			if err := t.store.Delete(c, tag.Name); err != nil {
				c.Logger.Errorf("cannot remove tag %s: %s", tags[i], err.Error())
			}
		}
	}
}

func validateTag(name string) bool {
	name = strings.TrimSpace(name)

	if len(name) < 1 {
		return false
	}

	if !strings.HasPrefix(name, "#") {
		return false
	}

	_, err := regexp.MatchString(`^[0-9A-Za-z_]$`, name[1:])

	return err == nil
}
