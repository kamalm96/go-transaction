package db

import (
	"context"
	"testing"

	"example.com/m/v2/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	acc1 := createRandomAccount(t)
	require.NotEmpty(t, acc1)

	arg := CreateEntryParams{
		AccountID: acc1.ID,
		Amount:    util.RandomMoney(),
	}

	entryCreated, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entryCreated)
	require.Equal(t, arg.AccountID, acc1.ID)
	require.True(t, arg.Amount > 0)

	return entryCreated
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {

	entry := createRandomEntry(t)

	getEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntry)
	require.Equal(t, entry.AccountID, getEntry.AccountID)
	require.Equal(t, entry.Amount, getEntry.Amount)
	require.Equal(t, entry.CreatedAt, getEntry.CreatedAt)
}
