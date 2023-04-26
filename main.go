package main

import (
	"github.com/gadc1996/reservation-manager/architecture/router"
	db "github.com/gadc1996/reservation-manager/src/user/infrastructure/interactors/database/mysql"
	repo "github.com/gadc1996/reservation-manager/src/user/infrastructure/repositories/memory"
)

func main() {
	database := &db.Interactor{}
	repository := &repo.Repository{}

	router.Setup(repository, database)
}