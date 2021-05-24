package models

import (
	"time"
)

type Stocks struct {
	Id uint64 `json:"id"`
	Status string `json:"status"`
	Jumlah int `json:"jumlah"`
	TglOrder time.Time `json:"tglorder"`
	BarangId uint64 `json:"barangid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}
