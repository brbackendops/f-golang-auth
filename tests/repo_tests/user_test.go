package repotests

import (
	"falcon/database/types"
	repo "falcon/repository"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	UserRepo repo.UserRepo
	Ctrl     *gomock.Controller
	Mock     sqlmock.Sqlmock
}

func (suite *UserRepoTestSuite) SetupSuite() {
	mockDB, mock, _ := sqlmock.New()
	sqlxDb := sqlx.NewDb(mockDB, "sqlmock")

	suite.UserRepo = *repo.UserNewRepo(sqlxDb)
	suite.Mock = mock

	suite.Ctrl = gomock.NewController(suite.T())
}

func (suite *UserRepoTestSuite) TestCreateUser() {

	data := types.UserRegister{
		Email:    "test@email.com",
		Password: "test123",
		Username: "test",
	}

	query := `INSERT INTO users (email,password,username) VALUES (?,?,?)`

	suite.Mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(data.Email, data.Password, data.Username).
		WillReturnResult(sqlmock.NewResult(1, 1))

	query = `
		SELECT id , email , username , created_at
		FROM users WHERE email=$1
	`

	columns := []string{"id", "email", "username", "created_at"}

	suite.Mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test@email.com", "test", time.Now()))

	userData, err := suite.UserRepo.CreateUser(&data)

	suite.NoError(err, "user should be created without any error")
	suite.Equal("test@email.com", userData.Email)

}

func (suite *UserRepoTestSuite) TestGetUserByEmail() {

	data := types.UserRegister{
		Email:    "test@email.com",
		Password: "test123",
		Username: "test",
	}

	query := "SELECT * FROM users WHERE email=$1"

	columns := []string{"id", "email", "username", "created_at"}

	suite.Mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test@email.com", "test", time.Now()))

	user, err := suite.UserRepo.GetByEmail(data.Email)

	suite.NoError(err, "should return a user with matched email")
	suite.Equal("test@email.com", user.Email)

}

func (suite *UserRepoTestSuite) TestGetUserById() {

	query := "SELECT * FROM users WHERE id=$1"

	columns := []string{"id", "email", "username", "created_at"}

	suite.Mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "test@email.com", "test", time.Now()))

	user, err := suite.UserRepo.GetById(1)

	suite.NoError(err, "should return a user with matched id")
	suite.Equal(1, user.Id)

}

func (suite *UserRepoTestSuite) TearDownSuite() {
	suite.UserRepo.DB.Close()
}

func TestUserRepo(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

// mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
