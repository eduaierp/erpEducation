package handler

import (
    "context"
    pb "academic-service/proto"
    "academic-service/model"
    "academic-service/repository"
)

type AcademicHandler struct {
    Repo *repository.AcademicRepository
}

func (h *AcademicHandler) CreateProgram(ctx context.Context, req *pb.CreateProgramRequest) (*pb.ProgramResponse, error) {
    program := &model.Program{
        Name: req.Name,
        Description: req.Description,
    }
    err := h.Repo.CreateProgram(program)
    if err != nil {
        return nil, err
    }
    return &pb.ProgramResponse{Program: &pb.Program{
        Name: req.Name,
        Description: req.Description,
    }}, nil
}
