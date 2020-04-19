package db

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber"
)

const (
	defaultLimit = "25"
	defaultPage  = "1"
	defaultOrder = "desc"
)

type Parameter struct {
	Filters  map[string]string
	Preloads string
	Sort     string
	Limit    int
	Page     int
	LastID   int
	Order    string
	IsLastID bool
}

func NewParameter(c *fiber.Ctx, model interface{}) (*Parameter, error) {
	parameter := &Parameter{}

	if err := parameter.initialize(c, model); err != nil {
		return nil, err
	}

	return parameter, nil
}

func DefaultValue(c *fiber.Ctx, queryParam string) string {
	var val string

	switch queryParam {
	case "limit":
		if c.Query("limit") != "" {
			val = c.Query("limit")
		}
		val = defaultLimit

	case "page":
		if c.Query("page") != "" {
			val = c.Query("page")
		}
		val = defaultPage

	case "order":
		if c.Query("order") != "" {
			val = c.Query("order")
		}
		val = defaultOrder

	case "fields":
		if c.Query("fields") != "" {
			val = c.Query("fields")
		}
		val = "*"
	}
	return val
}

func (p *Parameter) initialize(c *fiber.Ctx, model interface{}) error {
	p.Filters = filterToMap(c, model)
	p.Preloads = c.Query("preloads")
	p.Sort = c.Query("sort")

	limit, err := validate(DefaultValue(c, "limit"))
	if err != nil {
		return err
	}

	p.Limit = int(math.Max(1, math.Min(10000, float64(limit))))
	page, err := validate(DefaultValue(c, "limit"))
	if err != nil {
		return err
	}

	p.Page = int(math.Max(1, float64(page)))
	lastID, err := validate(c.Query("last_id"))
	if err != nil {
		return err
	}

	if lastID != -1 {
		p.IsLastID = true
		p.LastID = int(math.Max(0, float64(lastID)))
	}

	p.Order = DefaultValue(c, "order")
	return nil
}

func validate(s string) (int, error) {
	if s == "" {
		return -1, nil
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}
