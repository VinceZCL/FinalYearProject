//go:build cli
// +build cli

package cmd

import (
	"fmt"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "default",
		Short: "Generate default user",
		Run: func(cmd *cobra.Command, args []string) {
			generateDefault()
		},
	})
}

func generateDefault() {
	fmt.Printf("Creating default admin\n")
	// TODO generate default admin account
	app := app.SetupApp(app.New())
	err := app.Services.Auth.Default()
	if err != nil {
		fmt.Printf("Default admin already exists.\n")
		fmt.Printf("Aborting.\n")
		return
	}
	fmt.Printf("Default admin created.\n")
	fmt.Printf("Email: %s\nPassword: %s\n", "admin@example.com", "admin")
	fmt.Printf("Please update credentials for security.\n")
}
