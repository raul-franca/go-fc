package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second} // cria o objeto http

	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	// setando o header
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	println(string(body))

}
