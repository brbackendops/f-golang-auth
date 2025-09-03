package repotests

import (
	"falcon/database/types"
	mockUser "falcon/repository/mocks"
	srv "falcon/services"
	"falcon/utils"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type UserSrvTestSuite struct {
	suite.Suite
	UserRepo *mockUser.MockUserI
	UserSrv  srv.UserService
	Ctrl     *gomock.Controller
}

func (suite *UserSrvTestSuite) SetupSuite() {
	suite.Ctrl = gomock.NewController(suite.T())
	suite.UserRepo = mockUser.NewMockUserI(suite.Ctrl)
	suite.UserSrv = *srv.UserNewSerive(suite.UserRepo)
}

func (suite *UserSrvTestSuite) TestUserLogin() {

	hashedP, _ := utils.Hash("test123")

	user := types.UserAll{
		Id:        1,
		Email:     "test@mail.com",
		Username:  "test",
		Password:  hashedP,
		CreatedAt: time.Now().String(),
	}

	data := types.UserLogin{
		Email:    "test@mail.com",
		Password: "test123",
	}

	suite.UserRepo.EXPECT().GetByEmail(data.Email).Return(&user, nil)

	token, err := suite.UserSrv.Login(&data)

	suite.NoError(err, "user should be retrieved successfuly with given email")
	suite.Equal("test@mail.com", user.Email)

	suite.Equal(true, utils.JWTVerify(token))
}

func (suite *UserSrvTestSuite) TestUserRegister() {

	user := types.User{
		Email:     "test@mail.com",
		Username:  "test",
		CreatedAt: time.Now().String(),
	}

	data := types.UserRegister{
		Email:    "test@mail.com",
		Username: "test",
		Password: "test123",
	}

	suite.UserRepo.EXPECT().CreateUser(&data).Return(&user, nil)

	userData, err := suite.UserSrv.SignUp(&data)

	suite.NoError(err, "user should be retrieved successfuly with given email")
	suite.Equal("test@mail.com", userData.Email)
	suite.Equal("test", userData.Username)
}

func (suite *UserSrvTestSuite) TestShouldThrowExistsError() {

	data := types.UserRegister{
		Email:    "test@mail.com",
		Username: "test",
		Password: "test123",
	}

	suite.UserRepo.EXPECT().CreateUser(&data).Return(nil, &pq.Error{
		Code:    "23505",
		Message: "unique voilation",
	})

	_, err := suite.UserSrv.SignUp(&data)
	err, ok := err.(*utils.ModelExistsError)

	suite.Error(err)
	suite.Equal(true, ok)

}

func (suite *UserSrvTestSuite) TestShouldThrowNotExistsError() {

	data := types.UserLogin{
		Email:    "test@mail.com",
		Password: "test123",
	}

	suite.UserRepo.EXPECT().GetByEmail(data.Email).Return(nil, &pq.Error{
		Code:    "42703",
		Message: "undefined column",
	})

	_, err := suite.UserSrv.Login(&data)
	err, ok := err.(*utils.ModelDoesNotExistsError)

	suite.Error(err)
	suite.Equal(true, ok)

}

func (suite *UserSrvTestSuite) TearDownSuite() {
}

func TestUserSrv(t *testing.T) {
	suite.Run(t, new(UserSrvTestSuite))
}
