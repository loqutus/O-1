package telegram

import (
	"errors"
	"os"
)

func getToken() (string, error) {
	token := os.Getenv("O1_TELEGRAM_TOKEN")
	if token == "" {
		return "", errors.New("O1_TELEGRAM_TOKEN is not set")
	} else {
		return token, nil
	}
}
