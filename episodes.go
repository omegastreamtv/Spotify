package spotify

import (
	"fmt"
	"strings"
)

type Episode struct {
	// A URL to a 30 second preview (MP3 format) of the episode. null if not available.
	AudioPreviewURL string `json:"audio_preview_url"`
	// A description of the episode. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the episode. This field may contain HTML tags.
	HTMLDescription string `json:"html_description"`
	// The episode length in milliseconds.
	DurationMS int `json:"duration_ms"`
	// Whether or not the episode has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this episode.
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	// A link to the Web API endpoint providing full details of the episode.
	Href string `json:"href"`
	// The Spotify ID for the episode.
	ID string `json:"id"`
	// The cover art for the episode in various sizes, widest first.
	Images []Image `json:"images"`
	// True if the episode is hosted outside of Spotify's CDN.
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// True if the episode is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable"`
	// A list of the languages used in the episode, identified by their ISO 639-1 code.
	Languages []string `json:"languages"`
	// The name of the episode.
	Name string `json:"name"`
	// The date the episode was first released, for example "1981-12-15". Depending on the precision, it might be shown as "1981" or "1981-12".
	ReleaseDate string `json:"release_date"`
	// The precision with which release_date value is known.
	ReleaseDatePrecision string `json:"release_date_precision"`
	// The user's most recent position in the episode. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
	ResumePoint ResumePoint `json:"resume_point"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the episode.
	URI string `json:"uri"`
	// Included in the response when a content restriction is applied.
	Restrictions Restrictions `json:"restrictions"`
}

type FullEpisode struct {
	Episode
	// The show on which the episode belongs.
	Show Show `json:"show"`
}

type SavedEpisode struct {
	AddedAt string `json:"added_at"`
	FullEpisode
}

type GetEpisodeParams struct {
	Market Market `url:"market"`
}

type GetEpisodeResponse struct {
	FullEpisode
}

// Get Spotify catalog information for a single episode identified by its unique Spotify ID.
func (c *Client) GetEpisode(id string, market Market) (*GetEpisodeResponse, error) {
	episode := GetEpisodeResponse{}
	var err *SpotifyError

	params := GetEpisodeParams{
		Market: market,
	}

	c.get(fmt.Sprintf("/episodes/%s", id)).QueryStruct(params).Receive(&episode, &err)

	if err != nil {
		return nil, err
	}

	return &episode, nil
}

type GetSeveralEpisodesParams struct {
	// A comma-separated list of the Spotify IDs for the episodes. Maximum: 50 IDs.
	IDs    string `url:"ids"`
	Market Market `url:"market"`
}

type GetSeveralEpisodesResponse struct {
	Episodes []FullEpisode `json:"episodes"`
}

// Get Spotify catalog information for several episodes based on their Spotify IDs.
//
// Required scope: user-read-playback-position
func (c *Client) GetSeveralEpisodes(ids []string, market Market) (*GetSeveralEpisodesResponse, error) {
	episodes := GetSeveralEpisodesResponse{}
	var err *SpotifyError

	params := GetSeveralEpisodesParams{
		IDs:    strings.Join(ids, ","),
		Market: market,
	}

	c.get("/episodes").QueryStruct(params).Receive(&episodes, &err)

	if err != nil {
		return nil, err
	}

	return &episodes, nil
}

type GetUsersSavedEpisodesParams struct {
	Market Market `url:"market"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetUsersSavedEpisodesResponse struct {
	Pagination
	Items []SavedEpisode `json:"items"`
}

// Get a list of the episodes saved in the current Spotify user's library.
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) GetUsersSavedEpisodes(params *GetUsersSavedEpisodesParams) (*GetUsersSavedEpisodesResponse, error) {
	episodes := GetUsersSavedEpisodesResponse{}
	var err *SpotifyError

	c.get("/me/episodes").QueryStruct(params).Receive(&episodes, &err)

	if err != nil {
		return nil, err
	}

	return &episodes, nil
}

type SaveEpisodesForCurrentUserBody struct {
	// A JSON array of the Spotify IDs. A maximum of 50 items can be specified in one request.
	IDs []string `json:"ids"`
}

// Save one or more episodes to the current user's library.
//
// Required scope: user-library-modify
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) SaveEpisodesForCurrentUser(ids []string) error {
	var res struct{}
	var err *SpotifyError

	payload := SaveEpisodesForCurrentUserBody{
		IDs: ids,
	}

	c.put("/me/episodes").BodyJSON(payload).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type RemoveUsersSavedEpisodesBody struct {
	// A JSON array of the Spotify IDs. A maximum of 50 items can be specified in one request.
	IDs []string `json:"ids"`
}

// Remove one or more episodes from the current user's library.
//
// Required scope: user-library-modify
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) RemoveUsersSavedEpisodes(ids []string) error {
	var res struct{}
	var err *SpotifyError

	payload := RemoveUsersSavedEpisodesBody{
		IDs: ids,
	}

	c.delete("/me/episodes").BodyJSON(payload).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type CheckUsersSavedEpisodesParams struct {
	// A comma-separated list of the Spotify IDs. Maximum: 50 IDs.
	IDs string `url:"ids"`
}

// Check if one or more episodes is already saved in the current Spotify user's 'Your Episodes' library.
//
// Required scope: user-library-read
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) CheckUsersSavedEpisodes(ids []string) ([]bool, error) {
	var res []bool
	var err *SpotifyError

	params := CheckUsersSavedEpisodesParams{
		IDs: strings.Join(ids, ","),
	}

	c.get("/me/episodes/contains").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return nil, err
	}

	return res, nil
}
