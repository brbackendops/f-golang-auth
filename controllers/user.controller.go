package controllers

import (
	"encoding/json"
	"falcon/database/types"
	"falcon/utils"

	"github.com/gofiber/fiber/v2"

	srv "falcon/services"
)

type UserController struct {
	userSrv *srv.UserService
}

func UserNewController(userSrv *srv.UserService) *UserController {
	return &UserController{
		userSrv: userSrv,
	}
}

func (uc *UserController) RegisterHandler(c *fiber.Ctx) error {

	var UserRegister types.UserRegister

	data := c.Body()

	if err := json.Unmarshal([]byte(data), &UserRegister); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	user, err := uc.userSrv.SignUp(&UserRegister)
	if err != nil {

		if err, ok := err.(*utils.ModelExistsError); ok {
			return c.Status(err.StatusCode).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		return c.Status(500).JSON(&fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func (uc *UserController) LoginHandler(c *fiber.Ctx) error {
	var userData types.UserLogin

	// fmt.Println("Hello....from login ")
	data := c.Body()
	if err := json.Unmarshal(data, &userData); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	token, err := uc.userSrv.Login(&userData)

	if err != nil {

		if err, ok := err.(*utils.ModelDoesNotExistsError); ok {
			return c.Status(err.StatusCode).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid email or password.",
			})
		}

		return c.Status(500).JSON(&fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"status": "success",
		"token":  token,
	})

}
