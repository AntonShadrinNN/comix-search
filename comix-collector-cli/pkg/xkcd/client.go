package xkcd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AntonShadrinNN/comix-search/comix-collector-cli/internal/entities"
)

const (
	jsonEndpoint = "info.0.json"
)

type Client struct {
	ctx        context.Context
	url        string
	httpClient *http.Client
}

func NewClient(ctx context.Context, url string) *Client {
	return &Client{
		ctx:        ctx,
		url:        url,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetComixesCount() (int, error) {
	archiveUrl := fmt.Sprintf("%s/%s", c.url, jsonEndpoint)
	req, err := buildGetRequest(c.ctx, archiveUrl)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var comix entities.Comix
	err = json.Unmarshal(jsonBody, &comix)
	if err != nil {
		return 0, err
	}

	return comix.Id, nil
}

func (c *Client) FetchComixById(id int) (entities.Comix, error) {
	comixUrl := fmt.Sprintf("%s/%d/%s", c.url, id, jsonEndpoint)
	req, err := buildGetRequest(c.ctx, comixUrl)
	if err != nil {
		return entities.Comix{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return entities.Comix{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return entities.Comix{}, ErrNotFound
	}
	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.Comix{}, err
	}

	var comix entities.Comix
	err = json.Unmarshal(jsonBody, &comix)
	if err != nil {
		return entities.Comix{}, err
	}

	return comix, nil
}
