package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"international_trade/config"
)

const (
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second

	
)


type Postgres struct {
	connAttempts int
	connTimeout  time.Duration
}

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {

	pg := &Postgres{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PG.Host, cfg.PG.Port, cfg.PG.Username, cfg.PG.DBName, cfg.PG.Password, cfg.PG.SSLMode))

	if err != nil {
		return nil, err
	}

	for pg.connAttempts > 0 {

		err = db.Ping()

		if err == nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)
		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}
	
	return db, nil
}
