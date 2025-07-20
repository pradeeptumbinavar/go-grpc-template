// internal/handler/user.go
package handler

import (
    "context"
    "go-grpc-template/internal/port"
    "go-grpc-template/proto"
)

type UserHandler struct {
    port port.UserPort
    proto.UnimplementedUserServiceServer
}

func NewUserHandler(p port.UserPort) *UserHandler {
    return &UserHandler{port: p}
}

func (h *UserHandler) ListUsers(ctx context.Context, _ *proto.Empty) (*proto.UserList, error) {
    users, err := h.port.List()
    if err != nil {
        return nil, err
    }
    return &proto.UserList{Users: users}, nil
}
