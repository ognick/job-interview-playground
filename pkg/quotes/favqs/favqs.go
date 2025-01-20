package favqs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const url = "https://favqs.com/api/qotd"

type FavqsQuoteProvider struct {
	url    string
	client *http.Client
}

func NewFavqsQuoteProvider() *FavqsQuoteProvider {
	return &FavqsQuoteProvider{
		url:    url,
		client: http.DefaultClient,
	}
}

type Response struct {
	QotdDate time.Time `json:"qotd_date"`
	Quote    struct {
		Id              int      `json:"id"`
		Dialogue        bool     `json:"dialogue"`
		Private         bool     `json:"private"`
		Tags            []string `json:"tags"`
		Url             string   `json:"url"`
		FavoritesCount  int      `json:"favorites_count"`
		UpvotesCount    int      `json:"upvotes_count"`
		DownvotesCount  int      `json:"downvotes_count"`
		Author          string   `json:"author"`
		AuthorPermalink string   `json:"author_permalink"`
		Body            string   `json:"body"`
	} `json:"quote"`
}

func (provider *FavqsQuoteProvider) GetData(ctx context.Context) (Response, error) {
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
