package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	token  = os.Getenv("TOKEN")
	chatID = os.Getenv("CHAT_ID")
)

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

func sendMessage(chatID, text string) error {
	u := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	contentType := "application/x-www-form-urlencoded"
	values := url.Values{}
	values.Add("chat_id", chatID)
	values.Add("text", text)
	reader := strings.NewReader(values.Encode())
	_, err := http.Post(u, contentType, reader)
	return err
}

func main() {
	sendMessage(chatID, "up")
	updates, err := getUpdates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(updates))
}
