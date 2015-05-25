package models

import (
	"time"
)

type User struct {
	ID        int `sql:"AUTO_INCREMENT"`
	Birthday  time.Time
	Name      string `sql:"size:255"`
	Email     string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
