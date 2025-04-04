package handler

import (
	"context"
	"reporting-service/model"
	pb "reporting-service/proto"
	"reporting-service/repository"
)

type ReportHandler struct {
	Repo *repository.ReportRepository
}

func (h *ReportHandler) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.ReportResponse, error) {
	report := &model.Report{
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: req.CreatedAt,
	}

	err := h.Repo.Create(report)
	if err != nil {
		return nil, err
	}

	return &pb.ReportResponse{
		Report: &pb.Report{
			Id:        report.ID,
			Title:     report.Title,
			Content:   report.Content,
			CreatedAt: report.CreatedAt,
		},
	}, nil
}
