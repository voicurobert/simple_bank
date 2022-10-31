package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/voicurobert/simple_bank/db/sqlc"
	"github.com/voicurobert/simple_bank/util"
	"testing"
	"time"
)

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) db.Account {
	args := db.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestQueries_GetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)

	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)

}

func TestQueries_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := db.UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)

	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.Equal(t, account.CreatedAt, account2.CreatedAt)

}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := db.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
