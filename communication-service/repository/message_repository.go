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

func (r *MessageRepository) Get(id int32) (*model.Message, error) {
	row := r.DB.QueryRow("SELECT id, recipient, subject, body, sent_at FROM messages WHERE id = $1", id)
	msg := &model.Message{}
	err := row.Scan(&msg.ID, &msg.To, &msg.Subject, &msg.Body, &msg.SentAt)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (r *MessageRepository) List() ([]*model.Message, error) {
	rows, err := r.DB.Query("SELECT id, recipient, subject, body, sent_at FROM messages ORDER BY sent_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		msg := &model.Message{}
		err := rows.Scan(&msg.ID, &msg.To, &msg.Subject, &msg.Body, &msg.SentAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
