package models

import "time"

type Transfer struct {
	ID                     string
	Account_Origin_ID      int
	Account_Destination_ID int
	Amount                 float64
	Created_At             time.Time
}
