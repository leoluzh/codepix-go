package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull" gorm:"column:owner_name; type:varchar(255); not null"`
	Bank *bank `valid:"-"`
	BankID string `gorm:"column:bank_id; type:uuid; not null" valid:"-"`
	Number string `json:"number" valid:"notnull" gorm:"type:varchar(20); not null"`
	PixKeys []*PixKey `valid:"-" gorm:"ForeingKey:AccountID"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount( bank *Bank , number string, ownerName string ) (*Account, error) {
	
	account := Account{
		OwnerName: ownerName,
		Bank: bank,
		Number: number
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	} 

	return &account

}
