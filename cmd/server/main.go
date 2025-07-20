// cmd/server/main.go
package main

import (
    "fmt"
    "log"
    "net"
    "go-grpc-template/config"
    "go-grpc-template/internal/handler"
    "go-grpc-template/internal/service"
	"google.golang.org/grpc/reflection"

    "go-grpc-template/proto"
    "google.golang.org/grpc"
)

func main() {
    cfg := config.Load()
    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ServerPort))
    if err != nil {
        log.Fatalf("listen failed: %v", err)
    }

    grpcServer := grpc.NewServer()
    userSvc := service.NewUserService()
    proto.RegisterUserServiceServer(grpcServer, handler.NewUserHandler(userSvc))
	reflection.Register(grpcServer)

    log.Printf("gRPC server listening on %s", cfg.ServerPort)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("serve failed: %v", err)
    }
}
