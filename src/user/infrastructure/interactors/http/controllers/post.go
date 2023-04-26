package controllers

import (
	"github.com/gadc1996/reservation-manager/src/user/application"
	d "github.com/gadc1996/reservation-manager/src/user/domain"
	"github.com/gin-gonic/gin"
)

func PostHandler(database d.DatabaseInteractor) gin.HandlerFunc {
	var entity d.Entity
	fn := func(c *gin.Context) {
		c.BindJSON(entity)
		user, _ := application.Save(&entity, database)
		c.JSON(200, user)
	}

	return gin.HandlerFunc(fn)
}
