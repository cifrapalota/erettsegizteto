package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB represents a connection to a database
type DB struct {
	gormDB *gorm.DB
}

// NewDB creates a new DB connection
func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "db_erettsegizteto.", // schema name
		},
	})
	if err != nil {
		return nil, err
	}

	return &DB{gormDB: db}, nil
}
