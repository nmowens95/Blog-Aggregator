package main

import (
	"context"
	"log"
	"sync"
	"time"

	"http.github.com/nmowens95/Blog-Aggregator/internal/database"
)

// Takes in connection to database
// Number of concurrent routines
// Time we want between requests per scrape
func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	// Using ;; so that it will run immediately the first time rather than waiting
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Could not fetch feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			// For every feed we add 1 to the wait group
			wg.Add(1)

			// Will spawn a go routine for each
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	// Done decrements counter by 1
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Unable to mark feed", err)
		return
	}

	rssFeed, err := fetchFeed(feed.Url)
	if err != nil {
		log.Println("Could not fetch feed", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
