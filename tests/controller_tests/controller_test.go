package repotests

import (
	"bytes"
	"encoding/json"
	"falcon/app"
	cnt "falcon/controllers"
	mocks "falcon/repository/mocks"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	srv "falcon/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

var fiberApp *fiber.App

type UserHandlerTestSuite struct {
	suite.Suite
	UserRepo *mocks.MockUserI
	UserSrv  srv.UserService
	Handler  cnt.UserController
	Ctrl     *gomock.Controller
	DB       *sqlx.DB
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "test")

	suite.Ctrl = gomock.NewController(suite.T())

	suite.UserRepo = mocks.NewMockUserI(suite.Ctrl)
	suite.UserSrv = *srv.UserNewSerive(suite.UserRepo)
	suite.Handler = *cnt.UserNewController(&suite.UserSrv)

	fiberApp = app.AppInstance()

}

func (suite *UserHandlerTestSuite) TestGetHealthRoute() {

	suite.T().Log("testing on path /health")

	req, _ := http.NewRequest("GET", "/health", nil)
	res, err := fiberApp.Test(req, -1)

	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(body)

	suite.NoError(err)
	suite.Equal(200, res.StatusCode)

}

func (suite *UserHandlerTestSuite) TestingRegisterAndLogin() {

	suite.T().Log("Integration test on register and login")

	// register

	data := map[string]any{
		"email":    "test@mail.com",
		"username": "test",
		"password": "test123",
	}

	dataJson, err := json.Marshal(data)

	suite.NoError(err)

	req, err := http.NewRequest("POST", "/signup", bytes.NewReader(dataJson))
	req.Header.Set("Content-Type", "application/json")

	suite.NoError(err)

	res, err := fiberApp.Test(req, -1)
	suite.NoError(err)

	suite.Equal(201, res.StatusCode)

	// login

	data = map[string]any{
		"email":    "test@mail.com",
		"password": "test123",
	}

	dataJson, err = json.Marshal(data)

	suite.NoError(err)

	req, err = http.NewRequest("POST", "/login", bytes.NewReader(dataJson))
	fmt.Println("request", req)

	req.Header.Set("Content-Type", "application/json")

	suite.NoError(err)

	res, err = fiberApp.Test(req, -1)
	suite.NoError(err)

	body, err := io.ReadAll(res.Body)
	suite.NoError(err)

	fmt.Println(string(body))

	suite.Equal(200, res.StatusCode)

	defer res.Body.Close()

}

func (suite *UserHandlerTestSuite) TearDownSuite() {
	fmt.Println()
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
