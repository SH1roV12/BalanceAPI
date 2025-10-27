package repository

import (
	"context"
	"database/sql"
	"errors"
	"users-balance/internal/models"
)

type Repository struct{
	Database *sql.DB
}

func NewRepository(db *sql.DB)*Repository{
	return &Repository{Database: db}
}

func (db *Repository) CreateUser(ctx context.Context, user *models.User)error{
	_,err:=db.Database.ExecContext(ctx,"INSERT INTO users (balance,reserved) VALUES(?,?)",user.Balance,user.Reserved)
	return err
}

func (db *Repository) GetUserBalance(ctx context.Context, userId int)(float64,float64, error){
	var balance,reserved float64
	query:="SELECT balance, reserved FROM users WHERE id = ?"
	err:=db.Database.QueryRowContext(ctx,query, userId).Scan(&balance,&reserved)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
			return 0,0,errors.New("user not found")
		}
		return 0,0,err
	}
	return balance,reserved,nil
}

func (db *Repository) ReserveUserBalance(ctx context.Context, userId int, amount float64)(float64, float64,error){
	tx,err:=db.Database.BeginTx(ctx,nil)
	if err !=nil{
		return 0,0,err
	}


	defer tx.Rollback()

	var balanceNow,reservedNow float64
	checkBalanceQuery:="SELECT balance,reserved FROM users WHERE id = ? FOR UPDATE"
	err = tx.QueryRowContext(ctx,checkBalanceQuery,userId).Scan(&balanceNow,&reservedNow)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return 0,0,errors.New("user not found")
		}
		return 0,0,err
	}
	if balanceNow < amount{
		return 0,0,errors.New("not enough money")
	}
	changeBalanceQuery:="UPDATE users SET balance = ?,reserved = ? WHERE id = ?"
	_,err = tx.ExecContext(ctx,changeBalanceQuery,balanceNow - amount,reservedNow+amount, userId)
	if err!=nil{
		return 0,0,err
	}
	err=tx.Commit()
	if err!=nil{
		return 0,0,err
	}

	return reservedNow+amount,balanceNow-amount,nil
}


func (db *Repository) ReplenishmentOfBalance(ctx context.Context, userId int, amount float64)(float64,error){
	tx,err :=db.Database.BeginTx(ctx,nil)
	if err!= nil{
		return 0,err
	}
	defer tx.Rollback()
	queryString:="UPDATE users SET balance = balance+? WHERE id = ?"
	_,err = tx.ExecContext(ctx,queryString, amount,userId)
	if err!=nil{
		return 0,err
	}
	err=tx.Commit()
	if err!=nil{
		return 0,err
	}
	balanceAfter,_,err:=db.GetUserBalance(ctx,userId)
	if err!=nil{
		return 0,err
	}
	return balanceAfter,nil
}