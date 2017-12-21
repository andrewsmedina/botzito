package main

import (
	"fmt"
	"log"

	"github.com/andrewsmedina/botzito/telegram"
)

func main() {
	t := telegram.New()
	t.SendMessage("up")

	updates, err := t.GetUpdates()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(updates))
}
