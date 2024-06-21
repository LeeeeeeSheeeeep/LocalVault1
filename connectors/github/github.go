package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"localvault/connectors"
	"localvault/storage"
)

type GitHubConnector struct {
	token string
}

func New() *GitHubConnector {
	return &GitHubConnector{}
}

func (c *GitHubConnector) ID() string {
	return "github"
}

func (c *GitHubConnector) Name() string {
	return "GitHub"
}

func (c *GitHubConnector) Auth(ctx context.Context, config map[string]string) error {
	token, ok := config["token"]
	if !ok || token == "" {
		return fmt.Errorf("github token is required")
	}
	c.token = token
	return nil
}

// Minimal struct for GitHub Issue
type githubIssue struct {
	ID        int64     `json:"id"`
	Number    int       `json:"number"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	State     string    `json:"state"`
	HTMLURL   string    `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Login string `json:"login"`
	} `json:"user"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

func (c *GitHubConnector) Sync(ctx context.Context, state connectors.SyncState, docChan chan<- *storage.Document) (connectors.SyncState, error) {
	// Simple implementation: Fetch user's assigned issues
	// A production version would use pagination and handle rate limits.
	url := "https://api.github.com/issues?filter=assigned&state=all&per_page=50"
	
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return state, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return state, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return state, fmt.Errorf("github API error: %s", string(body))
	}

	var issues []githubIssue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return state, err
	}

	for _, issue := range issues {
		// Only sync newer issues if we have a LastSync
		if !state.LastSync.IsZero() && issue.UpdatedAt.Before(state.LastSync) {
			continue
		}

		rawMap := make(map[string]interface{})
		rawBytes, _ := json.Marshal(issue)
		json.Unmarshal(rawBytes, &rawMap)

		doc := &storage.Document{
			ProviderID: c.ID(),
			SourceID:   fmt.Sprintf("%d", issue.ID),
			DocType:    "issue",
			Title:      fmt.Sprintf("[%s] %s", issue.Repository.FullName, issue.Title),
			Content:    issue.Body,
			RawData:    rawMap,
			URL:        issue.HTMLURL,
			Author:     issue.User.Login,
			CreatedAt:  issue.CreatedAt,
			UpdatedAt:  issue.UpdatedAt,
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
