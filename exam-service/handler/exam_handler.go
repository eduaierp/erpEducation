package handler

import (
	"context"
	"exam-service/model"
	pb "exam-service/proto"
	"exam-service/repository"
)

type ExamHandler struct {
	Repo *repository.ExamRepository
}

func (h *ExamHandler) CreateExam(ctx context.Context, req *pb.CreateExamRequest) (*pb.ExamResponse, error) {
	exam := &model.Exam{
		Subject: req.Subject,
		Date:    req.Date,
		Type:    req.Type,
	}

	err := h.Repo.Create(exam)
	if err != nil {
		return nil, err
	}

	return &pb.ExamResponse{
		Exam: &pb.Exam{
			Id:      exam.ID,
			Subject: exam.Subject,
			Date:    exam.Date,
			Type:    exam.Type,
		},
	}, nil
}
