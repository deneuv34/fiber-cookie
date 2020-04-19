package routes

import (
	"gitlab.com/fdnetworks/api/{{cookiecutter.app_name}}/pkg/user"
	"github.com/gofiber/fiber"
)

func SetupRoute(app *fiber.App) {
	userProvider := user.SetProvider()
	app.Get("/user", userProvider.Handler.ListUser)
}
