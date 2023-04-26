package controllers

import (
	db "github.com/gadc1996/reservation-manager/src/user/infrastructure/interactors/database/mysql"
	r "github.com/gadc1996/reservation-manager/src/user/infrastructure/repositories/memory"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

var router *gin.Engine
var recorder *httptest.ResponseRecorder

var database *db.Interactor
var repository *r.Repository

func init() {
	router = gin.Default()
	recorder = httptest.NewRecorder()

	database = &db.Interactor{}
	repository = &r.Repository{}
}
