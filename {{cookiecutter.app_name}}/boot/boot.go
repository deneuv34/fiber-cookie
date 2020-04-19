package boot

import (
	dbpkg "gitlab.com/fdnetworks/api/{{cookiecutter.app_name}}/db"

	"gitlab.com/fdnetworks/api/{{cookiecutter.app_name}}/middleware"
	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func InitServer() *fiber.App {
	s := fiber.New()
	db := dbpkg.Connect()

	s.Use(middleware.SetDB(db))
	s.Use(compression.New())
	s.Use(logger.New())

	return s
}
