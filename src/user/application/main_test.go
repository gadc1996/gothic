package application

import (
	r "github.com/gadc1996/reservation-manager/src/user/infrastructure/repositories/memory"
	d "github.com/gadc1996/reservation-manager/src/user/infrastructure/interactors/database/memory"
)

var repository *r.Repository
var database *d.Interactor

func init() {
	repository = &r.Repository{}
	database = &d.Interactor{}
}
