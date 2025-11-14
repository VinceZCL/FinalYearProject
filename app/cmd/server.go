package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"scrum.com/app"
	"scrum.com/internal/endpoint"
)

func init() {

	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run Go Echo server",
		Run: func(cmd *cobra.Command, args []string) {

			e := echo.New()

			app.SetupApp(e.AcquireContext())

			e.Use(middleware.Logger())
			e.Use(middleware.Recover())

			e.Logger.SetLevel(log.INFO)

			endpoint.RegisterRoutes(*e.AcquireContext().Echo())

			go func() {
				if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
					e.Logger.Fatalf("Shutting down the server: %v", err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
			<-quit

			// Context with timeout for shutdown
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			if err := e.Shutdown(shutdownCtx); err != nil {
				e.Logger.Fatalf("Server forced to shutdown: %v", err)
			}

			e.Logger.Infof("Shutting down the server gracefully...")

		},
	})
}
