package handler

import (
	"context"
	"hr-service/model"
	pb "hr-service/proto"
	"hr-service/repository"
)

type HRHandler struct {
	pb.UnimplementedHRServiceServer
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

func (h *HRHandler) GetEmployee(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.EmployeeResponse, error) {
	emp, err := h.Repo.Get(req.Id)
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

func (h *HRHandler) ListEmployees(ctx context.Context, req *pb.ListEmployeesRequest) (*pb.ListEmployeesResponse, error) {
	emps, err := h.Repo.List()
	if err != nil {
		return nil, err
	}

	var protoEmps []*pb.Employee
	for _, emp := range emps {
		protoEmps = append(protoEmps, &pb.Employee{
			Id:         emp.ID,
			Name:       emp.Name,
			Department: emp.Department,
			Role:       emp.Role,
			Doj:        emp.DOJ,
		})
	}

	return &pb.ListEmployeesResponse{
		Employees: protoEmps,
	}, nil
}
