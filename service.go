package main

import (
	"fmt"
	"net/url"

	"github.com/teris-io/shortid"
)

func shortenURL(longURL string) (string, error) {
	u, err := url.ParseRequestURI(longURL)
	if err != nil {
		return "", err
	}

	fmt.Println(u.String())

	id, err := shortid.Generate()
	if err != nil {
		return "", err
	}

	return id, nil
}

func getURLByCode(code string) string {
	return "GetURLByCode"
}
