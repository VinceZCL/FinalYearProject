package interfaces

import (
	"scrum.com/internal/repository"
	"scrum.com/internal/service"
)

type Repositories struct {
	User repository.UserRepository
}

type Services struct {
	User service.UserService
}
