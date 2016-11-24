package rssreader

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

// poll RSS feeds once per 15 minutes
const rssPollingMillis = 1000 * 60 * 15

type Post struct {
	Id      string
	Url     string
	Author  string
	Subject string
	Body    string
	Feed    *Feed
}

type Feed struct {
	Id  string
	Url string
}

type FeedReader interface {
	GetFeed() Feed
	FetchNewPosts() []Post
}

type Listener interface {
	OnPost(Post)
}

type Config struct {
	Feeds     []Feed
}

func ParseConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

type Context struct {
	Readers        []FeedReader
	Listeners      []Listener
}

func ParseContext(config *Config) (*Context, error) {
	if len(config.Feeds) == 0 {
		return nil, errors.New("configuration error: configure at least one feed")
	}

	readers := parseReaders(config)

	listeners := []Listener{ConsolePrinterListener{}}

	context := &Context{
		Readers:        readers,
		Listeners:      listeners,
	}
	return context, nil
}

func parseReaders(config *Config) []FeedReader {
	readers := make([]FeedReader, len(config.Feeds))
	for i, feed := range config.Feeds {
		log.Println("adding feed:", feed.Id, feed.Url)
		readers[i] = NewRssReader(feed.Url, feed)
	}
	return readers
}

// the default number of posts to read; normally infinity, set to 0 by some tests
var defaultCount int

func init() {
	defaultCount = getDefaultCount()
}

func getDefaultCount() int {
	maxUint := ^uint(0)
	maxInt := int(maxUint >> 1)
	return maxInt
}

func RunForever(path string) error {
	config, err := ParseConfig(path)
	if err != nil {
		return err
	}

	return runForever(config)
}

func runForever(config *Config) error {
	context, err := ParseContext(config)
	if err != nil {
		return err
	}

	run(context, defaultCount)

	return nil
}

func run(context *Context, count int) {
	posts := make(chan Post)

	for _, reader := range context.Readers {
		go waitForPosts(reader, posts, count)
	}

	for i := 0; i < count; i++ {
		post := <-posts
		processNewPost(context, post)
	}
}

func waitForPosts(reader FeedReader, posts chan<- Post, count int) {
	log.Println("listening on feed:", reader.GetFeed().Id)
	for i := 0; i < count; i++ {
		for _, post := range reader.FetchNewPosts() {
			posts <- post
		}
		time.Sleep(rssPollingMillis * time.Millisecond)
	}
}

func processNewPost(context *Context, post Post) {
	for _, listener := range context.Listeners {
		listener.OnPost(post)
	}
}
