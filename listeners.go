package rssreader

import (
	"fmt"
	"log"
)

type ConsolePrinterListener struct{}

func (listener ConsolePrinterListener) OnPost(post Post) {
	log.Printf(formatPost(post))
}

func formatPost(post Post) string {
	return fmt.Sprintf("(%s) %s %s", post.Feed.Id, post.Id, post.Subject)
}
