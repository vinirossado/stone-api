package models

import "time"

type Account struct {
	ID        int `json:"id" gorm:"primary_key"`
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt time.Time `gorm:"->;<-:create"`
}
