package service

import (
	"context"
	"scratch/internal/database"
	"time"
)

type UserService struct {
	Queries *database.Queries
}

func (s *UserService) CreateUser(ctx context.Context, name string) (*database.User, error) {
	now := time.Now().UTC()
	user, err := s.Queries.CreateUser(ctx, database.CreateUserParams{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
