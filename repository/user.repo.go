package repository

import (
	"falcon/database/types"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func UserNewRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) CreateUser(data *types.UserRegister) (*types.User, error) {

	var userData types.UserRegister = types.UserRegister{
		Email:    data.Email,
		Password: data.Password,
		Username: data.Username,
	}

	query := `
		INSERT INTO users (email,password,username) 
		VALUES (:email,:password,:username)
	`

	_, err := u.DB.NamedExec(query, &userData)

	if err != nil {
		return nil, err
	}

	query = `
		SELECT id , email , username , created_at
		FROM users WHERE email=$1
	
	`

	user := types.User{}
	err = u.DB.Get(&user, query, data.Email)

	return &user, err

}

func (u *UserRepo) GetById(userId int) (*types.UserAll, error) {
	user := types.UserAll{}

	query := `
		SELECT *
		FROM users WHERE id=$1	
	`

	err := u.DB.Get(&user, query, userId)

	return &user, err
}

func (u *UserRepo) GetByEmail(email string) (*types.UserAll, error) {

	user := types.UserAll{}

	query := `
		SELECT *
		FROM users WHERE email=$1	
	`

	err := u.DB.Get(&user, query, email)

	return &user, err
}
