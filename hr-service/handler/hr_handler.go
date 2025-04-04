package handler

import (
	"context"
	"hr-service/model"
	pb "hr-service/proto"
	"hr-service/repository"
)

type HRHandler struct {
	Repo *repository.HRRepository
}

func (h *HRHandler) AddEmployee(ctx context.Context, req *pb.AddEmployeeRequest) (*pb.EmployeeResponse, error) {
	emp := &model.Employee{
		Name:       req.Name,
		Department: req.Department,
		Role:       req.Role,
		DOJ:        req.Doj,
	}

	err := h.Repo.Add(emp)
	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{
		Employee: &pb.Employee{
			Id:         emp.ID,
			Name:       emp.Name,
			Department: emp.Department,
			Role:       emp.Role,
			Doj:        emp.DOJ,
		},
	}, nil
}
