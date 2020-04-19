package user

type Provider struct {
	Handler IUserHandler
}

func SetProvider() *Provider {
	repo := NewRepository(User{})
	service := NewService(repo)
	handler := NewHandler(service)

	return &Provider{Handler: handler}
}
