package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"slack-app/internal/config"
	"slack-app/internal/slack/commands"

	"github.com/slack-go/slack"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	slackClient := slack.New(cfg.SlackBotToken)
	cmdHandler := commands.NewCommandHandler(slackClient)

	http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, cfg.SlackSigningSecret)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		bodyReader := io.TeeReader(r.Body, &verifier)
		body, err := io.ReadAll(bodyReader)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := verifier.Ensure(); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))
		cmd, err := slack.SlashCommandParse(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch cmd.Command {
		case "/signup":
			msg, err := cmdHandler.HandleSignup(cmd)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
		case "/deploy", "/status":
			msg, err := cmdHandler.HandleCommands(cmd)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(msg)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
