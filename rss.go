package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title         string        `xml:"title"`
		Link          string        `xml:"link"`
		Description   string        `xml:"description"`
		Category      []RSSCategory `xml:"category"`
		LastBuildDate string        `xml:"lastBuildDate"`
		Item          []RSSItem     `xml:"item"`
	} `xml:"channel"`
}

type RSSCategory struct {
	Category string `xml:"category"`
}

type RSSItem struct {
	Title       string      `xml:"title"`
	Link        string      `xml:"link"`
	Description string      `xml:"description"`
	PubDate     string      `xml:"pubDate"`
	Guid        string      `xml:"guid"`
	Category    RSSCategory `xml:"category"`
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

	// Preprocess XML data to replace &nbsp; with a space
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		log.Println("Reading problem")
		return RSSFeed{}, err
	}
	xmlData := buf.Bytes()
	xmlData = bytes.Replace(xmlData, []byte("&nbsp;"), []byte(" "), -1)
	rssFeed := RSSFeed{}

	// Unmarshal preprocessed XML data
	rssFeed = RSSFeed{}
	if err := xml.Unmarshal(xmlData, &rssFeed); err != nil {
		log.Println("Unmarshal problem")
		return RSSFeed{}, err
	}
	return rssFeed, nil

}
