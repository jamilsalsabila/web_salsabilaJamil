package models

import (
	"time"
)

type Users struct {
	Id uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nama string `json:"nama"`
	Email string `json:"email"`
	NoHP string `json:"nohp"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
