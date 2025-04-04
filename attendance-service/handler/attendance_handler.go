package handler

import (
	"attendance-service/model"
	pb "attendance-service/proto"
	"attendance-service/repository"
	"context"
)

type AttendanceHandler struct {
	Repo *repository.AttendanceRepository
}

func (h *AttendanceHandler) MarkAttendance(ctx context.Context, req *pb.MarkAttendanceRequest) (*pb.AttendanceResponse, error) {
	a := &model.Attendance{
		StudentID: req.StudentId,
		Date:      req.Date,
		Status:    req.Status,
	}

	err := h.Repo.Mark(a)
	if err != nil {
		return nil, err
	}

	return &pb.AttendanceResponse{
		Attendance: &pb.Attendance{
			Id:        a.ID,
			StudentId: a.StudentID,
			Date:      a.Date,
			Status:    a.Status,
		},
	}, nil
}
