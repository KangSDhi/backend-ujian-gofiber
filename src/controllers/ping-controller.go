package controllers

import "github.com/gofiber/fiber/v2"

func Ping(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"ping": "pong",
	})
}
