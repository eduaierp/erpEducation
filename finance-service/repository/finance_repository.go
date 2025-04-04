package repository

import (
    "database/sql"
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
