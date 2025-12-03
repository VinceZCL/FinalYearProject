package app

import (
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

func (a *appClient) getDBClient(c echo.Context) *client.PostgresClient {
	if a.DBClient == nil {
		var err error
		c.Logger().Info("Initializing to postgres")
		a.DBClient, err = client.NewPostgres()
		if err != nil {
			c.Logger().Fatal("Unable to connect to postgres")
			panic(err)
		}
		c.Logger().Info("Connected to postgres")
	}
	return a.DBClient
}

func (a *appClient) getDB(c echo.Context) *gorm.DB {
	if a.DB == nil {
		a.DB = a.getDBClient(c).DB
	}
	return a.DB
}

func WithApp(ctx echo.Context, a *app) {
	ctx.Set(AppContextKey, a)
}

func FromContext(ctx echo.Context) *app {
	a := ctx.Get(AppContextKey)
	if a != nil {
		return a.(*app)
	}

	return New()
}

func SetupApp(c echo.Context) {
	contextApp := FromContext(c)
	contextApp.Client.getDB(c)
	contextApp.Client.getDBClient(c)
	contextApp.Repos = &interfaces.Repositories{
		User:     repository.NewUserRepository(contextApp.Client.DBClient),
		UserTeam: repository.NewUserTeamRepository(contextApp.Client.DBClient),
		Team:     repository.NewTeamRepository(contextApp.Client.DBClient),
		CheckIn:  repository.NewCheckInRepository(contextApp.Client.DBClient),
	}
	contextApp.Services = &interfaces.Services{
		User:     *service.NewUserService(contextApp.Repos.User),
		UserTeam: *service.NewUserTeamService(contextApp.Repos.UserTeam),
		Team:     *service.NewTeamRepository(contextApp.Repos.Team),
		CheckIn:  *service.NewCheckInService(contextApp.Repos.CheckIn),
	}
	WithApp(c, contextApp)
}
