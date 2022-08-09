package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/NajmiddinAbdulhakim/mybook/config"
	"github.com/NajmiddinAbdulhakim/mybook/pkg/db"
)

var aRepo *AuthorRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	db, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	aRepo = NewAuthorRepo(db)
	
	os.Exit(m.Run())
}
