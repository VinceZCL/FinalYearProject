//go:build cron
// +build cron

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/VinceZCL/FinalYearProject/app"
	"github.com/VinceZCL/FinalYearProject/app/config"
	"github.com/spf13/cobra"
)

// TeamsMessage represents the payload for Teams workflow webhook
type TeamsMessage struct {
	Type        string `json:"type"`
	Attachments []struct {
		ContentType string       `json:"contentType"`
		ContentURL  *string      `json:"contentUrl"`
		Content     AdaptiveCard `json:"content"`
	} `json:"attachments"`
}

// AdaptiveCard defines the structure of the card
type AdaptiveCard struct {
	Schema  string        `json:"$schema"`
	Type    string        `json:"type"`
	Version string        `json:"version"`
	Body    []interface{} `json:"body"`
	Actions []interface{} `json:"actions,omitempty"`
}

var empty = []string{
	"You can be the first!",
	"It is time to lock in.",
	"Someone has to go first.",
	"Be the one who starts it.",
	"The floor is yours.",
	"Are we cooking or are we cooked?",
}

type block struct {
	user    string
	blocker string
}

var blocks []block
var txt string
var subtxt string

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "cron",
		Short: "Run daily Cron",
		Run: func(cmd *cobra.Command, args []string) {

			cronApp := app.SetupApp(app.New())

			webhookUrl := config.Get().Cron.Webhook

			teamID := 1
			loc, err := time.LoadLocation(config.Get().Database.Location)
			if err != nil {
				panic(err)
			}

			today := time.Now().In(loc).Format(time.DateOnly)

			cis, err := cronApp.Services.CheckIn.GetTeamCheckIns(nil, uint(teamID), today)
			if err != nil {
				panic(err)
			}

			numCI := len(cis)

			if numCI < 1 {
				txt = fmt.Sprintf("Nobody submitted their Daily Check-Ins.\n%s\n", empty[rand.Intn(len(empty))])
			} else {
				for _, v := range cis {
					if len(v.Blockers) >= 1 {
						for _, b := range v.Blockers {
							blocks = append(blocks, block{
								user:    v.Username,
								blocker: b.Item,
							})
						}
					}
				}
				if len(blocks) < 1 {
					subtxt = "No blockers found.\n"
				} else {
					subtxt = fmt.Sprintf("\n%d blockers found:\n", len(blocks))
					for _, v := range blocks {
						subtxt += fmt.Sprintf("• %s: %s\n", v.user, v.blocker)
					}
				}
				txt = fmt.Sprintf("%d has submitted their Daily Check-Ins.\n%s", numCI, subtxt)
			}

			// Build the card
			card := AdaptiveCard{
				Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
				Type:    "AdaptiveCard",
				Version: "1.4",
				Body: []interface{}{
					map[string]interface{}{
						"type":   "TextBlock",
						"text":   "Daily Check-Ins",
						"weight": "Bolder",
						"size":   "Medium",
					},
					// Date
					map[string]interface{}{
						"type":     "TextBlock",
						"text":     today,
						"size":     "Small",
						"isSubtle": true,
						"spacing":  "None",
					},
					map[string]interface{}{
						"type": "TextBlock",
						"text": txt + "\nSincerely, Vince.",
						"wrap": true,
					},
				},
				Actions: []interface{}{
					map[string]interface{}{
						"type":  "Action.OpenUrl",
						"title": "Daily Check-Ins",
						"url":   "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
					},
				},
			}

			msg := TeamsMessage{
				Type: "message",
			}
			msg.Attachments = append(msg.Attachments, struct {
				ContentType string       `json:"contentType"`
				ContentURL  *string      `json:"contentUrl"`
				Content     AdaptiveCard `json:"content"`
			}{
				ContentType: "application/vnd.microsoft.card.adaptive",
				ContentURL:  nil,
				Content:     card,
			})

			// Marshal payload
			payload, err := json.Marshal(msg)
			if err != nil {
				panic(err)
			}

			// Send POST request to the webhook
			resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(payload))
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			fmt.Printf("Message sent, status: %s\n", resp.Status)

		},
	})
}
