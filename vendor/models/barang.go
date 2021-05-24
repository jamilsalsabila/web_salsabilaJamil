package models

import (
	"time"
)

type Barang struct {
	Id uint64 `json:"id"`
	Nama string `json:"nama"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
