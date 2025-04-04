package handler

import (
    "context"
    "communication-service/model"
    "communication-service/repository"
    pb "communication-service/proto"
)

type MessageHandler struct {
    Repo *repository.MessageRepository
}

func (h *MessageHandler) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.MessageResponse, error) {
    msg := &model.Message{
        To:      req.To,
        Subject: req.Subject,
        Body:    req.Body,
        SentAt:  req.SentAt,
    }

    err := h.Repo.Send(msg)
    if err != nil {
        return nil, err
    }

    return &pb.MessageResponse{
        Message: &pb.Message{
            Id:      msg.ID,
            To:      msg.To,
            Subject: msg.Subject,
            Body:    msg.Body,
            SentAt:  msg.SentAt,
        },
    }, nil
}
