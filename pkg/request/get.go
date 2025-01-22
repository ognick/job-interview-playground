package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func Get[T any](ctx context.Context, url string) (T, error) {
	var response T

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return response, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
