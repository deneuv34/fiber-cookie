package db

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/serenize/snaker"

	"fmt"
	"log"
	"os"
	"strings"
	"time"

	{% if cookiecutter.database == "mysql" %}
	_ "github.com/jinzhu/gorm/dialects/mysql"
	{% else %}
	_ "github.com/jinzhu/gorm/dialects/postgres"
	{% endif %}
)

// Connect DB
func Connect() *gorm.DB {
	err := godotenv.Load()
	{% if cookiecutter.database == "mysql" %}
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	{% else %}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS")))
	{% endif %}
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(1)
	db.DB().SetConnMaxLifetime(3 * time.Second)

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	return db
}

func DBInstance(c *fiber.Ctx) *gorm.DB {
	return c.Locals("DB").(*gorm.DB)
}

func (p *Parameter) SetPreloads(db *gorm.DB) *gorm.DB {
	if p.Preloads == "" {
		return db
	}

	for _, preload := range strings.Split(p.Preloads, ",") {
		var a []string

		for _, s := range strings.Split(preload, ".") {
			a = append(a, snaker.SnakeToCamel(s))
		}

		db = db.Preload(strings.Join(a, "."))
	}

	return db
}
