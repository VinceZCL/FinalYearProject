package interfaces

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/internal/service"
)

type Repositories struct {
	User repository.UserRepository
}

type Services struct {
	User service.UserService
}
