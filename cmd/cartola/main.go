package main

import (
	"context"
	"database/sql"
	"fullcycle/internal/infra/db"
	"fullcycle/internal/infra/repository"
	"fullcycle/pkg/uow"
)

func main() {
	ctx := context.Background()

	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer dtb.Close()

	uow, err := uow.NewUow(ctx, dtb)

	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})
	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})
	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})
	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)

		return repo
	})
}
