package controllers

import (
	application "github.com/gadc1996/reservation-manager/src/user/application"
	d "github.com/gadc1996/reservation-manager/src/user/domain"
	"github.com/gin-gonic/gin"
)

func GetHandler(repository d.Repository) gin.HandlerFunc {
	var data *d.Entity
	fn := func(c *gin.Context) {
		user, _ := application.FindById(data.Id, repository)
		c.JSON(200, user)
	}

	return gin.HandlerFunc(fn)
}
