package handler

import (
	"context"
	"finance-service/model"
	pb "finance-service/proto"
	"finance-service/repository"
)

type FinanceHandler struct {
	Repo *repository.FinanceRepository
}

func (h *FinanceHandler) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	tx := &model.Transaction{
		Type:        req.Type,
		Amount:      req.Amount,
		Date:        req.Date,
		Description: req.Description,
	}

	err := h.Repo.Create(tx)
	if err != nil {
		return nil, err
	}

	return &pb.TransactionResponse{
		Transaction: &pb.Transaction{
			Id:          tx.ID,
			Type:        tx.Type,
			Amount:      tx.Amount,
			Date:        tx.Date,
			Description: tx.Description,
		},
	}, nil
}
