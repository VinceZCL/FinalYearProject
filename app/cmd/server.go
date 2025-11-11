package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"scrum.com/internal/service"
)

func init() {

	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run Go Echo server",
		Run: func(cmd *cobra.Command, args []string) {

			e := echo.New()
			e.GET("/", service.Hello)

			e.Logger.Fatal(e.Start(":8080"))

			// TODO Introduce Logger
			// TODO Introduce graceful shutdown

		},
	})
}
