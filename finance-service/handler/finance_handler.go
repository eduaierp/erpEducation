package handler

import (
	"context"
	"finance-service/model"
	pb "finance-service/proto"
	"finance-service/repository"
)

type FinanceHandler struct {
	pb.UnimplementedFinanceServiceServer
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

func (h *FinanceHandler) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.TransactionResponse, error) {
	tx, err := h.Repo.GetByID(req.Id)
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

func (h *FinanceHandler) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	transactions, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var pbTransactions []*pb.Transaction
	for _, tx := range transactions {
		pbTransactions = append(pbTransactions, &pb.Transaction{
			Id:          tx.ID,
			Type:        tx.Type,
			Amount:      tx.Amount,
			Date:        tx.Date,
			Description: tx.Description,
		})
	}

	return &pb.ListTransactionsResponse{
		Transactions: pbTransactions,
	}, nil
}
