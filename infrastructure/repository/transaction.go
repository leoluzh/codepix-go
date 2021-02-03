package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/leoluzh/codepix-go/domain/model"
)

// type TransactionRepository interface {
// 	Register(transaction *Transaction) error
// 	Save(transaction *Transaction) error
// 	FindBy( id string ) (*Transaction,error)
// }

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (r *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryDb) FindBy(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	r.Db.Preload("AccountFrom.Bank").First(&transaction, " id = ? ", id)
	if transaction.ID == "" {
		return nil, fmt.Errorf("Transaction not found")
	}
	return &transaction, nil
}
