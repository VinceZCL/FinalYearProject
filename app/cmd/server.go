package cmd

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"scrum.com/app"
	"scrum.com/internal/service"
)

func init() {

	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run Go Echo server",
		Run: func(cmd *cobra.Command, args []string) {

			e := echo.New()

			ctx := app.WithApp(context.Background(), app.New(app.WithLogger(e.Logger)))
			app.SetupApp(ctx)

			e.Use(middleware.Logger())
			e.Use(middleware.Recover())

			e.GET("/", service.Hello)

			e.Logger.Fatal(e.Start(":8080"))

		},
	})
}
