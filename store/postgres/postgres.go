package postgres

import (
	"webot/config"

	"github.com/go-pg/pg/v9"
)

// postgres client struct
type Client struct {
	*pg.DB
}

//creates new client for postgres DB
func NewClient(cfg config.PostgresConfig) (*Client, error) {
	db := pg.Connect(&pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Addr:     cfg.Address,
		Database: cfg.Db,
		PoolSize: cfg.Poolsize,
	})
	return &Client{db}, nil
}
