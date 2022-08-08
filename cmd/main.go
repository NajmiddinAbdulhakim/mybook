package main

import (
	"log"

	"github.com/NajmiddinAbdulhakim/mybook/api"
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

	service := service.NewService(storage.NewStoragePg(db))

	server := api.NewRouter(api.Option{

		Conf:           cfg,
		ServiceManager: service,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal(`failed to run http server: `, err)
		panic(err)
	}

}
