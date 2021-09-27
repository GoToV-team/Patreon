package sqlstore

import (
	"patreon/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatorRepository_Create(t *testing.T) {
	db, teardown := TestDB(t, dbUrl)
	defer teardown("users", "creator_profile")

	s := New(db)
	cr := models.TestCreator(t)
	u := models.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, cr.ID)
	err = s.Creator().Create(cr)
	assert.NoError(t, err)
}