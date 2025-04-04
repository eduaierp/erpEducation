package handler

import (
	pb "auth-service/proto"
	"auth-service/repository"
	"auth-service/utils"
	"context"
	"errors"
)

type AuthHandler struct {
	Repo *repository.UserRepository
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.Repo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if req.Password != user.Password {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
		Role:  user.Role,
	}, nil
}
