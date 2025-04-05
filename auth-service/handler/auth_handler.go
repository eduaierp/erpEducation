package handler

import (
	"auth-service/proto"
	"auth-service/repository"
	"auth-service/utils"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Repo *repository.UserRepository
	proto.UnimplementedAuthServiceServer
}

func (h *AuthHandler) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := h.Repo.GetByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Token: token,
		Role:  user.Role,
	}, nil
}
