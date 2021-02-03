package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/leoluzh/codepix-go/domain/model"
)

// type PixKeyRepository {
// 	RegisterKey( pixKey *PixKey) (*PixKey, error)
// 	FindKeyByKind( key string , kind string) (*PixKey, error)
// 	AddBank( bank *Bank) error
// 	AddAccount( account *Account ) error
// 	FindAccount( id string) (*Account,error)
// }

type PixKeyRepositoryDb struct {
	//db connection
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {

	var pixKey model.PixKey
	//preload - same concept of fetch eager - hibernate
	r.Db.Preload("Account.Bank").First(&pixKey, " kind = ? AND key = ?", kind, key)

	//verify if return result
	if pixKey.ID == "" {
		return nil, fmt.Errorf("Key not found")
	}

	return &pixKey, nil

}

func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	//preload
	r.Db.Preload("Bank").First(&account, " id = ? ", id)

	if account.ID == "" {
		return nil, fmt.Errorf("Account not found")
	}

	return &account, nil

}

func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	//preload
	r.Db.Preload("").First(&bank, " id = ? ", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("Bank not found")
	}

	return &bank, nil

}
