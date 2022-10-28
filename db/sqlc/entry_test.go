package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateEntry(t *testing.T) {
	entry, account, err := createRandomEntry(t)
	require.NoError(t, err)

	require.NotEmpty(t, account)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, entry.Amount, account.Balance)
}

func createRandomEntry(t *testing.T) (Entry, Account, error) {
	account := createRandomAccount(t)
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    account.Balance,
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	return entry, account, err
}

func TestQueries_GetEntry(t *testing.T) {
	entry, account, err := createRandomEntry(t)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotEmpty(t, entry)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.Equal(t, entry.CreatedAt, entry2.CreatedAt)
}
