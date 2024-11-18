package controllers

import (
	"backend-ujian-gofiber/src/config"
	"backend-ujian-gofiber/src/dto"
	"backend-ujian-gofiber/src/models"
	"backend-ujian-gofiber/src/utils"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"regexp"
)

var validate = validator.New()

func Login(ctx *fiber.Ctx) error {
	var request dto.LoginRequest

	if err := ctx.BodyParser(&request); err != nil {
		return badRequest(ctx, "Invalid Request Payload!")
	}

	if err := validate.Struct(request); err != nil {
		return validationError(ctx, err)
	}

	user, err := findUser(request.EmailOrIDSiswa)
	if err != nil {
		return badRequest(ctx, err.Error())
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return badRequest(ctx, "Password Salah!")
	}

	token, err := config.GenerateJWT(user)
	if err != nil {
		return badRequest(ctx, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Berhasil Login!",
		"data": fiber.Map{
			"token":         token,
			"nama_pengguna": user.NamaPengguna,
			"role_pengguna": user.RolePengguna,
		},
	})
}

func isEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func findUser(identifier string) (models.Pengguna, error) {
	if isEmail(identifier) {
		return models.FindPenggunaByEmail(identifier)
	}
	return models.FindPenggunaByIDSiswa(identifier)
}

func badRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"code":    fiber.StatusBadRequest,
		"message": message,
	})
}

func validationError(ctx *fiber.Ctx, err error) error {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		errorMap := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMap[fieldError.Field()] = fieldError.Error()
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Validasi Error",
			"errors":  errorMap,
		})
	}
	return badRequest(ctx, "Failed To Validate Input")
}
