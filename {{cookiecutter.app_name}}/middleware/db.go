package middleware

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	es "github.com/olivere/elastic/v7"
)

func SetDB(db *gorm.DB) func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.Locals("DB", db)
		c.Next()
	}
}

func SetES(es *es.Client) func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		c.Locals("ES", es)
		c.Next()
	}
}
