package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{db: db, Queries: New(db)}
}

func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer
	FromAccount Account `json:"from_account"`
	ToAccount   Account `json:"to_account"`
	FromEntry   Entry   `json:"from_entry"`
	ToEntry     Entry   `json:"to_entry"`
}

func (s *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(ctx, func(queries *Queries) error {
		var err error

		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return nil
		}

		result.FromEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return nil
		}

		result.ToEntry, err = queries.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return nil
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, queries, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, queries, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return nil
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1, amount1, accountID2, amount2 int64) (account1 Account, account2 Account, err error) {
	err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount1,
		ID:     accountID1,
	})
	if err != nil {
		return
	}
	account1, err = q.GetAccount(ctx, accountID1)
	if err != nil {
		return
	}

	err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		Amount: amount2,
		ID:     accountID2,
	})
	if err != nil {
		return
	}
	account2, err = q.GetAccount(ctx, accountID2)
	if err != nil {
		return
	}

	return
}
