package services

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
)

type ScraperService struct {
	db                  *database.Queries
	concurrency         int
	timeBetweenRequests time.Duration
	httpClient          *http.Client
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func NewScraperService(db *database.Queries, concurrency int, timeBetweenRequests time.Duration) *ScraperService {
	return &ScraperService{
		db:                  db,
		concurrency:         concurrency,
		timeBetweenRequests: timeBetweenRequests,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (s *ScraperService) Start() {
	log.Printf("Starting RSS scraper with concurrency %d and interval %s", s.concurrency, s.timeBetweenRequests)

	ticker := time.NewTicker(s.timeBetweenRequests)
	defer ticker.Stop()

	for range ticker.C {
		s.scrapeFeeds()
	}
}

func (s *ScraperService) scrapeFeeds() {
	ctx := context.Background()

	feeds, err := s.db.GetNextFeedsToFetch(ctx, int32(s.concurrency))
	if err != nil {
		log.Printf("Error fetching feeds to scrape: %v", err)
		return
	}

	if len(feeds) == 0 {
		return
	}

	log.Printf("Found %d feeds to scrape", len(feeds))

	wg := &sync.WaitGroup{}
	for _, feed := range feeds {
		wg.Add(1)
		go s.scrapeFeed(wg, feed)
	}
	wg.Wait()
}

func (s *ScraperService) scrapeFeed(wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	ctx := context.Background()

	_, err := s.db.MarkFeedAsFetched(ctx, feed.ID)
	if err != nil {
		log.Printf("Error marking feed %s as fetched: %v", feed.Name, err)
		return
	}

	rssFeed, err := s.fetchRSSFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching RSS feed %s (%s): %v", feed.Name, feed.Url, err)
		return
	}

	postsCreated := 0
	for _, item := range rssFeed.Channel.Item {
		if s.createPost(ctx, item, feed.ID) {
			postsCreated++
		}
	}

	log.Printf("Feed %s processed: %d new posts created out of %d items",
		feed.Name, postsCreated, len(rssFeed.Channel.Item))
}

func (s *ScraperService) fetchRSSFeed(url string) (RSSFeed, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "RSS-Aggregator/1.0")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RSSFeed{}, fmt.Errorf("HTTP error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("error reading response body: %w", err)
	}

	var rssFeed RSSFeed
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return RSSFeed{}, fmt.Errorf("error parsing XML: %w", err)
	}

	return rssFeed, nil
}

func (s *ScraperService) createPost(ctx context.Context, item RSSItem, feedID uuid.UUID) bool {
	description := sql.NullString{}
	if item.Description != "" {
		description.String = item.Description
		description.Valid = true
	}

	pubDate, err := s.parseDate(item.PubDate)
	if err != nil {
		log.Printf("Error parsing date '%s': %v", item.PubDate, err)
		pubDate = time.Now().UTC()
	}

	_, err = s.db.CreatePost(ctx, database.CreatePostParams{
		ID:          uuid.New(),
		Title:       item.Title,
		Description: description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Url:         item.Link,
		PublishedAt: pubDate,
		FeedID:      feedID,
	})

	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") ||
			strings.Contains(err.Error(), "duplicate key") {
			return false
		}
		log.Printf("Error creating post '%s': %v", item.Title, err)
		return false
	}

	return true
}

func (s *ScraperService) parseDate(dateStr string) (time.Time, error) {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
	}

	for _, format := range formats {
		if parsed, err := time.Parse(format, dateStr); err == nil {
			return parsed.UTC(), nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateStr)
}
