package main

import (
	"fmt"
	"io"
	"net/http"
)

type Sender struct {
	HttpClient *http.Client
}

func NewSender() {
	sender := new(Sender)
	sender.HttpClient = &http.Client{}
}

func (sender *Sender) SendRequest(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := sender.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}
