package db

import (
	"database/sql"
	"fmt"
	"server/config"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.AppConfig.DB.User,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.Name,
		config.AppConfig.DB.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
