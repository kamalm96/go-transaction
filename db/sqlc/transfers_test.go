package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        10,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, acc1.ID)
	require.Equal(t, transfer.ToAccountID, acc2.ID)
	require.Equal(t, transfer.Amount, arg.Amount)
	return transfer
}

func TestCreateTransfer(t *testing.T) {

	createRandomTransfer(t)
}
