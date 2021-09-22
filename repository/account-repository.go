package repository

import (
	"errors"

	"gorm.io/gorm"

	"stone.com/api/models"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db}
}

var _ AccountRepository = &AccountRepo{}

func (r *AccountRepo) GetAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := r.db.Debug().Find(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *AccountRepo) GetBalance(id uint64) (*models.Account, error) {
	var account models.Account

	err := r.db.Debug().Where("id = ?", id).Take(&account).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &account, nil
}

// func (r *UserRepo) GetUser(id uint64) (*entity.User, error) {
// 	var user entity.User
// 	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	if gorm.IsRecordNotFoundError(err) {
// 		return nil, errors.New("user not found")
// 	}
// 	return &user, nil
// }
