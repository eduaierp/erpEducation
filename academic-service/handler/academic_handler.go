package handler

import (
	"academic-service/model"
	pb "academic-service/proto"
	"academic-service/repository"
	"context"
)

type AcademicHandler struct {
	pb.UnimplementedAcademicServiceServer
	Repo *repository.AcademicRepository
}

// Program Handlers

func (h *AcademicHandler) CreateProgram(ctx context.Context, req *pb.CreateProgramRequest) (*pb.ProgramResponse, error) {
	program := &model.Program{
		Name:        req.Name,
		Description: req.Description,
	}
	err := h.Repo.CreateProgram(program)
	if err != nil {
		return nil, err
	}
	return &pb.ProgramResponse{Program: &pb.Program{
		Id:          program.ID,
		Name:        program.Name,
		Description: program.Description,
	}}, nil
}

func (h *AcademicHandler) GetProgram(ctx context.Context, req *pb.GetProgramRequest) (*pb.ProgramResponse, error) {
	program, err := h.Repo.GetProgram(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ProgramResponse{Program: &pb.Program{
		Id:          program.ID,
		Name:        program.Name,
		Description: program.Description,
	}}, nil
}

func (h *AcademicHandler) UpdateProgram(ctx context.Context, req *pb.UpdateProgramRequest) (*pb.ProgramResponse, error) {
	program := &model.Program{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}
	err := h.Repo.UpdateProgram(program)
	if err != nil {
		return nil, err
	}
	return &pb.ProgramResponse{Program: &pb.Program{
		Id:          program.ID,
		Name:        program.Name,
		Description: program.Description,
	}}, nil
}

func (h *AcademicHandler) DeleteProgram(ctx context.Context, req *pb.DeleteProgramRequest) (*pb.DeleteProgramResponse, error) {
	err := h.Repo.DeleteProgram(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProgramResponse{Message: "Program deleted successfully"}, nil
}

func (h *AcademicHandler) ListPrograms(ctx context.Context, req *pb.ListProgramsRequest) (*pb.ListProgramsResponse, error) {
	programs, err := h.Repo.ListPrograms()
	if err != nil {
		return nil, err
	}
	var pbPrograms []*pb.Program
	for _, p := range programs {
		pbPrograms = append(pbPrograms, &pb.Program{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return &pb.ListProgramsResponse{Programs: pbPrograms}, nil
}

// Course Handlers

func (h *AcademicHandler) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CourseResponse, error) {
	course := &model.Course{
		Title:     req.Title,
		Code:      req.Code,
		ProgramID: req.ProgramId,
	}
	err := h.Repo.CreateCourse(course)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{Course: &pb.Course{
		Id:        course.ID,
		Title:     course.Title,
		Code:      course.Code,
		ProgramId: course.ProgramID,
	}}, nil
}

func (h *AcademicHandler) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.CourseResponse, error) {
	course, err := h.Repo.GetCourse(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{Course: &pb.Course{
		Id:        course.ID,
		Title:     course.Title,
		Code:      course.Code,
		ProgramId: course.ProgramID,
	}}, nil
}

func (h *AcademicHandler) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.CourseResponse, error) {
	course := &model.Course{
		ID:        req.Id,
		Title:     req.Title,
		Code:      req.Code,
		ProgramID: req.ProgramId,
	}
	err := h.Repo.UpdateCourse(course)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{Course: &pb.Course{
		Id:        course.ID,
		Title:     course.Title,
		Code:      course.Code,
		ProgramId: course.ProgramID,
	}}, nil
}

func (h *AcademicHandler) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	err := h.Repo.DeleteCourse(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCourseResponse{Message: "Course deleted successfully"}, nil
}

func (h *AcademicHandler) ListCourses(ctx context.Context, req *pb.ListCoursesRequest) (*pb.ListCoursesResponse, error) {
	courses, err := h.Repo.ListCourses()
	if err != nil {
		return nil, err
	}
	var pbCourses []*pb.Course
	for _, c := range courses {
		pbCourses = append(pbCourses, &pb.Course{
			Id:        c.ID,
			Title:     c.Title,
			Code:      c.Code,
			ProgramId: c.ProgramID,
		})
	}
	return &pb.ListCoursesResponse{Courses: pbCourses}, nil
}
