// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transaction.sql

package repository

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transaction(email_id,status)
VALUES ($1,$2)
RETURNING id, email_id, status, created_at
`

type CreateTransactionParams struct {
	EmailID int64  `json:"email_id"`
	Status  string `json:"status"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction, arg.EmailID, arg.Status)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.EmailID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, email_id, status, created_at FROM transaction
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.EmailID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listTransactions = `-- name: ListTransactions :many
SELECT id, email_id, status, created_at FROM transaction
        WHERE email_id = $1
        ORDER BY id
        LIMIT $2
        OFFSET $3
`

type ListTransactionsParams struct {
	EmailID int64 `json:"email_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListTransactions(ctx context.Context, arg ListTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactions, arg.EmailID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.EmailID,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}