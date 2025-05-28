package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port                string
	DatabaseURL         string
	ScrapingConcurrency int
	ScrapingInterval    time.Duration
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT environment variable is required")
	}

	databaseURL := os.Getenv("DB_URL")
	if databaseURL == "" {
		return nil, errors.New("DB_URL environment variable is required")
	}

	scrapingConcurrency := 5
	scrapingInterval := time.Minute

	if concurrencyStr := os.Getenv("SCRAPING_CONCURRENCY"); concurrencyStr != "" {
		if concurrency, err := strconv.Atoi(concurrencyStr); err == nil {
			scrapingConcurrency = concurrency
		}
	}

	if intervalStr := os.Getenv("SCRAPING_INTERVAL"); intervalStr != "" {
		if interval, err := time.ParseDuration(intervalStr); err == nil {
			scrapingInterval = interval
		}
	}

	return &Config{
		Port:                port,
		DatabaseURL:         databaseURL,
		ScrapingConcurrency: scrapingConcurrency,
		ScrapingInterval:    scrapingInterval,
	}, nil
}
