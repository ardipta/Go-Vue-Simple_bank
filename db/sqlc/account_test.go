package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ashiqur/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account11 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account11.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account11.ID, account2.ID)
	require.Equal(t, account11.Owner, account2.Owner)
	require.Equal(t, account11.Balance, account2.Balance)
	require.Equal(t, account11.Currency, account2.Currency)
	require.WithinDuration(t, account11.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
	account11 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account11.ID,
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account11.ID, account2.ID)
	require.Equal(t, account11.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, arg.Currency, account2.Currency)
	require.WithinDuration(t, account11.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T){
	account11 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account11.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account11.Balance)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts{
		require.NotEmpty(t, account)
	}
}