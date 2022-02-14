package middleware

import "github.com/gofiber/fiber/v2"

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("key-access")
	if token != "Indra19" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return ctx.Next()
}

func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
