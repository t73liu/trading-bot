package newsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	ID          string `json:"id"`
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

type ArticlesQueryParams struct {
	Query    string
	Sources  []string
	PageSize int
	Page     int
	Country  string
	Category string
}

type SortArticlesBy string

const (
	Relevancy   SortArticlesBy = "relevancy"
	Popularity  SortArticlesBy = "popularity"
	PublishedAt SortArticlesBy = "publishedAt"
)

type AllArticlesQueryParams struct {
	Query     string
	Sources   []string
	Domains   []string
	PageSize  int
	Page      int
	StartTime time.Time
	EndTime   time.Time
	Language  string
	SortBy    SortArticlesBy
}

func (c *Client) GetTopHeadlinesWithSources(params ArticlesQueryParams) (result ArticlesResponse, err error) {
	req, err := http.NewRequest("GET", newsAPIHost+"/v2/top-headlines", nil)
	if err != nil {
		return result, err
	}

	c.setHeaders(req)
	queryParams := req.URL.Query()
	if params.Query != "" {
		queryParams.Add("q", params.Query)
	}
	if len(params.Sources) > 0 {
		queryParams.Add("sources", strings.Join(params.Sources, ","))
	}
	if params.PageSize > 0 {
		queryParams.Add("pageSize", strconv.Itoa(params.PageSize))
	}
	if params.Page > 1 {
		queryParams.Add("page", strconv.Itoa(params.Page))
	}
	req.URL.RawQuery = queryParams.Encode()

	return c.getArticlesResponse(req)
}

func (c *Client) GetAllHeadlinesWithSources(params AllArticlesQueryParams) (result ArticlesResponse, err error) {
	req, err := http.NewRequest("GET", newsAPIHost+"/v2/everything", nil)
	if err != nil {
		return result, err
	}

	c.setHeaders(req)
	queryParams := req.URL.Query()
	if params.Query != "" {
		queryParams.Add("q", params.Query)
	}
	if len(params.Sources) > 0 {
		queryParams.Add("sources", strings.Join(params.Sources, ","))
	}
	if len(params.Domains) > 0 {
		queryParams.Add("domains", strings.Join(params.Domains, ","))
	}
	if params.PageSize > 0 {
		queryParams.Add("pageSize", strconv.Itoa(params.PageSize))
	}
	if params.Page > 1 {
		queryParams.Add("page", strconv.Itoa(params.Page))
	}
	if params.SortBy != "" {
		queryParams.Add("sortBy", string(params.SortBy))
	}
	if !params.StartTime.IsZero() {
		queryParams.Add("from", formatTime(params.StartTime))
	}
	if !params.EndTime.IsZero() {
		queryParams.Add("to", formatTime(params.EndTime))
	}
	req.URL.RawQuery = queryParams.Encode()

	return c.getArticlesResponse(req)
}

func (c *Client) getArticlesResponse(req *http.Request) (result ArticlesResponse, err error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		var requestError RequestError
		if err := json.NewDecoder(resp.Body).Decode(&requestError); err != nil {
			return result, err
		}
		return result, &requestError
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (c *Client) GetSources(category, language, country string) (sources []Source, err error) {
	req, err := http.NewRequest("GET", newsAPIHost+"/v2/sources", nil)
	if err != nil {
		return sources, err
	}
	c.setHeaders(req)
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
		return sources, err
	}
	if resp.StatusCode != http.StatusOK {
		return sources, errors.New("Response failed with status code: " + resp.Status)
	}
	var result SourcesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return sources, err
	}
	return result.Sources, nil
}

func (c *Client) setHeaders(request *http.Request) {
	request.Header.Set("X-Api-Key", c.apiKey)
}

func formatTime(time time.Time) string {
	return time.Format("2006-01-02T15:04:05")
}
