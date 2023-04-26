package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	d "github.com/gadc1996/reservation-manager/src/user/domain"
)

func TestGinCreate(t *testing.T) {
	router.POST("users", PostHandler(database))

	payload, _ := json.Marshal(d.Mock())
	request, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(payload))
	
	router.ServeHTTP(recorder, request)
}
