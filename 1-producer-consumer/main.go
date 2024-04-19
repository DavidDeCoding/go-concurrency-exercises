//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, ch chan<- Tweet) error {
	for {
		tweet, err := stream.Next()

		if err != nil {
			close(ch)
			return err
		}
		ch <- *tweet
	}
}

func consumer(tweet_channel <-chan Tweet) {
	for t := range tweet_channel {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	tweet_channel := make(chan Tweet)

	// Producer
	go producer(stream, tweet_channel)

	// Consumer
	consumer(tweet_channel)

	fmt.Printf("Process took %s\n", time.Since(start))
}
