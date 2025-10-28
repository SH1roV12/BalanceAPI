package service

import (
	"context"
	"users-balance/internal/errmsg"
	"users-balance/internal/handlers/dto"
	"users-balance/internal/models"
	"users-balance/internal/repository"
)


type UsersService interface{
	CreateUser(ctx context.Context, user *models.User) error 
	GetUserBalance(ctx context.Context, usr dto.GetUserBalance) (float64, float64, error)
	ReserveUserBalance(ctx context.Context, usr dto.ReserveUserBalance) (float64, float64, error)
	ReplenishmentOfBalance(ctx context.Context, usr dto.ReplenishmentOfBalance) (float64, error)
}

type Service struct {
	Repo repository.UserRepository
}


//func for DI
func NewService(repo repository.UserRepository) *Service {
	return &Service{Repo: repo}
}

//Create new user method
func (s *Service) CreateUser(ctx context.Context, user *models.User) error {
	return s.Repo.CreateUser(ctx, user)
}

//Method to getting user balance by id
func (s *Service) GetUserBalance(ctx context.Context, usr dto.GetUserBalance) (float64, float64, error) {
	return s.Repo.GetUserBalance(ctx, usr.UserId)
}

//Method for reserve user balance
func (s *Service) ReserveUserBalance(ctx context.Context, usr dto.ReserveUserBalance) (float64, float64, error) {
	balance, reserved, err := s.Repo.ReserveUserBalance(ctx, usr.UserId, usr.Amount)


	if err != nil {
		return 0, 0, err
	}


	return balance, reserved, err
}


//Method to replenish user balance
func (s *Service) ReplenishmentOfBalance(ctx context.Context, usr dto.ReplenishmentOfBalance) (float64, error) {
	
	
	if usr.Amount <= 0 {
		return 0, errmsg.ErrIncorrectAmount
	}


	balance, error := s.Repo.ReplenishmentOfBalance(ctx, usr.UserID, usr.Amount)


	if error != nil {
		return 0, error
	}


	return balance, nil
}
