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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid Request Payload",
		})
	}

	if err := validate.Struct(request); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {

			errorMap := make(map[string]string)
			for _, fieldError := range validationErrors {
				errorMap[fieldError.Field()] = fieldError.Error()
			}

			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": errorMap,
			})
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Failed To Validate Input",
		})
	}

	if isEmail(request.EmailOrIDSiswa) {
		penggunaAdmin, err := models.FindPenggunaByEmail(request.EmailOrIDSiswa)

		if err != nil {
			//return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			//	"errors": "Pengguna Tidak Ditemukan",
			//})
			if err.Error() == "Admin Tidak Ditemukan!" {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"errors": err.Error(),
				})
			}
		}

		if utils.CheckPasswordHash(request.Password, penggunaAdmin.Password) {
			jwt, err := config.GenerateJWT(penggunaAdmin)

			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"errors": err.Error(),
				})
			}

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Berhasil Login!",
				"data": fiber.Map{
					"token":         jwt,
					"nama_pengguna": penggunaAdmin.NamaPengguna,
				},
			})

		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": "Password Salah!",
			})
		}
	} else {

		penggunaSiswa, err := models.FindPenggunaByIDSiswa(request.EmailOrIDSiswa)

		if err != nil {
			if err.Error() == "Siswa Tidak Ditemukan!" {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"errors": err.Error(),
				})
			}
		}

		if utils.CheckPasswordHash(request.Password, penggunaSiswa.Password) {
			jwt, err := config.GenerateJWT(penggunaSiswa)

			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"errors": err.Error(),
				})
			}

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Berhasil Login!",
				"data": fiber.Map{
					"token":         jwt,
					"nama_pengguna": penggunaSiswa.NamaPengguna,
				},
			})
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": "Password Salah!",
			})
		}
	}
}

func isEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
