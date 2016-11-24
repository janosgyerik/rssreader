package rssreader

import (
	rss "github.com/jteeuwen/go-pkg-rss"
	"io"
	"log"
)

type rssReader struct {
	uri      string
	feed     Feed
	rssFeed  *rss.Feed
	newPosts []Post
}

func NewRssReader(uri string, feed Feed) FeedReader {
	reader := rssReader{uri: uri, feed: feed}
	timeout := 0
	reader.rssFeed = rss.New(timeout, true, reader.chanHandler, reader.itemHandler)
	return &reader
}

func (reader *rssReader) chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	log.Printf("%d new channel(s) in %s", len(newchannels), reader.feed.Id)
}

func (reader *rssReader) itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	posts := make([]Post, len(newitems))
	for i, item := range newitems {
		id := extractPostId(item)
		post := Post{
			Id:      id,
			Url:     id,
			Author:  item.Author.Name,
			Subject: item.Title,
			Body:    item.Description,
			Feed:    &reader.feed,
		}
		posts[i] = post
	}

	reader.newPosts = posts
}

func (reader *rssReader) GetFeed() Feed {
	return reader.feed
}

func (reader *rssReader) FetchNewPosts() []Post {
	reader.newPosts = nil

	// note: itemHandler will get called synchronously when there are new posts
	if err := reader.rssFeed.Fetch(reader.uri, charsetReader); err != nil {
		log.Printf("error: %s: %s", reader.uri, err)
		return []Post{}
	}

	return reader.newPosts
}

func charsetReader(charset string, r io.Reader) (io.Reader, error) {
	return r, nil
}

// RSS feeds store the post id in different non-standard places
func extractPostId(item *rss.Item) string {
	if len(item.Id) > 0 {
		return item.Id
	}
	for _, link := range item.Links {
		if len(link.Href) > 0 {
			return link.Href
		}
	}
	return ""
}
