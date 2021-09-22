package application

import (
	"stone.com/api/models"
	"stone.com/api/repository"
)

type accountApp struct {
	us repository.AccountRepository
}

var _ AccountAppInterface = &accountApp{}

type AccountAppInterface interface {
	GetAccounts() ([]models.Account, error)
	GetBalance(uint64) (*models.Account, error)
}

func (a *accountApp) GetAccounts() ([]models.Account, error) {
	return a.us.GetAccounts()
}

func (a *accountApp) GetBalance(accountId uint64) (*models.Account, error) {
	return a.us.GetBalance(accountId)
}
