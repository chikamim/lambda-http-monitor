package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gregdel/pushover"
)

var (
	ErrResponse = errors.New("HTTP error response")
)

func checkStatus(url string, timeout time.Duration) error {
	client := http.Client{
		Timeout: timeout,
	}
	res, err := client.Head(url)
	if err != nil {
		return err
	}
	if res.StatusCode > 400 {
		return ErrResponse
	}
	return nil
}

func pushOver(apikey, userkey, message string) error {
	app := pushover.New(apikey)
	r := pushover.NewRecipient(userkey)
	m := pushover.NewMessage(message)
	_, err := app.SendMessage(m, r)
	if err != nil {
		return err
	}
	return nil
}
