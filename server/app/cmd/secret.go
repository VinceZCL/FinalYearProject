//go:build cli
// +build cli

package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "secret",
		Short: "Generate secret",
		Run: func(cmd *cobra.Command, args []string) {
			key := generateJWTSecretKey()

			// Path to .env.local, relative to where the command is run (i.e., 'server')
			envPath := ".env.local"

			// Open file for writing (will create the file if it doesn't exist)
			f, err := os.OpenFile(envPath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("Failed to open %s: %v", envPath, err)
			}
			defer f.Close()

			// Write the SECURITY_SECRETKEY line into the file
			line := fmt.Sprintf("export SECURITY_SECRETKEY=%s\n", key)
			if _, err := f.WriteString(line); err != nil {
				log.Fatalf("Failed to write secret: %v", err)
			}

			fmt.Printf("Secret written to server/%s\n", envPath)

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
