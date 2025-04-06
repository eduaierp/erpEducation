package handler

import (
	"context"
	"reporting-service/model"
	pb "reporting-service/proto"
	"reporting-service/repository"
)

type ReportHandler struct {
	pb.UnimplementedReportingServiceServer
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

func (h *ReportHandler) GetReport(ctx context.Context, req *pb.GetReportRequest) (*pb.ReportResponse, error) {
	report, err := h.Repo.GetByID(req.Id)
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

func (h *ReportHandler) ListReports(ctx context.Context, req *pb.ListReportsRequest) (*pb.ListReportsResponse, error) {
	reports, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var pbReports []*pb.Report
	for _, r := range reports {
		pbReports = append(pbReports, &pb.Report{
			Id:        r.ID,
			Title:     r.Title,
			Content:   r.Content,
			CreatedAt: r.CreatedAt,
		})
	}

	return &pb.ListReportsResponse{Reports: pbReports}, nil
}
