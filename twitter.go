package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Tweet struct {
	User     string
	Contents string
}

// Start loop looking for  new events.
// Will push tweet struct to channel on new tweets.
func StartTweetListener(tweetQueue chan Tweet) {
	config := oauth1.NewConfig("E8aFRSq6IbcBhJQSBf9GYVgiR", "WPnxdAAXvMcJs161x7wQGhcIKoJNpt8yeoo9hBDz7Ov2zMIaBc")
	token := oauth1.NewToken("717783264-KWFLpQl9grduyMvvpYp8CZoCSCa5dHQjANG1qvy8", "g5ewB7YKTjNWcP7NBczgiuoZSgX3HZ7KpvbiNcBO1yb42")
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	params := &twitter.StreamSampleParams{StallWarnings: twitter.Bool(true), Language: []string{"en"}}
	stream, err := client.Streams.Sample(params)

	if err != nil {
		fmt.Println("Error connecting to twitter")
	}

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		tweetQueueItem := Tweet{Contents: tweet.Text, User: ""}
		tweetQueue <- tweetQueueItem
	}

	demux.HandleChan(stream.Messages)
}
