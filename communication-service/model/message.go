package model

type Message struct {
	ID      int32
	To      string
	Subject string
	Body    string
	SentAt  string
}
