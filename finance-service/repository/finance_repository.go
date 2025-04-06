package repository

import (
	"database/sql"
	"errors"
	"finance-service/model"
)

type FinanceRepository struct {
	DB *sql.DB
}

func (r *FinanceRepository) Create(tx *model.Transaction) error {
	return r.DB.QueryRow(
		"INSERT INTO transactions (type, amount, date, description) VALUES ($1, $2, $3, $4) RETURNING id",
		tx.Type, tx.Amount, tx.Date, tx.Description,
	).Scan(&tx.ID)
}

func (r *FinanceRepository) GetByID(id int32) (*model.Transaction, error) {
	row := r.DB.QueryRow("SELECT id, type, amount, date, description FROM transactions WHERE id = $1", id)

	tx := &model.Transaction{}
	err := row.Scan(&tx.ID, &tx.Type, &tx.Amount, &tx.Date, &tx.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}
	return tx, nil
}

func (r *FinanceRepository) List() ([]*model.Transaction, error) {
	rows, err := r.DB.Query("SELECT id, type, amount, date, description FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*model.Transaction
	for rows.Next() {
		tx := &model.Transaction{}
		if err := rows.Scan(&tx.ID, &tx.Type, &tx.Amount, &tx.Date, &tx.Description); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
