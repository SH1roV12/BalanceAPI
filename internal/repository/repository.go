package repository

import (
	"context"
	"database/sql"
	"errors"
	"users-balance/internal/errmsg"
	"users-balance/internal/models"
)

type UserRepository interface{
	CreateUser(ctx context.Context, user *models.User)error
	GetUserBalance(ctx context.Context, userId int) (float64, float64, error)
	ReserveUserBalance(ctx context.Context, userId int, amount float64) (float64, float64, error) 
	ReplenishmentOfBalance(ctx context.Context, userId int, amount float64) (float64, error) 
}

type Repository struct {
	Database *sql.DB
}

//func for DI
func NewRepository(db *sql.DB) *Repository {
	return &Repository{Database: db}
}


//Create new user method
func (r *Repository) CreateUser(ctx context.Context, user *models.User) error {
	query:="INSERT INTO users (balance,reserved) VALUES(?,?)"

	_, err := r.Database.ExecContext(ctx, query, user.Balance, user.Reserved)
	if err!=nil{
		return err
	}

	return nil
}

//Method to getting user balance by id
func (r *Repository) GetUserBalance(ctx context.Context, userId int) (float64, float64, error) {
	var balance, reserved float64

	query := "SELECT balance, reserved FROM users WHERE id = ?"

	err := r.Database.QueryRowContext(ctx, query, userId).Scan(&balance, &reserved)


	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, errmsg.ErrUserNotFound
		}
		return 0, 0, err
	}


	return balance, reserved, nil
}


//Method for reserve user balance
func (r *Repository) ReserveUserBalance(ctx context.Context, userId int, amount float64) (float64, float64, error) {
	tx, err := r.Database.BeginTx(ctx, nil)
	if err != nil {
		return 0, 0, err
	}

	defer tx.Rollback()

	var balanceNow, reservedNow float64

	query := "SELECT balance,reserved FROM users WHERE id = ? FOR UPDATE"

	err = tx.QueryRowContext(ctx, query, userId).Scan(&balanceNow, &reservedNow)


	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, 0, errmsg.ErrUserNotFound
		}
		return 0, 0, err
	}


	if balanceNow < amount {
		return 0, 0, errmsg.ErrNotEnoughMoney
	}


	changeBalanceQuery := "UPDATE users SET balance = ?,reserved = ? WHERE id = ?"

	_, err = tx.ExecContext(ctx, changeBalanceQuery, balanceNow-amount, reservedNow+amount, userId)


	if err != nil {
		return 0, 0, err
	}


	err = tx.Commit()
	if err != nil {
		return 0, 0, err
	}


	return reservedNow + amount, balanceNow - amount, nil
}

//Method to replenish user balance
func (r *Repository) ReplenishmentOfBalance(ctx context.Context, userId int, amount float64) (float64, error) {
	tx, err := r.Database.BeginTx(ctx, nil)


	if err != nil {
		tx.Rollback()
		return 0, err
	}


	
	queryString := "UPDATE users SET balance = balance+? WHERE id = ?"
	_, err = tx.ExecContext(ctx, queryString, amount, userId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}


	err = tx.Commit()


	if err != nil {
		tx.Rollback()
		return 0, err
	}


	balanceAfter, _, err := r.GetUserBalance(ctx, userId)


	if err != nil {
		tx.Rollback()
		return 0, err
	}


	return balanceAfter, nil
}
