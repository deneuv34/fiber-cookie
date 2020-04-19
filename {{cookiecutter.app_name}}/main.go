package main

import (
	"gitlab.com/fdnetworks/api/{{cookiecutter.app_name}}/boot"
)

func main() {
	app := boot.InitServer()

	app.Listen(8000)
}
