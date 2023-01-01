package service

import (
	"context"
	"github.com/f-robo/pb-api-gateway/pkg/auth/pb"
	"net/http"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	//var user model.User

	if req.Email != "gigasu" {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token := "token12345"

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
