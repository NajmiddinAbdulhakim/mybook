package postgres

import (
	"fmt"
	"log"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/huandu/go-assert"
	"github.com/stretchr/testify/require"

	"github.com/NajmiddinAbdulhakim/mybook/models"
)

func TestAuthors(t *testing.T) {
	t.Run(`CRAD full`, func(t *testing.T) {
		t.Cleanup(cleanUpFc(aRepo))

		uuid, err := uuid.NewV4()
		id := uuid.String()
		require.NoError(t, err)

		want, err := aRepo.CreateAuthor(&models.Author{
			Id:        id,
			FirstName: "John",
			LastName:  "Karl",
			PhotoLink: `1.com`,
		})
		require.NoError(t, err)

		got, err := aRepo.GetAuthor(id)
		require.NoError(t, err)

		assert.Equal(t, want, got)

		want2, err := aRepo.UpdateAuthor(&models.Author{
			Id:        id,
			FirstName: "Adam",
			LastName:  "Fred",
			PhotoLink: `1231.com`,
		})
		require.NoError(t, err)

		got2, err := aRepo.GetAuthor(id)
		require.NoError(t, err)

		assert.Equal(t, want2, got2)

		err = aRepo.DeleteAuthor(id)
		require.NoError(t, err)
	})

	t.Run(`Check error`, func(t *testing.T) {
		t.Cleanup(cleanUpFc(aRepo))

		uuid, err := uuid.NewV4()
		id := uuid.String()
		fmt.Println(id)
		require.NoError(t, err)

		want, err := aRepo.CreateAuthor(&models.Author{
			Id:        "",
			FirstName: "John",
			LastName:  "Karl",
			PhotoLink: `1.com`,
		})
		require.Error(t, err)

		got, err := aRepo.GetAuthor(id)
		require.Error(t, err)

		assert.Equal(t, want, got)

		want2, err := aRepo.UpdateAuthor(&models.Author{
			Id:        id,
			FirstName: "Adam",
			LastName:  "Fred",
			PhotoLink: `1231.com`,
		})
		require.Error(t, err)

		got2, err := aRepo.GetAuthor(id)
		require.Error(t, err)

		assert.Equal(t, want2, got2)

		err = aRepo.DeleteAuthor(id)
		require.Error(t, err)

	})

}

func cleanUpFc(r *AuthorRepo) func() {
	return func() {
		if err := r.cleanTable(); err != nil {
			log.Println(`CLEANUP OF DB FAILED`)
		}
	}

}
