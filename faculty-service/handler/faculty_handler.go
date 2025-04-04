package handler

import (
	"context"
	"faculty-service/model"
	pb "faculty-service/proto"
	"faculty-service/repository"
)

type FacultyHandler struct {
	Repo *repository.FacultyRepository
}

func (h *FacultyHandler) CreateFaculty(ctx context.Context, req *pb.CreateFacultyRequest) (*pb.FacultyResponse, error) {
	faculty := &model.Faculty{
		Name:        req.Name,
		Email:       req.Email,
		Department:  req.Department,
		Designation: req.Designation,
	}

	err := h.Repo.Create(faculty)
	if err != nil {
		return nil, err
	}

	return &pb.FacultyResponse{
		Faculty: &pb.Faculty{
			Id:          faculty.ID,
			Name:        faculty.Name,
			Email:       faculty.Email,
			Department:  faculty.Department,
			Designation: faculty.Designation,
		},
	}, nil
}
