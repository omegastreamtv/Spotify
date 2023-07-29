package spotify

import (
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/dghubble/sling"
)

const URL = "https://api.spotify.com/v1"

type Client struct {
	// mu is used to lock the client for concurrent requests
	mu sync.RWMutex
	// userAcccessToken is used to set the user access token for the current request when using WithUserAccessToken
	userAccessToken string
	// appAccessToken is used when the userAccessToken is not provided
	appAccessToken string
	// baseURL is the URL to use when making requests
	baseURL string
	// Custom HTTP client
	httpClient *http.Client
	// The Client ID generated after registering your application.
	clientID string
	// The Client secret generated after registering your application.
	clientSecret string
	// The URI to redirect to after the user grants or denies permission.
	redirectURI string
}

type ClientOption func(*Client)

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		if baseURL == "" {
			c.baseURL = URL
			return
		}

		c.baseURL = baseURL
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		if httpClient == nil {
			c.httpClient = http.DefaultClient
			return
		}

		c.httpClient = httpClient
	}
}

// WithClientID Provides The client the Client ID generated after registering your application.
func WithClientID(clientID string) ClientOption {
	return func(c *Client) {
		c.clientID = clientID
	}
}

// WithClientSecret Provides the client with a secret generated after registering your application.
func WithClientSecret(clientSecret string) ClientOption {
	return func(c *Client) {
		c.clientSecret = clientSecret
	}
}

// WithRedirectURI Provides the URI to redirect to after the user grants or denies permission. This URI needs to have been entered in the Redirect URI allowlist that you specified when you registered your application (See the app guide). The value of redirect_uri here must exactly match one of the values you entered when you registered your application, including upper or lowercase, terminating slashes, and such.
func WithRedirectURI(redirectURI string) ClientOption {
	return func(c *Client) {
		c.redirectURI = redirectURI
	}
}

func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{
		mu:              sync.RWMutex{},
		userAccessToken: "",
		httpClient:      http.DefaultClient,
		baseURL:         URL,
	}

	for _, option := range options {
		option(c)
	}

	if c.clientID == "" {
		return nil, errors.New("ClientID cannot be empty")
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

// SetAppAccessToken will set the app access token for the client.
func (c *Client) SetAppAccessToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.appAccessToken = token
}

// WithUserAccessToken will set the user access token for the current request.
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
