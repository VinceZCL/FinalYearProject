//go:build server

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "Run Go Echo server",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Ran Server")

		},
	})
}
