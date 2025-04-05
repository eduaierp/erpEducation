package handler

import (
	"admission-service/model"
	pb "admission-service/proto"
	"admission-service/repository"
	"context"
)

type AdmissionHandler struct {
	pb.UnimplementedAdmissionServiceServer
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

func (h *AdmissionHandler) GetAdmission(ctx context.Context, req *pb.GetAdmissionRequest) (*pb.AdmissionResponse, error) {
	admission, err := h.Repo.GetByID(req.Id)
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

func (h *AdmissionHandler) ListAdmissions(ctx context.Context, req *pb.ListAdmissionsRequest) (*pb.ListAdmissionsResponse, error) {
	list, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var admissions []*pb.Admission
	for _, a := range list {
		admissions = append(admissions, &pb.Admission{
			Id:          a.ID,
			StudentName: a.StudentName,
			Course:      a.Course,
			Status:      a.Status,
			AppliedDate: a.AppliedDate,
		})
	}

	return &pb.ListAdmissionsResponse{Admissions: admissions}, nil
}
