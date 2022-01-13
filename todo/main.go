package main

import (
	"fmt"
	"log"
	"net"

	"todo/config"
	"todo/repository"
	"todo/todo"
	"todo/usecase"

	"github.com/hashicorp/go-hclog"

	"google.golang.org/grpc"
)

func main() {
	l := hclog.Default()
	l.Debug("Todo Service")
	db := config.ConnectDB()
	TodoRepository := repository.InitTodoRepository(db)
	todoUsecase := usecase.InitUserUsecase(TodoRepository)

	s := todo.InitServer(todoUsecase, l)

	grpcServer := grpc.NewServer()

	todo.RegisterTodoServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		l.Error("failed to listen", "error", err)
		panic(err)
	}

	l.Info("Server started", "port", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %s\n", err)
	}
}
