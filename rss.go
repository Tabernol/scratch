package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title         string    `xml:"title"`
		Link          string    `xml:"link"`
		Description   string    `xml:"description"`
		Category      []string  `xml:"category"`
		LastBuildDate string    `xml:"lastBuildDate"`
		Item          []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Category    []string `xml:"category"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Println("Error with responce")
		return RSSFeed{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read problem")
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	// Unmarshal preprocessed XML data
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		log.Println("Unmarshal problem")
		return RSSFeed{}, err
	}
	return rssFeed, nil

}
