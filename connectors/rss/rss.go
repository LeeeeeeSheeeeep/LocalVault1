package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"localvault/connectors"
	"localvault/storage"
)

type RSSConnector struct {
	feedURL string
}

func New() *RSSConnector {
	return &RSSConnector{}
}

func (c *RSSConnector) ID() string {
	return "rss"
}

func (c *RSSConnector) Name() string {
	return "RSS Feed"
}

func (c *RSSConnector) Auth(ctx context.Context, config map[string]string) error {
	feedURL, ok := config["url"]
	if !ok || feedURL == "" {
		return fmt.Errorf("RSS feed url is required")
	}
	c.feedURL = feedURL
	return nil
}

type RSSFeed struct {
	Channel struct {
		Title string    `xml:"title"`
		Items []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func (c *RSSConnector) Sync(ctx context.Context, state connectors.SyncState, docChan chan<- *storage.Document) (connectors.SyncState, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.feedURL, nil)
	if err != nil {
		return state, err
	}

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return state, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return state, fmt.Errorf("failed to fetch RSS feed, status: %d", resp.StatusCode)
	}

	var feed RSSFeed
	if err := xml.NewDecoder(resp.Body).Decode(&feed); err != nil {
		return state, err
	}

	for _, item := range feed.Channel.Items {
		pubDate, _ := time.Parse(time.RFC1123Z, item.PubDate)
		
		if !state.LastSync.IsZero() && pubDate.Before(state.LastSync) {
			continue
		}

		rawMap := map[string]interface{}{
			"guid":    item.GUID,
			"pubDate": item.PubDate,
		}

		doc := &storage.Document{
			ProviderID: c.ID(),
			SourceID:   item.GUID,
			DocType:    "article",
			Title:      item.Title,
			Content:    item.Description,
			RawData:    rawMap,
			URL:        item.Link,
			Author:     feed.Channel.Title,
			CreatedAt:  pubDate,
			UpdatedAt:  pubDate,
		}

		select {
		case <-ctx.Done():
			return state, ctx.Err()
		case docChan <- doc:
		}
	}

	state.LastSync = time.Now()
	return state, nil
}
