package services

import (
	"database/sql"
	"errors"
	"falcon/database/types"
	UserRepo "falcon/repository"
	"falcon/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
)

type UserService struct {
	UserRepo UserRepo.UserI
}

func UserNewSerive(UserRepo UserRepo.UserI) *UserService {
	return &UserService{
		UserRepo: UserRepo,
	}
}

func (us *UserService) Login(data *types.UserLogin) (string, error) {

	var pqError *pq.Error

	user, err := us.UserRepo.GetByEmail(data.Email)
	if err != nil {

		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", &utils.ModelDoesNotExistsError{
				StatusCode: 404,
				ModelName:  "user",
			}
		case errors.As(err, &pqError):
			if pqError.Code.Name() == "undefined_column" {
				return "", &utils.ModelDoesNotExistsError{
					StatusCode: 404,
					ModelName:  "user",
				}
			}
		}
	}

	if !utils.CompareHash(data.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.JWTSign(jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) SignUp(data *types.UserRegister) (*types.User, error) {

	hPassword, err := utils.Hash(data.Password)
	if err != nil {
		return nil, err
	}

	data.Password = hPassword

	user, err := us.UserRepo.CreateUser(data)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return nil, &utils.ModelExistsError{
					StatusCode: 409,
					ModelName:  "user",
					Cause:      "email",
				}
			}
		}
	}

	return user, err
}
