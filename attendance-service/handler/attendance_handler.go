package handler

import (
	"attendance-service/model"
	pb "attendance-service/proto"
	"attendance-service/repository"
	"context"
)

type AttendanceHandler struct {
	Repo *repository.AttendanceRepository
	pb.UnimplementedAttendanceServiceServer
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

func (h *AttendanceHandler) GetAttendance(ctx context.Context, req *pb.GetAttendanceRequest) (*pb.AttendanceResponse, error) {
	a, err := h.Repo.GetByID(req.Id)
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

func (h *AttendanceHandler) ListAttendance(ctx context.Context, req *pb.ListAttendanceRequest) (*pb.ListAttendanceResponse, error) {
	records, err := h.Repo.ListByStudentID(req.StudentId)
	if err != nil {
		return nil, err
	}

	var result []*pb.Attendance
	for _, a := range records {
		result = append(result, &pb.Attendance{
			Id:        a.ID,
			StudentId: a.StudentID,
			Date:      a.Date,
			Status:    a.Status,
		})
	}

	return &pb.ListAttendanceResponse{
		Records: result,
	}, nil
}
