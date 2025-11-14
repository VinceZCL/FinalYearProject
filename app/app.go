package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"scrum.com/internal/client"
)

const appContextKey = "scrum:app"

type AppOption func(*appClient)

type app struct {
	Client *appClient
}

type appClient struct {
	DB       *gorm.DB
	DBClient *client.PostgresClient
	Logger   echo.Logger
}

func New(opts ...AppOption) *app {
	appClient := newAppClient(opts...)
	return &app{
		Client: appClient,
	}
}

func newAppClient(opts ...AppOption) *appClient {
	a := &appClient{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func (a *appClient) getDBClient() *client.PostgresClient {
	if a.DBClient == nil {
		var err error
		a.Logger.Info("Initializing to postgres")
		a.DBClient, err = client.NewPostgres()
		if err != nil {
			a.Logger.Fatal("Unable to connect to postgres")
			panic(err)
		}
		a.Logger.Info("Connected to postgres")
	}
	return a.DBClient
}

func (a *appClient) getDB() *gorm.DB {
	if a.DB == nil {
		a.DB = a.getDBClient().DB
	}
	return a.DB
}

func WithApp(ctx context.Context, a *app) context.Context {
	return context.WithValue(ctx, appContextKey, a)
}

func WithLogger(logger echo.Logger) AppOption {
	return func(c *appClient) {
		c.Logger = logger
	}
}

func FromContext(ctx context.Context) *app {
	a := ctx.Value(appContextKey)
	if a != nil {
		return a.(*app)
	}

	return New()
}

func SetupApp(ctx context.Context) {
	contextApp := FromContext(ctx)
	contextApp.Client.getDB()
	contextApp.Client.getDBClient()
}
