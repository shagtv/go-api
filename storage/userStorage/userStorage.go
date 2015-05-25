package userStorage

import (
	"github.com/shagtv/go-api/library/db"
	"github.com/shagtv/go-api/models"
)

func CreateTable() {
	db.Connection().CreateTable(&models.User{})
}
