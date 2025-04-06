package handler

import (
	"context"
	"student-service/model"
	pb "student-service/proto"
	"student-service/repository"
)

type StudentHandler struct {
	pb.UnimplementedStudentServiceServer
	Repo *repository.StudentRepository
}

func (h *StudentHandler) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.StudentResponse, error) {
	student := &model.Student{
		Name:   req.Name,
		Email:  req.Email,
		Course: req.Course,
	}
	err := h.Repo.Create(student)
	if err != nil {
		return nil, err
	}
	return &pb.StudentResponse{Student: convertToProto(student)}, nil
}

func (h *StudentHandler) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.StudentResponse, error) {
	student, err := h.Repo.Get(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.StudentResponse{Student: convertToProto(student)}, nil
}

func (h *StudentHandler) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.StudentResponse, error) {
	student := &model.Student{
		ID:     req.Id,
		Name:   req.Name,
		Email:  req.Email,
		Course: req.Course,
	}
	err := h.Repo.Update(student)
	if err != nil {
		return nil, err
	}
	return &pb.StudentResponse{Student: convertToProto(student)}, nil
}

func (h *StudentHandler) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	err := h.Repo.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentResponse{Message: "Student deleted successfully"}, nil
}

func (h *StudentHandler) ListStudents(ctx context.Context, req *pb.ListStudentsRequest) (*pb.ListStudentsResponse, error) {
	students, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var protoStudents []*pb.Student
	for _, s := range students {
		protoStudents = append(protoStudents, convertToProto(s))
	}
	return &pb.ListStudentsResponse{Students: protoStudents}, nil
}

func convertToProto(s *model.Student) *pb.Student {
	return &pb.Student{
		Id:     s.ID,
		Name:   s.Name,
		Email:  s.Email,
		Course: s.Course,
	}
}
