package router

import (
	"github.com/gin-gonic/gin"

	d "github.com/gadc1996/reservation-manager/src/user/domain"
	user "github.com/gadc1996/reservation-manager/src/user/infrastructure/interactors/http"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func Setup(repository d.Repository,database d.DatabaseInteractor) {
	defer router.Run()

	user.Register(router, repository, database)
}