package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"scrum.com/internal/client"
	"scrum.com/types/models"
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
				&models.User{},
				&models.Team{},
				&models.UserTeam{},
				&models.CheckIn{},
			); err != nil {
				log.Fatalf("Migration failed: %v", err)
			}

			fmt.Println("Migration completed successfully!")

		},
	})
}
