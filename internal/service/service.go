package service

import (
	"context"
	"errors"
	"users-balance/internal/handlers/dto"
	"users-balance/internal/models"
	"users-balance/internal/repository"
)

type UsersService struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *UsersService {
	return &UsersService{Repo: repo}
}

func (s *UsersService) CreateUser(ctx context.Context, user *models.User) error {
	return s.Repo.CreateUser(ctx, user)
}

func (s *UsersService) GetUserBalance(ctx context.Context, usr dto.GetUserBalance) (float64, float64, error) {
	return s.Repo.GetUserBalance(ctx, usr.UserId)
}

func (s *UsersService) ReserveUserBalance(ctx context.Context, usr dto.ReserveUserBalance) (float64, float64, error) {
	balance, reserved, err := s.Repo.ReserveUserBalance(ctx, usr.UserId, usr.Amount)
	if err != nil {
		return 0, 0, err
	}
	return balance, reserved, err
}

func (s *UsersService) ReplenishmentOfBalance(ctx context.Context, usr dto.ReplenishmentOfBalance) (float64, error) {
	if usr.Amount <= 0 {
		return 0, errors.New("incorrect amount")
	}
	balance, error := s.Repo.ReplenishmentOfBalance(ctx, usr.UserID, usr.Amount)
	if error != nil {
		return 0, error
	}
	return balance, nil
}
