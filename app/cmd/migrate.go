//go:build cli
// +build cli

package cmd

import (
	"fmt"
	"log"

	"github.com/VinceZCL/FinalYearProject/internal/client"
	"github.com/VinceZCL/FinalYearProject/types/model"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Run DB migration",
		Run: func(cmd *cobra.Command, args []string) {

			pg, err := client.NewPostgres()
			if err != nil {
				log.Fatalf("Failed to connect to DB: %v", err)
			}
			db := pg.DB

			if err := db.AutoMigrate(
				&model.User{},
				&model.Team{},
				&model.UserTeam{},
				&model.CheckIn{},
			); err != nil {
				log.Fatalf("Migration failed: %v", err)
			}

			fmt.Println("Migration completed successfully!")

		},
	})
}
