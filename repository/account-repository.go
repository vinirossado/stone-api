package repository

import (
	"gorm.io/gorm"

	"stone.com/api/models"
)

type AccountRepo struct {
	db *gorm.DB
}

type AccountRepoInterface interface {
	NewAccount(*models.Account) (*models.Account, map[string]string)
	GetAccounts() ([]*models.Account, error)
	GetBalance(uint64) (*models.Account, error)
}

func NewAccountRepository(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db}
}

func (r *AccountRepo) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := r.db.Debug().Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	// if gorm.err(err) {
	// 	return nil, errors.New("items not found")
	// }
	return accounts, nil
}
