package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func Problem9() {
	url := "https://google.com"
	body, err := FetchURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Body:", body)
}
