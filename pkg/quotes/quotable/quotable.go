package quotable

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "http://api.quotable.io/random"

type QuotableQuoteProvider struct {
	url    string
	client *http.Client
}

func NewQuotableQuoteProvider() *QuotableQuoteProvider {
	return &QuotableQuoteProvider{
		url:    url,
		client: http.DefaultClient,
	}
}

type Response struct {
	Id           string   `json:"_id"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

func (provider *QuotableQuoteProvider) GetData(ctx context.Context) (Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", provider.url, nil)
	if err != nil {
		return Response{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := provider.client.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
