package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RestClient interface {
	GetRequest(ctx context.Context, url string, header map[string]string, responseBody any) error
	PostRequest(ctx context.Context, url string, header map[string]string, body any, responseBody any) error
	PutRequest(ctx context.Context, url string, header map[string]string, body any, responseBody any) error
	DeleteRequest(ctx context.Context, url string, header map[string]string, body any, responseBody any) error
}
type restClientImpl struct {
	BaseUrl    string
	HttpClient *http.Client
}

func NewRestClient(baseUrl string, timeout time.Duration) RestClient {
	return &restClientImpl{
		BaseUrl:    baseUrl,
		HttpClient: &http.Client{Timeout: timeout},
	}
}

func (c *restClientImpl) GetRequest(ctx context.Context, url string, header map[string]string, responseBody any) error {

	request, err := http.NewRequestWithContext(ctx, "GET", c.BaseUrl+url, nil)
	if err != nil {
		return err
	}
	// Header setup
	for k, v := range header {
		request.Header.Add(k, v)
	}

	// Make the request
	resp, err := c.HttpClient.Do(request)
	// Check for errors
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for non-200 status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("HTTP Error(%d) : %s", resp.StatusCode, resp.Status))
	}

	// Unmarshal the response body into the provided responseBody variable
	if err = json.Unmarshal(body, responseBody); err != nil {
		return err
	}

	return nil
}

func (c *restClientImpl) PostRequest(ctx context.Context, url string, header map[string]string, body any, responseBody any) error {
	//TODO implement me
	panic("implement me")
}

func (c *restClientImpl) PutRequest(ctx context.Context, url string, header map[string]string, body any, responseBody any) error {
	//TODO implement me
	panic("implement me")
}

func (c *restClientImpl) DeleteRequest(ctx context.Context, url string, header map[string]string, body any, responseBody any) error {
	//TODO implement me
	panic("implement me")
}
