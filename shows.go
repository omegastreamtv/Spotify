package spotify

import (
	"fmt"
	"strings"
)

type Show struct {
	// A list of the countries in which the show can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The copyright statements of the show.
	Copyrights []Copyright `json:"copyrights"`
	// A description of the show. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the show. This field may contain HTML tags.
	HTMLDescription string `json:"html_description"`
	// Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this show.
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	// A link to the Web API endpoint providing full details of the show.
	Href string `json:"href"`
	// The Spotify ID for the show.
	ID string `json:"id"`
	// The cover art for the show in various sizes, widest first.
	Images []Image `json:"images"`
	// True if all of the shows episodes are hosted outside of Spotify's CDN. This field might be null in some cases.
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// A list of the languages used in the show, identified by their ISO 639 code.
	Languages []string `json:"languages"`
	// The media type of the show.
	MediaType string `json:"media_type"`
	// The name of the episode.
	Name string `json:"name"`
	// The publisher of the show.
	Publisher string `json:"publisher"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the show.
	URI string `json:"uri"`
	// The total number of episodes in the show.
	TotalEpisodes int `json:"total_episodes"`
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
	Market Market `url:"market"`
}

type GetShowResponse struct {
	Show
}

// Get Spotify catalog information for a single show identified by its unique Spotify ID.
//
// Required scope: user-read-playback-position
func (c *Client) GetShow(showId string, market Market) (*GetShowResponse, error) {
	show := GetShowResponse{}
	var err *SpotifyError

	params := GetShowParams{
		Market: market,
	}

	c.get(fmt.Sprintf("/shows/%s", showId)).QueryStruct(params).Receive(&show, &err)

	if err != nil {
		return nil, err
	}

	return &show, nil
}

type GetSeveralShowsParams struct {
	// A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	IDs    string `url:"ids,omitempty"`
	Market Market `url:"market"`
}

type GetSeveralShowsResponse struct {
	Shows []Show `json:"shows"`
}

// Get Spotify catalog information for several shows based on their Spotify IDs.
func (c *Client) GetSeveralShows(showIds []string, market Market) (*GetSeveralShowsResponse, error) {
	shows := GetSeveralShowsResponse{}
	var err *SpotifyError

	params := GetSeveralShowsParams{
		IDs:    strings.Join(showIds, ","),
		Market: market,
	}

	c.get("/shows").QueryStruct(params).Receive(&shows, &err)

	if err != nil {
		return nil, err
	}

	return &shows, nil
}

type GetShowEpisodesParams struct {
	Market Market `url:"market"`
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
func (c *Client) GetShowEpisodes(showId string, params *GetShowEpisodesParams) (*GetShowEpisodesResponse, error) {
	episodes := GetShowEpisodesResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/shows/%s/episodes", showId)).QueryStruct(params).Receive(&episodes, &err)

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
func (c *Client) SaveShowsForCurrentUser(showIds []string) error {
	var res struct{}
	var err *SpotifyError

	params := SaveShowsForCurrentUserParams{
		IDs: strings.Join(showIds, ","),
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
	Market Market `url:"market"`
}

// Delete one or more shows from current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedShows(showIds []string, market Market) error {
	var res struct{}
	var err *SpotifyError

	params := RemoveUsersSavedShowsParams{
		IDs:    strings.Join(showIds, ","),
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
func (c *Client) CheckUsersSavedShows(showIds []string) error {
	var res struct{}
	var err *SpotifyError

	params := CheckUsersSavedShowsParams{
		IDs: strings.Join(showIds, ","),
	}

	c.get("/me/shows/contains").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}
