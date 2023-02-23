package database

import (
	"github.com/go-pg/pg"

	"github.com/madnaaaaas/crud/pkg/config"
)

func NewDatabaseConnection(c *config.Config) (*pg.DB, error) {
	return pg.Connect(&pg.Options{
		Addr:     c.PGAddress,
		User:     c.PGUser,
		Password: c.PGPassword,
		Database: c.PGDatabase,
	}), nil

}
