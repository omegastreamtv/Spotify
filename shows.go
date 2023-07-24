package spotify

import (
	"fmt"
	"strings"
)

type Show struct {
	AvailableMarkets   []string     `json:"available_markets"`
	Copyrights         []Copyright  `json:"copyrights"`
	Description        string       `json:"description"`
	HTMLDescription    string       `json:"html_description"`
	Explicit           bool         `json:"explicit"`
	ExternalURLs       ExternalURLs `json:"external_ur_ls"`
	Href               string       `json:"href"`
	ID                 string       `json:"id"`
	Images             []Image      `json:"images"`
	IsExternallyHosted bool         `json:"is_externally_hosted"`
	Languages          []string     `json:"languages"`
	MediaType          string       `json:"media_type"`
	Name               string       `json:"name"`
	Publisher          string       `json:"publisher"`
	Type               string       `json:"type"`
	URI                string       `json:"uri"`
	TotalEpisodes      int          `json:"total_episodes"`
}

type FullShow struct {
	Show
	Episodes struct {
		Pagination
		Items []Episode `json:"items"`
	} `json:"episodes"`
}

type SavedShow struct {
	AddedAt string `json:"added_at"`
	Show
}

type GetShowParams struct {
	Market string `url:"market"`
}

type GetShowResponse struct {
	Show
}

// Get Spotify catalog information for a single show identified by its unique Spotify ID.
//
// Required scope: user-read-playback-position
func (c *Client) GetShow(id string, market string) (*GetShowResponse, error) {
	show := GetShowResponse{}
	var err *SpotifyError

	params := GetShowParams{
		Market: market,
	}

	c.get(fmt.Sprintf("/shows/%s", id)).QueryStruct(params).Receive(&show, &err)

	if err != nil {
		return nil, err
	}

	return &show, nil
}

type GetSeveralShowsParams struct {
	Market string `url:"market"`
}

type GetSeveralShowsResponse struct {
	Shows []Show `json:"shows"`
}

// Get Spotify catalog information for several shows based on their Spotify IDs.
func (c *Client) GetSeveralShows(ids []string, market string) (*GetSeveralShowsResponse, error) {
	shows := GetSeveralShowsResponse{}
	var err *SpotifyError

	params := GetSeveralShowsParams{
		Market: market,
	}

	c.get("/shows").QueryStruct(params).Receive(&shows, &err)

	if err != nil {
		return nil, err
	}

	return &shows, nil
}

type GetShowEpisodesParams struct {
	Market string `url:"market"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetShowEpisodesResponse struct {
	Pagination
	Items []Episode `json:"items"`
}

// Get Spotify catalog information about an show’s episodes. Optional parameters can be used to limit the number of episodes returned.
func (c *Client) GetShowEpisodes(id string, params *GetShowEpisodesParams) (*GetShowEpisodesResponse, error) {
	episodes := GetShowEpisodesResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/shows/%s/episodes", id)).QueryStruct(params).Receive(&episodes, &err)

	if err != nil {
		return nil, err
	}

	return &episodes, nil
}

type GetUsersSavedShowsParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetUsersSavedShowsResponse struct {
	Pagination
	Items []SavedShow `json:"items"`
}

// Get a list of shows saved in the current Spotify user's library. Optional parameters can be used to limit the number of shows returned.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedShows(params *GetUsersSavedShowsParams) (*GetUsersSavedShowsResponse, error) {
	shows := GetUsersSavedShowsResponse{}
	var err *SpotifyError

	c.get("/me/shows").QueryStruct(params).Receive(&shows, &err)

	if err != nil {
		return nil, err
	}

	return &shows, nil
}

type SaveShowsForCurrentUserParams struct {
	IDs string `url:"ids"`
}

// Save one or more shows to current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) SaveShowsForCurrentUser(ids []string) error {
	var res struct{}
	var err *SpotifyError

	params := SaveShowsForCurrentUserParams{
		IDs: strings.Join(ids, ","),
	}

	c.put("/me/shows").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type RemoveUsersSavedShowsParams struct {
	// A comma-separated list of the Spotify IDs for the shows. Maximum: 50 IDs.
	IDs    string `url:"ids"`
	Market string `url:"market"`
}

// Delete one or more shows from current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedShows(ids []string, market string) error {
	var res struct{}
	var err *SpotifyError

	params := RemoveUsersSavedShowsParams{
		IDs:    strings.Join(ids, ","),
		Market: market,
	}

	c.delete("/me/shows").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type CheckUsersSavedShowsParams struct {
	IDs string `url:"ids"`
}

// Check if one or more shows is already saved in the current Spotify user's library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedShows(ids []string) error {
	var res struct{}
	var err *SpotifyError

	params := CheckUsersSavedShowsParams{
		IDs: strings.Join(ids, ","),
	}

	c.get("/me/shows/contains").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}
