package models

import "time"

type Account struct {
	ID        int
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt time.Time
}
