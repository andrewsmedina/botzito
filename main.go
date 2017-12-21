package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var token = os.Getenv("TOKEN")

func getUpdates() ([]byte, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}

func main() {
	updates, err := getUpdates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(updates))
}
