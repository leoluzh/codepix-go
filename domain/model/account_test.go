package model_test

import (
	"testing"

	"github.com/leoluzh/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {

	code := "013"
	name := "Caixa Economica Federal"
	bank, err := model.NewBank(code, name)

	accountNumber := ""
	ownerName := "leoluzh"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankID, bank.ID)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)

}
