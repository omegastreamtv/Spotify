package spotifyauth

import (
	"context"
	"errors"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

const (
	// AuthURL is the URL to Spotify Accounts Service's OAuth2 endpoint.
	AuthURL = "https://accounts.spotify.com/authorize"
	// TokenURL is the URL to the Spotify Accounts Service's OAuth2
	// token endpoint.
	TokenURL = "https://accounts.spotify.com/api/token"
)

// Scopes let you specify exactly which types of data your application wants to access.
// The set of scopes you pass in your authentication request determines what access the
// permissions the user is asked to grant.
const (
	// ScopeImageUpload seeks permission to upload images to Spotify on your behalf.
	ScopeImageUpload = "ugc-image-upload"
	// ScopePlaylistReadPrivate seeks permission to read
	// a user's private playlists.
	ScopePlaylistReadPrivate = "playlist-read-private"
	// ScopePlaylistModifyPublic seeks write access
	// to a user's public playlists.
	ScopePlaylistModifyPublic = "playlist-modify-public"
	// ScopePlaylistModifyPrivate seeks write access to
	// a user's private playlists.
	ScopePlaylistModifyPrivate = "playlist-modify-private"
	// ScopePlaylistReadCollaborative seeks permission to
	// access a user's collaborative playlists.
	ScopePlaylistReadCollaborative = "playlist-read-collaborative"
	// ScopeUserFollowModify seeks write/delete access to
	// the list of artists and other users that a user follows.
	ScopeUserFollowModify = "user-follow-modify"
	// ScopeUserFollowRead seeks read access to the list of
	// artists and other users that a user follows.
	ScopeUserFollowRead = "user-follow-read"
	// ScopeUserLibraryModify seeks write/delete access to a
	// user's "Your Music" library.
	ScopeUserLibraryModify = "user-library-modify"
	// ScopeUserLibraryRead seeks read access to a user's "Your Music" library.
	ScopeUserLibraryRead = "user-library-read"
	// ScopeUserReadPrivate seeks read access to a user's
	// subscription details (type of user account).
	ScopeUserReadPrivate = "user-read-private"
	// ScopeUserReadEmail seeks read access to a user's email address.
	ScopeUserReadEmail = "user-read-email"
	// ScopeUserReadCurrentlyPlaying seeks read access to a user's currently playing track
	ScopeUserReadCurrentlyPlaying = "user-read-currently-playing"
	// ScopeUserReadPlaybackState seeks read access to the user's current playback state
	ScopeUserReadPlaybackState = "user-read-playback-state"
	// ScopeUserModifyPlaybackState seeks write access to the user's current playback state
	ScopeUserModifyPlaybackState = "user-modify-playback-state"
	// ScopeUserReadRecentlyPlayed allows access to a user's recently-played songs
	ScopeUserReadRecentlyPlayed = "user-read-recently-played"
	// ScopeUserTopRead seeks read access to a user's top tracks and artists
	ScopeUserTopRead = "user-top-read"
	// ScopeStreaming seeks permission to play music and control playback on your other devices.
	ScopeStreaming = "streaming"
)

type Authenticator struct {
	config *oauth2.Config
}

type AuthenticatorOption func(a *Authenticator)

func WithClientID(id string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientID = id
	}
}

func WithClientSecret(secret string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.ClientSecret = secret
	}
}

// WithScopes configures the oauth scopes that the client should request.
func WithScopes(scopes ...string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.Scopes = scopes
	}
}

func WithRedirectURL(url string) AuthenticatorOption {
	return func(a *Authenticator) {
		a.config.RedirectURL = url
	}
}

func New(opts ...AuthenticatorOption) *Authenticator {
	cfg := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}

	a := &Authenticator{
		config: cfg,
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

// ShowDialog forces the user to approve the app, even if they have already done so.
// Without this, users who have already approved the app are immediately redirected to the redirect uri.
var ShowDialog = oauth2.SetAuthURLParam("show_dialog", "true")

func (a Authenticator) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return a.config.AuthCodeURL(state, opts...)
}

func (a Authenticator) Token(ctx context.Context, state string, r *http.Request, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return nil, errors.New("spotify: auth failed - " + e)
	}
	code := values.Get("code")
	if code == "" {
		return nil, errors.New("spotify: didn't get access code")
	}
	actualState := values.Get("state")
	if actualState != state {
		return nil, errors.New("spotify: redirect state parameter doesn't match")
	}
	return a.config.Exchange(ctx, code, opts...)
}

// Exchange is like Token, except it allows you to manually specify the access
// code instead of pulling it out of an HTTP request.
func (a Authenticator) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return a.config.Exchange(ctx, code, opts...)
}

// Client creates a *http.Client that will use the specified access token for its API requests.
// Combine this with spotify.HTTPClientOpt.
func (a Authenticator) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return a.config.Client(ctx, token)
}
