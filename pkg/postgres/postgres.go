package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/satriowisnugroho/book-store/internal/config"
)

// Postgres holds database connection to postgreSQL
type Postgres struct {
	Db *sqlx.DB
}

// NewPostgres initializes postgres database connection from configs
func NewPostgres(opt *config.DatabaseConfig) (*Postgres, error) {
	postgresDb, err := sqlx.Open(opt.Driver, fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", opt.Username, opt.Password, opt.Name, opt.Host, opt.Port))
	if err != nil {
		return &Postgres{}, err
	}

	postgresDb.SetMaxOpenConns(opt.Pool)

	return &Postgres{Db: postgresDb}, nil
}
