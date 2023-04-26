package domain

import (
	"github.com/brianvoe/gofakeit/v6"
)

func Mock() *Entity {
	return &Entity{
		Name: gofakeit.Name(),
		Email: gofakeit.Email(),
		Password: gofakeit.Word(),
	}
}