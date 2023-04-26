package controllers

import (
	d "github.com/gadc1996/reservation-manager/src/user/domain"
	"github.com/gin-gonic/gin"
)

func IndexHandler(repository d.Repository) gin.HandlerFunc {
	data := map[string]interface{}{}
	fn := func(c *gin.Context) {
		entities, _ := repository.ListAll(data)
		c.JSON(200, entities)
	}

	return gin.HandlerFunc(fn)
}

