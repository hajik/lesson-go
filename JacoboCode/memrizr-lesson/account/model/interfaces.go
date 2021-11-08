package model

import (
	"context"

	"github.com/google/uuid"
)

// UserService define method the handler layer expects
// any service it interacts with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

// UserRepoitory define method the service layer expects
// any respository it interacts with to implement
type UserRespository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}
