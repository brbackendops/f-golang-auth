package routers

import (
	"falcon/controllers"
	db "falcon/database"
	"falcon/database/types"
	middle "falcon/middlewares"
	UserRepo "falcon/repository"
	UserSrv "falcon/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutesInit(app *fiber.App) {

	userService := UserSrv.UserNewSerive(UserRepo.UserNewRepo(db.DB))

	cnt := controllers.UserNewController(userService)

	app.Route("/", func(api fiber.Router) {
		api.Post("/signup", middle.Validator(types.UserRegister{}), cnt.RegisterHandler)
		api.Post("/login", middle.Validator(types.UserLogin{}), cnt.LoginHandler)
	})
}
