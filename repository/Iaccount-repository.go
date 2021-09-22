package repository

import "stone.com/api/models"

type AccountRepository interface {
	// SaveUser(*models.Account) (*models.Account, map[string]string)
	GetAccounts() ([]models.Account, error)
	GetBalance(uint64) (*models.Account, error)
}
