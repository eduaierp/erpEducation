package handler

import (
    "context"
    pb "student-service/proto"
    "student-service/model"
    "student-service/repository"
)

type StudentHandler struct {
    Repo *repository.StudentRepository
}

func (h *StudentHandler) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.StudentResponse, error) {
    student := &model.Student{
        Name: req.Name,
        Email: req.Email,
        Course: req.Course,
    }
    err := h.Repo.Create(student)
    if err != nil {
        return nil, err
    }
    return &pb.StudentResponse{Student: &pb.Student{
        Name: req.Name,
        Email: req.Email,
        Course: req.Course,
    }}, nil
}
