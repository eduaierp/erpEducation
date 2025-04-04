package repository

import (
	"communication-service/model"
	"database/sql"
)

type MessageRepository struct {
	DB *sql.DB
}

func (r *MessageRepository) Send(msg *model.Message) error {
	return r.DB.QueryRow(
		"INSERT INTO messages (recipient, subject, body, sent_at) VALUES ($1, $2, $3, $4) RETURNING id",
		msg.To, msg.Subject, msg.Body, msg.SentAt,
	).Scan(&msg.ID)
}
