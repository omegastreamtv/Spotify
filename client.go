package spotify

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/dghubble/sling"
)

const (
	URL = "https://api.spotify.com/v1"
)

type Client struct {
	// mu is used to lock the client for concurrent requests
	mu sync.RWMutex
	// opts are the options used to create the client
	opts *Options
	// userAcccessToken is used to set the user access token for the current request when using WithUserAccessToken
	userAccessToken string
	// appAccessToken is used when the userAccessToken is not provided
	appAccessToken string
	// baseURL is the URL to use when making requests
	baseURL string
	// Custom HTTP client
	httpClient *http.Client
}

type Options struct {
	// The Client ID generated after registering your application.
	ClientID string
	// The Client secret generated after registering your application.
	ClientSecret string
	// The URI to redirect to after the user grants or denies permission. This URI needs to have been entered in the Redirect URI allowlist that you specified when you registered your application (See the app guide). The value of redirect_uri here must exactly match one of the values you entered when you registered your application, including upper or lowercase, terminating slashes, and such.
	RedirectURI string
	// BaseURL is the URL to use when making requests
	BaseURL string
}

func NewClient(options *Options, client *http.Client) (*Client, error) {
	c := &Client{
		mu:              sync.RWMutex{},
		userAccessToken: "",
		httpClient:      http.DefaultClient,
		baseURL:         URL,
	}

	if options.BaseURL != "" {
		c.baseURL = options.BaseURL
	}

	if options.ClientID == "" {
		return nil, errors.New("ClientID cannot be empty")
	}

	c.opts = options

	if client == nil {
		c.httpClient = client
	}

	return c, nil
}

type SpotifyError struct {
	Err struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
}

func (s *SpotifyError) Error() string {
	return fmt.Sprintf("Spotify API Error - Status: %d, Message: %s", s.Err.Status, s.Err.Message)
}

type Pagination struct {
	// A link to the Web API endpoint returning the full result of the request
	Href string `json:"href"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int `json:"limit"`
	// URL to the next page of items. (null if none)
	Next string `json:"next"`
	// The offset of the items returned (as set in the query or by default)
	Offset int `json:"offset"`
	// URL to the previous page of items. (null if none)
	Previous string `json:"previous"`
	// The total number of items available to return.
	Total int `json:"total"`
}

func (c *Client) SetAppAccessToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.appAccessToken = token
}

func (c *Client) WithUserAccessToken(userAccessToken string) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.userAccessToken = userAccessToken

	return c
}

func (c *Client) get(path string) *sling.Sling {
	c.mu.Lock()

	tokenToUse := c.appAccessToken
	if c.userAccessToken != "" {
		tokenToUse = c.userAccessToken
	}

	req := sling.New().Get(c.baseURL+path).Set("Authorization", "Bearer "+tokenToUse)
	c.userAccessToken = ""
	c.mu.Unlock()

	return req
}

func (c *Client) post(path string) *sling.Sling {
	c.mu.Lock()

	tokenToUse := c.appAccessToken
	if c.userAccessToken != "" {
		tokenToUse = c.userAccessToken
	}

	req := sling.New().Post(c.baseURL+path).Set("Authorization", "Bearer "+tokenToUse)
	c.userAccessToken = ""
	c.mu.Unlock()

	return req
}

func (c *Client) put(path string) *sling.Sling {
	c.mu.Lock()

	tokenToUse := c.appAccessToken
	if c.userAccessToken != "" {
		tokenToUse = c.userAccessToken
	}

	req := sling.New().Put(c.baseURL+path).Set("Authorization", "Bearer "+tokenToUse)
	c.userAccessToken = ""
	c.mu.Unlock()

	return req
}

func (c *Client) delete(path string) *sling.Sling {
	c.mu.Lock()

	tokenToUse := c.appAccessToken
	if c.userAccessToken != "" {
		tokenToUse = c.userAccessToken
	}

	req := sling.New().Delete(c.baseURL+path).Set("Authorization", "Bearer "+tokenToUse)
	c.userAccessToken = ""
	c.mu.Unlock()

	return req
}
