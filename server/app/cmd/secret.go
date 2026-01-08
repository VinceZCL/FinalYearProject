//go:build cli
// +build cli

package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "secret",
		Short: "Generate secret",
		Run: func(cmd *cobra.Command, args []string) {
			key := generateJWTSecretKey()
			fmt.Println("Add this into `./server/.env.local`")
			fmt.Printf("export SECURITY_SECRETKEY=%s\n", key)
		},
	})
}

func generateJWTSecretKey() string {
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatalf("Error generating secret: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(secret)
}
