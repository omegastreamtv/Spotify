package spotify

import (
	"fmt"
	"strings"
)

// An artist
type Artist struct {
	// Known external URLs for this artist.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// Information about the followers of the artist.
	Followers Followers `json:"followers,omitempty"`
	// A list of the genres the artist is associated with. If not yet classified, the array is empty.
	Genres []string `json:"genres,omitempty"`
	// A link to the Web API endpoint providing full details of the artist.
	Href string `json:"href"`
	// The Spotify ID for the artist.
	ID string `json:"id"`
	// Images of the artist in various sizes, widest first.
	Images []Image `json:"images,omitempty"`
	// The name of the artist.
	Name string `json:"name"`
	// The popularity of the artist. The value will be between 0 and 100, with 100 being the most popular. The artist's popularity is calculated from the popularity of all the artist's tracks.
	Popularity int `json:"popularity,omitempty"`
	// The object type.
	//
	// Allowed values: "artist"
	Type string `json:"type"`
	// The Spotify URI for the artist.
	URI string `json:"uri"`
}

type Followers struct {
	Href  string `json:"href,omitempty"`
	Total int    `json:"total"`
}

type GetArtistsResponse struct {
	Artist
}

// Get Spotify catalog information for a single artist identified by their unique Spotify ID.
func (c *Client) GetArtist(id string) (*GetArtistsResponse, error) {
	artist := GetArtistsResponse{}
	var err *SpotifyError
	c.get(fmt.Sprintf("/artists/%s", id)).Receive(&artist, &err)
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

type GetSeveralArtistsParams struct {
	// A comma-separated list of the Spotify IDs for the artists. Maximum: 50 IDs.
	IDs string `url:"ids"`
}

type GetSeveralArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

// Get Spotify catalog information for several artists based on their Spotify IDs.
func (c *Client) GetSeveralArtists(ids []string) (*GetSeveralArtistsResponse, error) {
	artists := GetSeveralArtistsResponse{}

	params := GetSeveralArtistsParams{
		IDs: strings.Join(ids, ","),
	}

	var err *SpotifyError
	c.get(fmt.Sprintf("/artists?ids=%s", params.IDs)).QueryStruct(params).Receive(&artists, &err)
	if err != nil {
		return nil, err
	}

	return &artists, nil
}

type GetArtistsAlbumsParams struct {
	// A comma-separated list of keywords that will be used to filter the response. If not supplied, all album types will be returned.
	//
	// Valid values are:
	//
	// - album
	//
	// - single
	//
	// - appears_on
	//
	// - compilation
	//
	// For example: include_groups=album,single.
	//
	// Example value: "single,appears_on"
	IncludeGroups string `url:"include_groups,omitempty"`
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	//
	// Example value: 10
	//
	// Default value: 20
	//
	// Range: 0-50
	Limit int `url:"limit,omitempty"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	//
	// Example value: 5
	//
	// Default value: 0
	Offset int `url:"offset,omitempty"`
}

type GetArtistsAlbumsResponse struct {
	Pagination
	Items []Album `json:"items"`
}

// Get Spotify catalog information about an artist's albums.
func (c *Client) GetArtistsAlbums(id string, params *GetArtistsAlbumsParams) (*GetArtistsAlbumsResponse, error) {
	albums := GetArtistsAlbumsResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/artists/%s/albums", id)).QueryStruct(params).Receive(&albums, &err)
	if err != nil {
		return nil, err
	}

	return &albums, nil
}

type GetArtistsTopTracksParams struct {
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
}

type GetArtistsTopTracksResponse struct {
	Tracks []Track `json:"tracks"`
}

// Get Spotify catalog information about an artist's top tracks by country.
func (c *Client) GetArtistsTopTracks(id string, params *GetArtistsTopTracksParams) (*GetArtistsTopTracksResponse, error) {
	tracks := GetArtistsTopTracksResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/artists/%s/top-tracks", id)).QueryStruct(params).Receive(&tracks, &err)
	if err != nil {
		return nil, err
	}

	return &tracks, nil
}

type GetArtistsRelatedArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

// Get Spotify catalog information about artists similar to a given artist. Similarity is based on analysis of the Spotify community's listening history.
func (c *Client) GetArtistsRelatedArtists(id string) (*GetArtistsRelatedArtistsResponse, error) {
	artists := GetArtistsRelatedArtistsResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/artists/%s/related-artists", id)).Receive(&artists, &err)
	if err != nil {
		return nil, err
	}

	return &artists, nil
}
