package db

import (
	"database/sql"
	"github.com/NajmiddinAbdulhakim/config"

	_ "github.com/lib/pq"

)

func ConnectDB(cfg config.Config ) (*sql.DB, error)   {
	db, err := sql.Open("postgres", 
	fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database))
	)
}