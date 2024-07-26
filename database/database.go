package database

import (
	"database/sql"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

var (
	client     *sql.DB
	clientOnce sync.Once
)

func Client() *sql.DB {
	clientOnce.Do(
		func() {
			conn, err := mysql.NewConnector(
				&mysql.Config{
					User:                 Config().User,
					Addr:                 Config().Host,
					Passwd:               Config().Pass,
					DBName:               Config().Name,
					AllowNativePasswords: true,
				},
			)
			if err != nil {
				log.Fatal().Err(err).Msg("Error creating MySQL connector")
			}
			client = sql.OpenDB(conn)
		},
	)
	return client
}
