package repository

import (
	"fmt"
	"github.com/leoluzh/codepix-go/domain/model"
	"gorm.io/gorm"
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

func (r *TransactionRepositoryDb) Save(transaction *Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepositoryDb) FindBy(id string) (*Transaction, error) {
	var transaction model.Transaction
	r.Db.Preload("AccountFrom.Bank").First(&transaction, " id = ? ", id)
	if transaction.ID == "" {
		return nil, fmt.Errorf("Transaction not found")
	}
	return &transaction, nil
}
