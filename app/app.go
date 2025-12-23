package app

import (
	"fmt"

	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/internal/interfaces"
	"github.com/VinceZCL/FinalYearProject/internal/repository"
	"github.com/VinceZCL/FinalYearProject/internal/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const AppContextKey = "fyp:app"

type AppOption func(*appClient)

type app struct {
	Client   *appClient
	Repos    *interfaces.Repositories
	Services *interfaces.Services
}

type appClient struct {
	DB       *gorm.DB
	DBClient *client.PostgresClient
}

var AppClientInstance *appClient

func New(opts ...AppOption) *app {

	if AppClientInstance == nil {
		AppClientInstance = newAppClient(opts...)
	}
	return &app{
		Client: AppClientInstance,
	}
}

func newAppClient(opts ...AppOption) *appClient {
	a := &appClient{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func FromContext(ctx echo.Context) *app {
	return ctx.Get(AppContextKey).(*app)
}

func SetupApp(a *app) *app {
	// --- Initialize DB Client ---
	if a.Client.DBClient == nil {
		dbClient, err := client.NewPostgres()
		if err != nil {
			panic(fmt.Errorf("failed to connect to postgres: %w", err))
		}
		a.Client.DBClient = dbClient
		a.Client.DB = dbClient.DB // gorm.DB
	}

	// --- Initialize Repositories ---
	if a.Repos == nil {
		a.Repos = &interfaces.Repositories{
			User:     repository.NewUserRepository(a.Client.DBClient),
			UserTeam: repository.NewUserTeamRepository(a.Client.DBClient),
			Team:     repository.NewTeamRepository(a.Client.DBClient),
			CheckIn:  repository.NewCheckInRepository(a.Client.DBClient),
			Auth:     repository.NewAuthRepository(a.Client.DBClient),
		}
	}

	// --- Initialize Services ---
	if a.Services == nil {
		a.Services = &interfaces.Services{
			User:     *service.NewUserService(a.Repos.User),
			UserTeam: *service.NewUserTeamService(a.Repos.UserTeam),
			Team:     *service.NewTeamService(a.Repos.Team),
			CheckIn:  *service.NewCheckInService(a.Repos.CheckIn),
			Auth:     *service.NewAuthService(a.Repos.Auth, a.Repos.User),
		}
	}

	return a
}
