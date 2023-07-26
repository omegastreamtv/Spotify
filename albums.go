package spotify

import (
	"fmt"
	"strings"
)

type Album struct {
	// The type of the album.
	//
	// Allowed values: "album", "single", "compilation"
	AlbumType string `json:"album_type"`
	// The number of tracks in the album.
	TotalTracks int `json:"total_tracks"`
	// The markets in which the album is available: ISO 3166-1 alpha-2 country codes.
	//
	// NOTE: an album is considered available in a market when at least 1 of its tracks is available in that market.
	AvailableMarkets []string `json:"available_markets"`
	// Known external URLs for this album.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the album.
	Href string `json:"href"`
	// The Spotify ID for the album.
	ID string `json:"id"`
	// The cover art for the album in various sizes, widest first.
	Images []Image `json:"images"`
	// The name of the album. In case of an album takedown, the value may be an empty string.
	Name string `json:"name"`
	// The date the album was first released.
	ReleaseDate string `json:"release_date"`
	// The precision with which release_date value is known.
	ReleaseDatePrecision string `json:"release_date_precision"`
	// Included in the response when a content restriction is applied.
	Restrictions Restrictions `json:"restrictions"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the album.
	URI string `json:"uri"`
	// The copyright statements of the album.
	Copyrights []Copyright `json:"copyrights"`
	// Known external IDs for the album.
	ExternalIDs ExternalIDs `json:"external_ids"`
	// A list of the genres the album is associated with. If not yet classified, the array is empty.
	Genres []string `json:"genres"`
	// The label associated with the album.
	Label string `json:"label"`
	// The popularity of the album. The value will be between 0 and 100, with 100 being the most popular.
	Popularity int `json:"popularity"`
	// The field is present when getting an artist's albums. Compare to album_type this field represents relationship between the artist and the album.
	//
	// Allowed values: "album", "single", "compilation", "appears_on"
	AlbumGroup string `json:"album_group,omitempty"`
	// The artists of the album. Each artist object includes a link in href to more detailed information about the artist.
	Artists []Artist `json:"artists"`
}

type FullAlbum struct {
	Album
	// The tracks of the album.
	Tracks []Track `json:"tracks"`
}

type SavedAlbum struct {
	AddedAt string    `json:"added_at"`
	Album   FullAlbum `json:"album"`
}

type GetAlbumResponse struct {
	Album
	Tracks struct {
		Pagination
		Items []Track `json:"items"`
	} `json:"tracks"`
}

// Get Spotify catalog information for a single album.
func (c *Client) GetAlbum(albumId string, market string) (*GetAlbumResponse, error) {
	album := GetAlbumResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/albums/%s", albumId)).Receive(&album, &err)

	if err != nil {
		return nil, err
	}

	return &album, nil
}

type GetSeveralAlbumsParams struct {
	// A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	IDs    string `url:"ids,omitempty"`
	Market Market `url:"market,omitempty"`
}

type GetSeveralAlbumsResponse struct {
	Albums []FullAlbum `json:"albums"`
}

// Get Spotify catalog information for multiple albums identified by their Spotify IDs.
func (c *Client) GetSeveralAlbums(albumIds []string, market Market) (*GetSeveralAlbumsResponse, error) {
	albums := GetSeveralAlbumsResponse{}
	var err *SpotifyError

	params := GetSeveralAlbumsParams{
		// Convert ids slice to a comma-separated string
		IDs:    strings.Join(albumIds, ","),
		Market: market,
	}

	c.get("/albums").QueryStruct(params).Receive(&albums, &err)

	if err != nil {
		return nil, err
	}

	return &albums, nil
}

type GetAlbumTracksParams struct {
	Market Market `url:"market,omitempty"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	//
	// Example value: 10
	//
	// Default value: 20
	//
	// Range: 0 - 50
	Limit int `url:"limit,omitempty"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	//
	// Example value: 5
	//
	// Default value: 0
	Offset int `url:"offset,omitempty"`
}

type GetAlbumTracksResponse struct {
	Pagination
	Items []Track `json:"items"`
}

// Get Spotify catalog information about an album’s tracks. Optional parameters can be used to limit the number of tracks returned.
func (c *Client) GetAlbumTracks(albumId string, params *GetAlbumTracksParams) (*GetAlbumTracksResponse, error) {
	tracks := GetAlbumTracksResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/albums/%s/tracks", albumId)).QueryStruct(params).Receive(&tracks, &err)

	if err != nil {
		return nil, err
	}

	return &tracks, nil
}

type GetUsersSavedAlbumsParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int    `url:"offset"`
	Market Market `url:"market"`
}

type GetUsersSavedAlbumsResponse struct {
	Pagination
	Items []SavedAlbum `json:"items"`
}

// Get a list of the albums saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedAlbums(params *GetUsersSavedAlbumsParams) (*GetUsersSavedAlbumsResponse, error) {
	albums := GetUsersSavedAlbumsResponse{}
	var err *SpotifyError

	c.get("/me/albums").QueryStruct(params).Receive(&albums, &err)

	if err != nil {
		return nil, err
	}

	return &albums, nil
}

type SaveAlbumsForCurrentUserBody struct {
	// A JSON array of the Spotify IDs. A maximum of 50 items can be specified in one request.
	IDs []string `json:"ids"`
}

// Save one or more albums to the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) SaveAlbumsForCurrentUser(albumIds []string) error {
	var res struct{}
	var err *SpotifyError

	payload := SaveAlbumsForCurrentUserBody{
		IDs: albumIds,
	}

	c.put("/me/albums").BodyJSON(payload).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type RemoveAlbumsForCurrentUserBody struct {
	// A JSON array of the Spotify IDs. A maximum of 50 items can be specified in one request.
	IDs []string `json:"ids"`
}

// Remove one or more albums from the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) RemoveAlbumsForCurrentUser(albumIds []string) error {
	var res struct{}
	var err *SpotifyError

	payload := RemoveAlbumsForCurrentUserBody{
		IDs: albumIds,
	}

	c.delete("/me/albums").BodyJSON(payload).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type CheckUsersSavedAlbumsParams struct {
	// A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	IDs string `url:"ids"`
}

// Check if one or more albums is already saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedAlbums(albumIds []string) ([]bool, error) {
	var res []bool
	var err *SpotifyError

	params := CheckUsersSavedAlbumsParams{
		IDs: strings.Join(albumIds, ","),
	}

	c.get("/me/albums/contains").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetNewReleasesParams struct {
	// A country: an ISO 3166-1 alpha-2 country code. Provide this parameter if you want the list of returned items to be relevant to a particular country. If omitted, the returned items will be relevant to all countries.
	Country string `url:"country"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetNewReleasesResponse struct {
	Albums struct {
		Pagination
		Items []Album `json:"items"`
	} `json:"albums"`
}

// Get a list of new album releases featured in Spotify (shown, for example, on a Spotify player’s “Browse” tab).
func (c *Client) GetNewReleases(params *GetNewReleasesParams) (*GetNewReleasesResponse, error) {
	releases := GetNewReleasesResponse{}
	var err *SpotifyError

	c.get("/browse/new-releases").QueryStruct(params).Receive(&releases, &err)

	if err != nil {
		return nil, err
	}

	return &releases, nil
}
