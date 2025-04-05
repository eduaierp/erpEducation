package handler

import (
	"context"
	"exam-service/model"
	pb "exam-service/proto"
	"exam-service/repository"
)

type ExamHandler struct {
	pb.UnimplementedExamServiceServer
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

func (h *ExamHandler) GetExam(ctx context.Context, req *pb.GetExamRequest) (*pb.ExamResponse, error) {
	exam, err := h.Repo.GetByID(req.Id)
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

func (h *ExamHandler) ListExams(ctx context.Context, req *pb.ListExamsRequest) (*pb.ListExamsResponse, error) {
	exams, err := h.Repo.ListAll()
	if err != nil {
		return nil, err
	}

	var pbExams []*pb.Exam
	for _, exam := range exams {
		pbExams = append(pbExams, &pb.Exam{
			Id:      exam.ID,
			Subject: exam.Subject,
			Date:    exam.Date,
			Type:    exam.Type,
		})
	}

	return &pb.ListExamsResponse{Exams: pbExams}, nil
}
