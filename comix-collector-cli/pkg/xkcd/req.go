package xkcd

import (
	"context"
	"net/http"
)

// buildGetRequest is a helper function to build request
func buildGetRequest(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return req, nil
}
