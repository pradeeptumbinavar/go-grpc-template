// internal/port/user.go
package port

import "go-grpc-template/proto"

type UserPort interface {
    List() ([]*proto.User, error)
}
