package di

import (
	postgresDb "authentication-service/infra/database/postgres"
	pgUserRepo "authentication-service/infra/database/postgres/repositories/userRepo"
	"authentication-service/infra/database/repositories/userRepo"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
)

func (d *DI) InjectPostgres() {
	d.Dig.Provide(postgresDb.New(), dig.As(new(sqlx.DB)))

	d.Dig.Invoke(func(db sqlx.DB) {

		driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
		if err != nil {
			panic(err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://infra/database/postgres/migrations",
			"postgres", driver)

		if err != nil {
			log.Fatalf("erro ao atualizar o dicionario: %v", err)
		}

		if err = m.Up(); err != nil {
			if err.Error() != "no change" {
				log.Fatalf("erro ao atualizar o dicionario: %v", err)
			}
		}
	})
}

func (d *DI) CloseDB() {
	err := d.Dig.Invoke(func(db sqlx.DB) {
		if err := db.Close(); err != nil {
			panic(err)
		}
	})

	if err != nil {
		panic(err)
	}
}

func (d *DI) InjectPostgresRepos() {
	d.Dig.Provide(pgUserRepo.New, dig.As(new(userRepo.IUserRepo)))
}
