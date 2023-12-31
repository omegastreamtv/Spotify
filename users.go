package spotify

import (
	"fmt"
	"strings"
)

type User struct {
	// Known external URLs for this user.
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	// Information about the followers of the user.
	Followers Followers `json:"followers"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The Spotify user ID for the user.
	ID string `json:"id"`
	// The object type: "user"
	Type string `json:"type"`
	// The Spotify URI for the user.
	URI string `json:"uri"`
}

type BasicUser struct {
	User
	// The name displayed on the user's profile. null if not available.
	DisplayName string `json:"display_name"`
	// The user's profile image.
	Images []Image `json:"images"`
}

type FullUser struct {
	BasicUser
	// The country of the user, as set in the user's account profile. An ISO 3166-1 alpha-2 country code. This field is only available when the current user has granted access to the user-read-private scope.
	Country string `json:"country"`
	// The user's email address, as entered by the user when creating their account. Important! This email address is unverified; there is no proof that it actually belongs to the user. This field is only available when the current user has granted access to the user-read-email scope.
	Email string `json:"email"`
	// The user's explicit content settings. This field is only available when the current user has granted access to the user-read-private scope.
	ExplicitContent struct {
		// When true, indicates that explicit content should not be played.
		FilterEnabled bool `json:"filter_enabled"`
		// When true, indicates that the explicit content setting is locked and can't be changed by the user.
		FilterLocked bool `json:"filter_locked"`
	} `json:"explicit_content"`
	// The user's Spotify subscription level: "premium", "free", etc. (The subscription level "open" can be considered the same as "free".) This field is only available when the current user has granted access to the user-read-private scope.
	Product string `json:"product"`
}

type GetCurrentUsersProfileResponse struct {
	FullUser
}

// Get detailed profile information about the current user (including the current user's username).
//
// Required scope: user-read-private, user-read-email
func (c *Client) GetCurrentUsersProfile() (*GetCurrentUsersProfileResponse, error) {
	profile := GetCurrentUsersProfileResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get("/me").Receive(&profile, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &profile, nil
}

type GetUsersTopItemsParams struct {
	// Over what time frame the affinities are computed. Valid values: long_term (calculated from several years of data and including all new data as it becomes available), medium_term (approximately last 6 months), short_term (approximately last 4 weeks). Default: medium_term
	TimeRange string `json:"time_range"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `json:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `json:"offset"`
}

type GetUsersTopArtistsResponse struct {
	Pagination
	Items []Artist `json:"items"`
}

type GetUsersTopTracksResponse struct {
	Pagination
	Items []Track `json:"items"`
}

// Get the current user's top artists based on calculated affinity.
//
// Required scope: user-top-read
func (c *Client) GetUsersTopArtists(params *GetUsersTopItemsParams) (*GetUsersTopArtistsResponse, error) {
	artists := GetUsersTopArtistsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get("/me/top/artists").QueryStruct(params).Receive(&artists, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &artists, nil
}

// Get the current user's top tracks based on calculated affinity.
//
// Required scope: user-top-read
func (c *Client) GetUsersTopTracks(params *GetUsersTopItemsParams) (*GetUsersTopTracksResponse, error) {
	tracks := GetUsersTopTracksResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get("/me/top/tracks").QueryStruct(params).Receive(&tracks, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &tracks, nil
}

type GetUsersProfileResponse struct {
	User
}

// Get public profile information about a Spotify user.
func (c *Client) GetUsersProfile(userId string) (*GetUsersProfileResponse, error) {
	profile := GetUsersProfileResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get(fmt.Sprintf("/users/%s", userId)).Receive(&profile, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &profile, nil
}

type FollowPlaylistBody struct {
	// Defaults to true. If true the playlist will be included in user's public playlists, if false it will remain private.
	Public bool `json:"public"`
}

// Add the current user as a follower of a playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) FollowPlaylist(playlistId string, payload *FollowPlaylistBody) error {
	var res struct{}
	var spotifyErr *SpotifyError

	_, err := c.put(fmt.Sprintf("/playlists/%s/followers", playlistId)).BodyJSON(payload).Receive(&res, &spotifyErr)
	if err != nil {
		return err
	}

	if spotifyErr != nil {
		return spotifyErr
	}

	return nil
}

// Remove the current user as a follower of a playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) UnfollowPlaylist(playlistId string) error {
	var res struct{}
	var spotifyErr *SpotifyError

	_, err := c.delete(fmt.Sprintf("/playlists/%s/followers", playlistId)).Receive(&res, &spotifyErr)
	if err != nil {
		return err
	}

	if spotifyErr != nil {
		return spotifyErr
	}

	return nil
}

type GetFollowedArtistsParams struct {
	// The ID type: currently only artist is supported.
	Type string `url:"type"`
	// The last artist ID retrieved from the previous request.
	After string `url:"after"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
}

