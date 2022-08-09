package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/NajmiddinAbdulhakim/mybook/config"
	
	_ "github.com/lib/pq"
	
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
		),
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connecting migrate driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:/home/najmiddin/go/src/github.com/NajmiddinAbdulhakim/mybook/migrations",
		"postgres", driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connecting migrate :%v",err)
	}
	
	if err = m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}
	return db, nil
}
