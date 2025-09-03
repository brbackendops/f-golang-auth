package repository

import (
	"falcon/database/types"
)

type UserI interface {
	GetById(userId int) (*types.UserAll, error)
	GetByEmail(email string) (*types.UserAll, error)
	CreateUser(*types.UserRegister) (*types.User, error)
}
