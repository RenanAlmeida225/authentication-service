package postgres

import (
	"authentication-service/utils"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func New() *sqlx.DB {
	dbConfig := utils.Env.DBConf
	db, err := sqlx.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}

	return db
}
