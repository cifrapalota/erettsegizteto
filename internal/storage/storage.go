package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "db_erettsegizteto.", // schema name
		},
	})
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}
