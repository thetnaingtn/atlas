package db

import (
	"atlas/internal/config"
	"atlas/store"
	"atlas/store/db/sqlite"
)

func NewDBDriver(config *config.Config) (store.Driver, error) {
	db, err := sqlite.NewDB(config)
	if err != nil {
		return nil, err
	}

	return db, nil
}
