package twitter

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Listen(filterString string, callback func([]byte)) {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter Client
	client := twitter.NewClient(httpClient)

	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		body, _ := json.Marshal(tweet)
		callback(body)
	}
	log.Printf("Starting Stream with `%s` filter", filterString)

	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{filterString},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	// Receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	cha := make(chan os.Signal)
	signal.Notify(cha, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-cha)

	log.Println("Stopping Stream...")
	stream.Stop()
}
