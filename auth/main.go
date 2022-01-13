package main

import (
	"fmt"
	"log"
	"net"

	"auth/auth"
	"auth/config"
	"auth/repository"
	"auth/usecase"

	"github.com/hashicorp/go-hclog"

	"google.golang.org/grpc"
)

func main() {
	l := hclog.Default()
	l.Debug("Core Service")

	db := config.ConnectDB()
	userRepository := repository.InitUserRepository(db)
	userUsecase := usecase.InitUserUsecase(userRepository)

	s := auth.InitServer(userUsecase, l)

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		l.Error("failed to listen", "error", err)
		panic(err)
	}

	l.Info("Server started", "port", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
	}
}
