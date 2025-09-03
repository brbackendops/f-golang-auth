package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Field string `json:"field"`
}

func Validator[T any](userSchema T) fiber.Handler {

	request := userSchema
	validationErrors := []ErrorResponse{}

	return func(c *fiber.Ctx) error {

		if err := c.BodyParser(&request); err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"status": "error",
				"errors": "invalid request data",
			})
		}

		checks := validator.New()

		if errs := checks.Struct(request); errs != nil {

			for _, err := range errs.(validator.ValidationErrors) {
				var elem ErrorResponse

				elem.Error = "field is required"
				elem.Field = err.Field()

				validationErrors = append(validationErrors, elem)
			}

			return c.Status(400).JSON(&fiber.Map{
				"status": "error",
				"errors": validationErrors,
			})
		}

		return c.Next()
	}

}
