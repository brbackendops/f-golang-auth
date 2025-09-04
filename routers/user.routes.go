package routers

import (
	"falcon/controllers"
	db "falcon/database"
	"falcon/database/types"
	middle "falcon/middlewares"
	UserRepo "falcon/repository"
	UserSrv "falcon/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func UserRoutesInit(app *fiber.App) {

	userService := UserSrv.UserNewSerive(UserRepo.UserNewRepo(db.DB))

	cnt := controllers.UserNewController(userService)

	app.Route("/", func(api fiber.Router) {
		api.Post("/signup",
			limiter.New(limiter.Config{
				Max:        25,
				Expiration: 30 * time.Minute,
				LimitReached: func(c *fiber.Ctx) error {
					return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
						"status": "error",
						"error":  "Too many requests, please try again later",
					})
				},
			}),
			middle.Validator(types.UserRegister{}),
			cnt.RegisterHandler,
		)
		api.Post("/login", middle.Validator(types.UserLogin{}), cnt.LoginHandler)
	})
}
