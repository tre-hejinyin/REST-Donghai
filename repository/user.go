package repository

import (
	"context"

	"server/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (domain.Users, error)
	Create(ctx context.Context, user *domain.Users) error
	UpdateByEmail(ctx context.Context, email string, user *domain.Users) (*domain.Users, error)
}
