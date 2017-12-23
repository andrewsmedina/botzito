package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andrewsmedina/botzito/telegram"
	"github.com/robfig/cron"
)

func main() {
	t := telegram.New()
	t.SendMessage("up")

	updates, err := t.GetUpdates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(updates))

	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() { t.SendMessage("time to sleep") })
	c.AddFunc("0 0 9 * * *", func() { t.SendMessage("time to wake up") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
