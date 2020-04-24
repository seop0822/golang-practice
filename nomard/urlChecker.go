package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errorRequestFailed = errors.New("Request failed")

type result struct {
	url    string
	status string
}

func main() {

	c := make(chan result)
	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "Ok"
	if err != nil && resp.StatusCode >= 400 {
		status = "Failed"
	} else {
		c <- result{url: url, status: status}
	}

}
