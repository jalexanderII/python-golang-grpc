package auth

import (
	"context"
	"log"

	"auth/usecase"

	"github.com/hashicorp/go-hclog"
)

type Server struct {
	userUsecase usecase.UserUsecase
	l           hclog.Logger
	UnimplementedAuthServiceServer
}

func InitServer(userUsecase usecase.UserUsecase, l hclog.Logger) Server {
	return Server{
		userUsecase,
		l,
		UnimplementedAuthServiceServer{},
	}
}

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	s.l.Info("Register called")
	_, err := s.userUsecase.Register(request.Username, request.Password)
	if err != nil {
		log.Println("Error to register user", err)
		return &RegisterResponse{Success: false, Message: "Failed to register"}, err
	}
	return &RegisterResponse{Success: true, Message: "Succeed to register"}, nil
}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	s.l.Info("Login called")
	token, err := s.userUsecase.Login(request.Username, request.Password)
	if err != nil {
		log.Println("Error to register user", err)
		return &LoginResponse{Success: false, Message: "Failed to login", Token: ""}, err
	}
	return &LoginResponse{Success: true, Message: "Login success", Token: token}, nil
}

func (s *Server) ValidateToken(ctx context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	result, err := s.userUsecase.ValidateToken(request.Token)
	if err != nil {
		log.Println("Error to validate token", err)
		return &ValidateTokenResponse{Success: false, Message: "Invalid token"}, err
	}
	s.l.Info("Token valid")
	return &ValidateTokenResponse{Success: true, Message: result}, nil
}
