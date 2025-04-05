package handler

import (
	"context"
	"faculty-service/model"
	pb "faculty-service/proto"
	"faculty-service/repository"
)

type FacultyHandler struct {
	Repo *repository.FacultyRepository
	pb.UnimplementedFacultyServiceServer // 👈 Add this line
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

func (h *FacultyHandler) GetFaculty(ctx context.Context, req *pb.GetFacultyRequest) (*pb.FacultyResponse, error) {
	faculty, err := h.Repo.GetByID(req.Id)
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

func (h *FacultyHandler) UpdateFaculty(ctx context.Context, req *pb.UpdateFacultyRequest) (*pb.FacultyResponse, error) {
	faculty := &model.Faculty{
		ID:          req.Id,
		Name:        req.Name,
		Email:       req.Email,
		Department:  req.Department,
		Designation: req.Designation,
	}

	err := h.Repo.Update(faculty)
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

func (h *FacultyHandler) DeleteFaculty(ctx context.Context, req *pb.DeleteFacultyRequest) (*pb.DeleteFacultyResponse, error) {
	err := h.Repo.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteFacultyResponse{Message: "Faculty deleted successfully"}, nil
}

func (h *FacultyHandler) ListFaculties(ctx context.Context, req *pb.ListFacultiesRequest) (*pb.ListFacultiesResponse, error) {
	facultyList, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var res []*pb.Faculty
	for _, f := range facultyList {
		res = append(res, &pb.Faculty{
			Id:          f.ID,
			Name:        f.Name,
			Email:       f.Email,
			Department:  f.Department,
			Designation: f.Designation,
		})
	}

	return &pb.ListFacultiesResponse{
		Faculties: res,
	}, nil
}
