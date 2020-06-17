package newsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const newsAPIHost = "https://newsapi.org"

type Client struct {
	client http.Client
	apiKey string
}

type ArticlesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	UrlToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type SourcesResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

type Source struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type RequestError struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	return &Client{
		client: *httpClient,
		apiKey: apiKey,
	}
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("%s - %s - %s", e.Status, e.Code, e.Message)
}

func (c *Client) GetTopHeadlinesBySources(query, sources string) (*ArticlesResponse, error) {
	return c.getHeadlines("/v2/top-headlines", query, sources)
}

func (c *Client) GetAllHeadlinesBySources(query, sources string) (*ArticlesResponse, error) {
	return c.getHeadlines("/v2/everything", query, sources)
}

func (c *Client) getHeadlines(path, query, sources string) (*ArticlesResponse, error) {
	req, err := http.NewRequest("GET", newsAPIHost+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", c.apiKey)
	queryParams := req.URL.Query()
	if query != "" {
		queryParams.Add("q", query)
	}
	if sources != "" {
		queryParams.Add("sources", sources)
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var requestError RequestError
		if err := json.NewDecoder(resp.Body).Decode(&requestError); err != nil {
			return nil, err
		}
		return nil, &requestError
	}
	var result ArticlesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSources(category, language, country string) ([]Source, error) {
	req, err := http.NewRequest("GET", newsAPIHost+"/v2/sources", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", c.apiKey)
	queryParams := req.URL.Query()
	if category != "" {
		queryParams.Add("category", category)
	}
	if language != "" {
		queryParams.Add("language", language)
	}
	if country != "" {
		queryParams.Add("country", country)
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var requestError RequestError
		if err := json.NewDecoder(resp.Body).Decode(&requestError); err != nil {
			return nil, err
		}
		return nil, &requestError
	}
	var result SourcesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Sources, nil
}
