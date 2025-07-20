// internal/service/user.go
package service

import (
    "go-grpc-template/internal/port"
    "go-grpc-template/proto"
)

// UserServiceImpl is our “business logic” adapter.
type UserServiceImpl struct{}

// Ensure it satisfies the port interface:
var _ port.UserPort = (*UserServiceImpl)(nil)

func NewUserService() *UserServiceImpl {
    return &UserServiceImpl{}
}

// List returns a hard‑coded slice for now.
func (s *UserServiceImpl) List() ([]*proto.User, error) {
    return []*proto.User{
        {Id: 1, Name: "Alice"},
        {Id: 2, Name: "Bob"},
    }, nil
}