type GetFollowedArtistsResponse struct {
	Artists struct {
		Pagination
		Items []Artist `json:"items"`
	} `json:"artists"`
}

// Get the current user's followed artists.
//
// Required scope: user-follow-read
func (c *Client) GetFollowedArtists(params *GetFollowedArtistsParams) (*GetFollowedArtistsResponse, error) {
	artists := GetFollowedArtistsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get("/me/following").Receive(&artists, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &artists, nil
}

type FollowArtistsOrUsersParams struct {
	// The ID type.
	Type string `url:"type"`
	// A comma-separated list of the artist or the user Spotify IDs. A maximum of 50 IDs can be sent in one request.
	IDs string `url:"ids"`
}

// Add the current user as a follower of one or more artists or other Spotify users.
//
// Required scope: user-follow-modify
func (c *Client) FollowArtistsOrUsers(typ string, artistIds []string) error {
	var res struct{}
	var spotifyErr *SpotifyError

	params := FollowArtistsOrUsersParams{
		Type: typ,
		IDs:  strings.Join(artistIds, ","),
	}

	_, err := c.put("/me/following").QueryStruct(params).Receive(&res, &spotifyErr)
	if err != nil {
		return err
	}

	if spotifyErr != nil {
		return spotifyErr
	}

	return nil
}

type UnfollowArtistsOrUsersParams struct {
	// The ID type: either artist or user.
	Type string `url:"type"`
	// A comma-separated list of the artist or the user Spotify IDs. A maximum of 50 IDs can be sent in one request.
	IDs string `url:"ids"`
}

// Remove the current user as a follower of one or more artists or other Spotify users.
//
// Required scope: user-follow-modify
func (c *Client) UnfollowArtistsOrUsers(typ string, artistIds []string) error {
	var res struct{}
	var spotifyErr *SpotifyError

	params := UnfollowArtistsOrUsersParams{
		Type: typ,
		IDs:  strings.Join(artistIds, ","),
	}

	_, err := c.delete("/me/following").QueryStruct(params).Receive(&res, &spotifyErr)
	if err != nil {
		return err
	}

	if spotifyErr != nil {
		return spotifyErr
	}

	return nil
}

type CheckIfUserFollowsArtistsOrUsersParams struct {
	// The ID type: either artist or user.
	Type string `url:"type"`
	// A comma-separated list of the artist or the user Spotify IDs to check. For example: ids=74ASZWbe4lXaubB36ztrGX,08td7MxkoHQkXnWAYD8d6Q. A maximum of 50 IDs can be sent in one request.
	IDs string `url:"ids"`
}

type CheckIfUserFollowsArtistsOrUsersResponse []bool

// Check to see if the current user is following one or more artists or other Spotify users.
//
//	Required scope: user-follow-read
func (c *Client) CheckIfUserFollowsArtistsOrUsers(typ string, artistIds []string) (*CheckIfUserFollowsArtistsOrUsersResponse, error) {
	resEach := CheckIfUserFollowsArtistsOrUsersResponse{}
	var spotifyErr *SpotifyError

	params := CheckIfUserFollowsArtistsOrUsersParams{
		Type: typ,
		IDs:  strings.Join(artistIds, ","),
	}

	_, err := c.get("/me/following/contains").QueryStruct(params).Receive(&resEach, &spotifyErr)
	if err != nil {
		return nil, err
	}

	return &resEach, nil
}

type CheckIfUsersFollowPlaylistParams struct {
	// A comma-separated list of Spotify User IDs ; the ids of the users that you want to check to see if they follow the playlist. Maximum: 5 ids.
	IDs string `url:"ids"`
}

type CheckIfUsersFollowPlaylistResponse []bool

// Check to see if one or more Spotify users are following a specified playlist.
func (c *Client) CheckIfUsersFollowPlaylist(playlistId string, userIds []string) (*CheckIfUsersFollowPlaylistResponse, error) {
	resEach := CheckIfUsersFollowPlaylistResponse{}
	var spotifyErr *SpotifyError

	params := CheckIfUsersFollowPlaylistParams{
		IDs: strings.Join(userIds, ","),
	}

	_, err := c.get(fmt.Sprintf("/playlists/%s/followers/contains", playlistId)).QueryStruct(params).Receive(&resEach, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &resEach, nil
}
