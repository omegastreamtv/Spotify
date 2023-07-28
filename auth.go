package spotify

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	// AuthURL is the URL to Spotify Accounts Service's OAuth2 endpoint.
	AuthURL = "https://accounts.spotify.com/authorize"
	// TokenURL is the URL to the Spotify Accounts Service's OAuth2 token endpoint.
	TokenURL = "https://accounts.spotify.com/api/token"
)

// Scopes provide Spotify users using third-party apps the confidence that only the information they choose to share will be shared, and nothing more.
type Scope string

const (
	// Write access to user-provided images.
	//
	// Visible to users: Upload images to Spotify on your behalf.
	ScopeUGCImageUpload Scope = "ugc-image-upload"
	// Read access to a user’s player state.
	//
	// Visible to users: Read your currently playing content and Spotify Connect devices information.
	ScopeUserReadPlaybackState Scope = "user-read-playback-state"
	// Write access to a user’s playback state
	//
	//Visible to users: Control playback on your Spotify clients and Spotify Connect devices.
	ScopeUserModifyPlaybackState Scope = "user-modify-playback-state"
	// Read access to a user’s currently playing content.
	//
	// Visible to users: Read your currently playing content.
	ScopeUserReadCurrentlyPlaying Scope = "user-read-currently-playing"
	// Remote control playback of Spotify. This scope is currently available to Spotify iOS and Android SDKs.
	//
	// Visible to users: Communicate with the Spotify app on your device.
	ScopeAppRemoteControl Scope = "app-remote-control"
	// Control playback of a Spotify track. This scope is currently available to the Web Playback SDK. The user must have a Spotify Premium account.
	//
	// Visible to users: Play content and control playback on your other devices.
	ScopeStreaming Scope = "streaming"
	// Read access to user's private playlists.
	//
	// Visible to users: Access your private playlists.
	ScopePlaylistReadPrivate Scope = "playlist-read-private"
	// Include collaborative playlists when requesting a user's playlists.
	//
	// Visible to users: Access your collaborative playlists.
	ScopePlaylistReadCollaborate Scope = "playlist-read-collaborative"
	// Write access to a user's private playlists.
	//
	// Visible to users: Manage your private playlists.
	ScopePlaylistModifyPrivate Scope = "playlist-modify-private"
	// Write access to a user's public playlists.
	//
	// Visible to users: Manage your public playlists.
	ScopePlaylistModifyPublic Scope = "playlist-modify-public"
	// Write/delete access to the list of artists and other users that the user follows.
	//
	// Visible to users: Manage who you are following.
	ScopeUserFollowModify Scope = "user-follow-modify"
	// Read access to a user’s playback position in a content.
	//
	// Visible to users: Read your position in content you have played.
	ScopeUserFollowRead Scope = "user-follow-read"
	// Read access to a user's top artists and tracks.
	//
	// Visible to users: Read your top artists and content.
	ScopeUserTopRead Scope = "user-top-read"
	// Read access to a user’s recently played tracks.
	//
	// Visible to users: Access your recently played items.
	ScopeUserReadRecentlyPlayed Scope = "user-read-recently-played"
	// Write/delete access to a user's "Your Music" library.
	//
	// Visible to users: Manage your saved content.
	ScopeUserLibraryModify Scope = "user-library-modify"
	// Read access to a user's library.
	//
	// Visible to users: Access your saved content.
	ScopeUserLibraryRead Scope = "user-library-read"
	// Read access to user’s email address.
	//
	// Visible to users: Get your real email address.
	ScopeUserReadEmail Scope = "user-read-email"
	// Read access to user’s subscription details (type of user account).
	//
	// Visible to users: Access your subscription details.
	ScopeUserReadPrivate Scope = "user-read-private"
)

type Authenticator struct {
	config *oauth2.Config
}

type AuthenticatorOption func(a *Authenticator)

// WithScopes configures the OAuth scopes that the client should request.
func WithScopes(scopes ...Scope) AuthenticatorOption {
	var strScopes []string

	for _, scope := range scopes {
		strScopes = append(strScopes, string(scope))
	}

	return func(a *Authenticator) {
		a.config.Scopes = strScopes
	}
}

// NewAuthenticator returns a new Authenticator.
func (c *Client) NewAuthenticator(opts ...AuthenticatorOption) *Authenticator {
	auth := &Authenticator{
		config: &oauth2.Config{
			ClientID:     c.clientID,
			ClientSecret: c.clientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:  AuthURL,
				TokenURL: TokenURL,
			},
			RedirectURL: c.redirectURI,
			Scopes:      []string{},
		},
	}

	for _, opt := range opts {
		opt(auth)
	}

	return auth
}

// ShowDialog will prompt the user to approve the app again if they’ve already done so.
var ShowDialog = oauth2.SetAuthURLParam("show_dialog", "true")

// AuthURL returns a URL to OAuth 2.0 provider's consent page
func (a Authenticator) AuthURL(state string, opts ...oauth2.AuthCodeOption) string {
	return a.config.AuthCodeURL(state, opts...)
}

// Token exchanges the OAuth 2.0 authorization code for a token.
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
