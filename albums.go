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
	Images []Images `json:"images"`
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
	Copyrights []Copyrights `json:"copyrights"`
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
	Artists []Artists `json:"artists"`
}

type FullAlbum struct {
	Album
	// The tracks of the album.
	Tracks []Track `json:"tracks"`
}

type GetAlbumResponse struct {
	Album FullAlbum `json:"album"`
}

// Get Spotify catalog information for a single album.
func (c *Client) GetAlbum(id string, market string) (*GetAlbumResponse, error) {
	album := FullAlbum{}
	var err *SpotifyError
	c.get(fmt.Sprintf("/albums/%s", id)).Receive(&album, &err)
	if err != nil {
		return nil, err
	}

	return &GetAlbumResponse{
		Album: album,
	}, nil
}

type GetSeveralAlbumsParams struct {
	// A comma-separated list of the Spotify IDs for the albums. Maximum: 20 IDs.
	IDs string `url:"ids,omitempty"`
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
}

type GetSeveralAlbumsResponse struct {
	Albums []FullAlbum `json:"albums"`
}

// Get Spotify catalog information for multiple albums identified by their Spotify IDs.
func (c *Client) GetSeveralAlbums(ids []string, market string) (*GetSeveralAlbumsResponse, error) {
	albums := GetSeveralAlbumsResponse{}
	var err *SpotifyError

	params := GetSeveralAlbumsParams{
		// Convert ids slice to a comma-separated string
		IDs:    strings.Join(ids, ","),
		Market: market,
	}

	c.get("/albums").QueryStruct(params).Receive(&albums, &err)
	if err != nil {
		return nil, err
	}

	return &albums, nil
}

// Get Spotify catalog information about an album’s tracks. Optional parameters can be used to limit the number of tracks returned.
func (c *Client) GetAlbumTracks() {}

// Get a list of the albums saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedAlbums() {}

// Save one or more albums to the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) SaveAlbumsForCurrentUser() {}

// Remove one or more albums from the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) RemoveAlbumsForCurrentUser() {}

// Check if one or more albums is already saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedAlbums() {}

// Get a list of new album releases featured in Spotify (shown, for example, on a Spotify player’s “Browse” tab).
func (c *Client) GetNewReleases() {}
