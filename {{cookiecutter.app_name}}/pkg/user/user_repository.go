package user

import (
	dbpkg "gitlab.com/fdnetworks/api/{{cookiecutter.app_name}}/db"
	"gitlab.com/fdnetworks/api/{{cookiecutter.app_name}}/helper"
	"github.com/gofiber/fiber"
)

type IUserRepository interface {
	FindUser(c *fiber.Ctx) (user User, err error)
}

type Repository struct {
	model User
}

func NewRepository(model User) *Repository {
	return &Repository{model: model}
}

func (r *Repository) FindUser(c *fiber.Ctx) (user User, err error) {
	db := dbpkg.DBInstance(c)
	param, err := dbpkg.NewParameter(c, r.model)
	if err != nil {
		panic(err)
	}
	db = param.SetPreloads(db)
	fields := helper.ParseFields(dbpkg.DefaultValue(c, "fields"))
	queryField := helper.QueryFields(r.model, fields)

	if err = db.Select(queryField).Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
		panic(err)
	}
	return
}
