package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull" gorm:"column:kind; type:varchar(20); not null"`
	Key       string   `json:"key" valid:"notnull" gorm:"column:key; type:varchar(20); not null"`
	AccountID string   `gorm:"column:account_id; type:uuid; not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull" gorm:"column:status;type:varchar(20);not null"`
}

func (pixKey *PixKey) isValid() error {

	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "telefone" {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Kind != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {

	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		AccountID: account.ID,
		Account:   account,
		Status:    "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &pixKey, nil

}
