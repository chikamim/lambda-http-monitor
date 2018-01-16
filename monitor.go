package main

import (
	"errors"
	"net/http"

	"github.com/gregdel/pushover"
)

var (
	ErrResponse = errors.New("HTTP error response")
)

func checkStatus(url string) error {
	res, err := http.Head(url)
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
