package interfaces

import (
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/internal/service"
)

type Repositories struct {
	User     repository.UserRepository
	UserTeam repository.UserTeamRepository
	Team     repository.TeamRepository
	CheckIn  repository.CheckInRepository
	Auth     repository.AuthRepository
}

type Services struct {
	User     service.UserService
	UserTeam service.UserTeamService
	Team     service.TeamService
	CheckIn  service.CheckInService
	Auth     service.AuthService
}
