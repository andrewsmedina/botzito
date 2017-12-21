package telegram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Telegram represents a api object
type Telegram struct {
	token    string
	chatID   string
	endpoint string
}

// New returns a new Telegram instance
func New() *Telegram {
	return &Telegram{
		token:    os.Getenv("TOKEN"),
		chatID:   os.Getenv("CHAT_ID"),
		endpoint: "https://api.telegram.org",
	}
}

func (t *Telegram) formatURL(url string) string {
	return fmt.Sprintf(url, t.endpoint, t.token)
}

// GetUpdates returns the bot updates
func (t *Telegram) GetUpdates() ([]byte, error) {
	url := t.formatURL("%s/bot%s/getUpdates")
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

// SendMessage sends messages to a channel
func (t *Telegram) SendMessage(message string) error {
	u := t.formatURL("%s/bot%s/sendMessage")
	contentType := "application/x-www-form-urlencoded"
	values := url.Values{}
	values.Add("chat_id", t.chatID)
	values.Add("text", message)
	reader := strings.NewReader(values.Encode())
	_, err := http.Post(u, contentType, reader)
	return err
}
