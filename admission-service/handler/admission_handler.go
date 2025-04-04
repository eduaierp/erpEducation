package handler

import (
	"admission-service/model"
	pb "admission-service/proto"
	"admission-service/repository"
	"context"
)

type AdmissionHandler struct {
	Repo *repository.AdmissionRepository
}

func (h *AdmissionHandler) ApplyAdmission(ctx context.Context, req *pb.ApplyAdmissionRequest) (*pb.AdmissionResponse, error) {
	admission := &model.Admission{
		StudentName: req.StudentName,
		Course:      req.Course,
		AppliedDate: req.AppliedDate,
	}

	err := h.Repo.Apply(admission)
	if err != nil {
		return nil, err
	}

	return &pb.AdmissionResponse{
		Admission: &pb.Admission{
			Id:          admission.ID,
			StudentName: admission.StudentName,
			Course:      admission.Course,
			Status:      admission.Status,
			AppliedDate: admission.AppliedDate,
		},
	}, nil
}
