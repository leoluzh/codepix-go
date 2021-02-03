package model_test

import (
	"testing"

	"github.com/leoluzh/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewTransaction(t *testing.T) {

	code := "013"
	name := "Caixa Economica Federal"
	bank, _ := model.NewBank(code, name)

	accountNumber := "#-abc123-#"
	ownerName := "leoluzh"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestionation"
	ownerName = "Jonh Wicket"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "leoluzh@foobar.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 3.10
	statusTransaction := model.TransactionPeding
	transaction, err := model.NewTransaction(account, amount, pixKey, "Custom description")

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "Custom description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, err := model.NewPixKey(kind, account, key)

	_, err = model.NewTransaction(account, amount, pixKeySameAccount, "Custom description")
	require.Nil(t, err)

	_, err = model.NewTransaction(account, 0, pixKey, "Custom description")
	require.Nil(t, err)

}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {

	code := "013"
	name := "Caixa Economica Federal"
	bank, _ := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "leoluzh"
	account, _ := model.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "abcdestination"
	ownerName = "Jonh Wicket"
	accountDestination, _ := model.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "leoluzh@foobar.com"
	pixKey, _ := model.NewPixKey(kind, accountDestination, key)

	amount := 3.10
	transaction, _ := model.NewTransaction(account, amount, pixKey, "Custom description")

	transaction.Complete()
	require.Equal(t, transaction.Status, model.TransactionCompleted)

	transaction.Cancel("Testing Error")
	require.Equal(t, transaction.Status, model.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Testing Error")

}
