package application

import (
	"stone.com/api/models"
	"stone.com/api/repository"
)

type accountApp struct {
	us repository.AccountRepo
}

var _ AccountAppInterface = &accountApp{}

type AccountAppInterface interface {
	GetAccounts() ([]models.Account, error)
}

func (a *accountApp) GetAccounts() ([]models.Account, error) {
	return a.us.GetAccounts()
}
