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

type GetArtistsResponse struct {
	Artist
}

// Get Spotify catalog information for a single artist identified by their unique Spotify ID.
func (c *Client) GetArtist(artistId string) (*GetArtistsResponse, error) {
	artist := GetArtistsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get(fmt.Sprintf("/artists/%s", artistId)).Receive(&artist, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
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
func (c *Client) GetSeveralArtists(artistIds []string) (*GetSeveralArtistsResponse, error) {
	artists := GetSeveralArtistsResponse{}

	params := GetSeveralArtistsParams{
		IDs: strings.Join(artistIds, ","),
	}

	var spotifyErr *SpotifyError
	_, err := c.get(fmt.Sprintf("/artists?ids=%s", params.IDs)).QueryStruct(params).Receive(&artists, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
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
	Market        Market `url:"market,omitempty"`
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
func (c *Client) GetArtistsAlbums(artistId string, params *GetArtistsAlbumsParams) (*GetArtistsAlbumsResponse, error) {
	albums := GetArtistsAlbumsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get(fmt.Sprintf("/artists/%s/albums", artistId)).QueryStruct(params).Receive(&albums, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &albums, nil
}

type GetArtistsTopTracksParams struct {
	Market Market `url:"market,omitempty"`
}

type GetArtistsTopTracksResponse struct {
	Tracks []Track `json:"tracks"`
}

// Get Spotify catalog information about an artist's top tracks by country.
func (c *Client) GetArtistsTopTracks(artistId string, params *GetArtistsTopTracksParams) (*GetArtistsTopTracksResponse, error) {
	tracks := GetArtistsTopTracksResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get(fmt.Sprintf("/artists/%s/top-tracks", artistId)).QueryStruct(params).Receive(&tracks, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &tracks, nil
}

type GetArtistsRelatedArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

// Get Spotify catalog information about artists similar to a given artist. Similarity is based on analysis of the Spotify community's listening history.
func (c *Client) GetArtistsRelatedArtists(artistId string) (*GetArtistsRelatedArtistsResponse, error) {
	artists := GetArtistsRelatedArtistsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get(fmt.Sprintf("/artists/%s/related-artists", artistId)).Receive(&artists, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &artists, nil
}
