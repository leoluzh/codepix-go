package model_test

import (
	"testing"

	"github.com/leoluzh/codepix-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {

	code := "013"
	name := "Caixa Economica Federal"

	bank, err := model.NewBank(code, name)

	accountNumber := "#--abc123--#"
	ownerName := "leoluzh"

	account, err := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "leoluz@foobar.com"

	pixKey, err := model.NewPixKey(kind, account, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"

	_, err = model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.Nil(t, err)

}
