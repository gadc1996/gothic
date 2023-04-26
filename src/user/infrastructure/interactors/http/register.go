package http

import (
	"github.com/gin-gonic/gin"
	c "github.com/gadc1996/reservation-manager/src/user/infrastructure/interactors/http/controllers"
	d "github.com/gadc1996/reservation-manager/src/user/domain"
)

func Register(router *gin.Engine, repository d.Repository, database d.DatabaseInteractor) {
	router.GET("users", c.IndexHandler(repository))
	router.GET("users/:id", c.GetHandler(repository))

	router.POST("users", c.PostHandler(database))
}