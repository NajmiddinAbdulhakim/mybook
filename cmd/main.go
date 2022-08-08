package main

import (
	"fmt"

	"github.com/NajmiddinAbdulhakim/mybook/config"
	"github.com/NajmiddinAbdulhakim/mybook/pkg/db"
	"github.com/NajmiddinAbdulhakim/mybook/service"
	"github.com/NajmiddinAbdulhakim/mybook/storage"
)

func main() {
	cfg := config.Load()

	db, err := db.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	server := service.NewService(storage.NewStoragePg(db))
	fmt.Print(server)
}
