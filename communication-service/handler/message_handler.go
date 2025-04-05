package handler

import (
	"communication-service/model"
	pb "communication-service/proto"
	"communication-service/repository"
	"context"
)

type MessageHandler struct {
	pb.UnimplementedCommunicationServiceServer
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
		Message: toProto(msg),
	}, nil
}

func (h *MessageHandler) GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.MessageResponse, error) {
	msg, err := h.Repo.Get(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.MessageResponse{
		Message: toProto(msg),
	}, nil
}

func (h *MessageHandler) ListMessages(ctx context.Context, req *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	msgs, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var pbMsgs []*pb.Message
	for _, msg := range msgs {
		pbMsgs = append(pbMsgs, toProto(msg))
	}

	return &pb.ListMessagesResponse{
		Messages: pbMsgs,
	}, nil
}

func toProto(msg *model.Message) *pb.Message {
	return &pb.Message{
		Id:      msg.ID,
		To:      msg.To,
		Subject: msg.Subject,
		Body:    msg.Body,
		SentAt:  msg.SentAt,
	}
}
