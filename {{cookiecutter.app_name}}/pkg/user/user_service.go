package user

import "github.com/gofiber/fiber"

type IUserService interface {
	FetchUser(c *fiber.Ctx) (user User, err error)
}

type Service struct {
	repo IUserRepository
}

func NewService(repo IUserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FetchUser(c *fiber.Ctx) (user User, err error) {
	user, err = s.repo.FindUser(c)
	return
}
