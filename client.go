package playrlo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	DefaultHost      = "https://play.rust-lang.org"
	DefaultUserAgent = "playrlo/0.0.0"
)

type Client struct {
	host       string
	userAgent  string
	httpClient *http.Client
}

type Config struct {
	Host       string
	UserAgent  string
	HTTPClient *http.Client
}

func NewClient(config *Config) *Client {
	if config == nil {
		config = new(Config)
	}

	host := config.Host
	if host == "" {
		host = DefaultHost
	}

	userAgent := config.UserAgent
	if userAgent == "" {
		userAgent = DefaultUserAgent
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		host:       host,
		userAgent:  userAgent,
		httpClient: httpClient,
	}
}

func (c *Client) Compile(params CompileRequest) (CompileResponse, error) {
	raw, err := clientPost[CompileRequest, rawCompileResponse](c, "/compile", params)
	if err != nil {
		return CompileResponse{}, err
	}

	if raw.Error != "" {
		return CompileResponse{}, fmt.Errorf("compile error: %v", raw.Error)
	}

	return CompileResponse{
		Success: raw.Success,
		Code:    raw.Code,
		Stdout:  raw.Stdout,
		Stderr:  raw.Stderr,
	}, nil
}

func (c *Client) Format(params FormatRequest) (FormatResponse, error) {
	raw, err := clientPost[FormatRequest, rawFormatResponse](c, "/format", params)
	if err != nil {
		return FormatResponse{}, err
	}

	if raw.Error != "" {
		return FormatResponse{}, fmt.Errorf("format error: %v", raw.Error)
	}

	return FormatResponse{
		Success: raw.Success,
		Code:    raw.Code,
		Stdout:  raw.Stdout,
		Stderr:  raw.Stderr,
	}, nil
}

func clientDo[T any](c *Client, req *http.Request) (T, error) {
	var zero T

	req.Header.Add("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return zero, fmt.Errorf("response status: %v", resp.Status)
	}

	var out T
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return zero, err
	}

	return out, nil
}

func clientPost[P, T any](c *Client, path string, params P) (T, error) {
	var zero T

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(params); err != nil {
		return zero, err
	}
	fmt.Println(body)

	req, err := http.NewRequest(http.MethodPost, c.host+path, body)
	if err != nil {
		return zero, err
	}
	req.Header.Add("Content-Type", "application/json")

	return clientDo[T](c, req)
}
