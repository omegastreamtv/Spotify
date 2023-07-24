package spotify

import (
	"fmt"
	"strings"
)

type Audiobook struct {
	// The author(s) for the audiobook.
	Authors []struct {
		// The name of the author.
		Name string `json:"name"`
	} `json:"authors"`
	// A list of the countries in which the audiobook can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The copyright statements of the audiobook.
	Copyrights []Copyright `json:"copyrights"`
	// A description of the audiobook. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the audiobook. This field may contain HTML tags.
	HTMLDescription string `json:"html_description"`
	// The edition of the audiobook.
	Edition string `json:"edition"`
	// Whether or not the audiobook has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this audiobook.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the audiobook.
	Href string `json:"href"`
	// The Spotify ID for the audiobook.
	ID string `json:"id"`
	// The cover art for the audiobook in various sizes, widest first.
	Images []Image `json:"images"`
	// A list of the languages used in the audiobook, identified by their ISO 639 code.
	Languages []string `json:"languages"`
	// The media type of the audiobook.
	MediaType string `json:"media_type"`
	// The name of the audiobook.
	Name string `json:"name"`
	// The narrator(s) for the audiobook.
	Narrators []struct {
		// The name of the Narrator.
		Name string `json:"name"`
	} `json:"narrators"`
	// The publisher of the audiobook.
	Publisher string `json:"publisher"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the audiobook.
	URI string `json:"uri"`
	// The number of chapters in this audiobook.
	TotalChapters int `json:"total_chapters"`
}

type FullAudiobook struct {
	Audiobook
	// The chapters of the audiobook.
	Chapters struct {
		Pagination
		Items []Chapter `json:"items"`
	} `json:"chapters"`
}

type GetAnAudiobookParams struct {
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market,omitempty"`
}

type GetAnAudiobookResponse struct {
	FullAudiobook
}

// Get Spotify catalog information for a single audiobook.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetAnAudiobook(id string, market string) (*GetAnAudiobookResponse, error) {
	audiobook := GetAnAudiobookResponse{}
	var err *SpotifyError

	params := GetAnAudiobookParams{
		Market: market,
	}

	c.get(fmt.Sprintf("/audiobooks/%s", id)).QueryStruct(params).Receive(&audiobook, &err)

	if err != nil {
		return nil, err
	}

	return &audiobook, nil
}

type GetSeveralAudiobooksParams struct {
	// A comma-separated list of the Spotify IDs.
	IDs string `url:"ids"`
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market"`
}

type GetSeveralAudiobooksResponse struct {
	Audiobooks []Audiobook `json:"audiobooks"`
}

// Get Spotify catalog information for several audiobooks identified by their Spotify IDs.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetSeveralAudiobooks(ids []string, market string) (*GetSeveralAudiobooksResponse, error) {
	audiobooks := GetSeveralAudiobooksResponse{}
	var err *SpotifyError

	params := GetSeveralAudiobooksParams{
		IDs:    strings.Join(ids, ","),
		Market: market,
	}

	c.get("/audiobooks").QueryStruct(params).Receive(&audiobooks, &err)

	if err != nil {
		return nil, err
	}

	return &audiobooks, nil
}

type GetAudiobookChaptersParams struct {
	// An ISO 3166-1 alpha-2 country code. If a country code is specified, only content that is available in that market will be returned.
	//
	// If a valid user access token is specified in the request header, the country associated with the user account will take priority over this parameter.
	//
	// Note: If neither market or user country are provided, the content is considered unavailable for the client.
	// Users can view the country that is associated with their account in the account settings.
	Market string `url:"market"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetAudiobookChaptersResponse struct {
	Pagination
	Items []Chapter `json:"items"`
}

// Get Spotify catalog information about an audiobook's chapters.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetAudiobookChapters(id string, params GetAudiobookChaptersParams) (*GetAudiobookChaptersResponse, error) {
	chapters := GetAudiobookChaptersResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/audiobooks/%s/chapters", id)).QueryStruct(params).Receive(&chapters, &err)

	if err != nil {
		return nil, err
	}

	return &chapters, nil
}

type GetUsersSavedAudioBooksParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetUsersSavedAudioBooksResponse struct {
	Pagination
	Items []Audiobook `json:"items"`
}

// Get a list of the audiobooks saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedAudioBooks(limit int, offset int) (*GetUsersSavedAudioBooksResponse, error) {
	audiobooks := GetUsersSavedAudioBooksResponse{}
	var err *SpotifyError

	params := GetUsersSavedAudioBooksParams{
		Limit:  limit,
		Offset: offset,
	}

	c.get("/me/audiobooks").QueryStruct(params).Receive(&audiobooks, &err)

	if err != nil {
		return nil, err
	}

	return &audiobooks, nil
}

type SaveAudiobooksForCurrentUserParams struct {
	// A comma-separated list of the Spotify IDs.
	IDs string `url:"ids"`
}

type SaveAudiobooksForCurrentUserBody struct {
	// A JSON array of the Spotify IDs.
	//
	// A maximum of 50 items can be specified in one request. Note: if the ids parameter is present in the query string, any IDs listed here in the body will be ignored.
	IDs []string `json:"ids"`
}

// Save one or more audiobooks to the current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) SaveAudiobooksForCurrentUser(ids []string) error {
	var res struct{}
	var err *SpotifyError

	params := SaveAudiobooksForCurrentUserParams{
		IDs: strings.Join(ids, ","),
	}

	payload := SaveAudiobooksForCurrentUserBody{
		IDs: ids,
	}

	c.put("/me/audiobooks").QueryStruct(params).BodyJSON(payload).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type RemoveUsersSavedAudiobooksParams struct {
	// A comma-separated list of the Spotify IDs.
	IDs string `url:"ids"`
}

// Remove one or more audiobooks from the Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedAudiobooks(ids []string) error {
	var res struct{}
	var err *SpotifyError

	params := RemoveUsersSavedAudiobooksParams{
		IDs: strings.Join(ids, ","),
	}

	c.delete("/me/audiobooks").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type CheckUsersSavedAudiobooksParams struct {
	// A comma-separated list of the Spotify IDs.
	IDs string `url:"ids"`
}

// Check if one or more audiobooks are already saved in the current Spotify user's library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedAudiobooks(ids []string) ([]bool, error) {
	var res []bool
	var err *SpotifyError

	params := CheckUsersSavedAudiobooksParams{
		IDs: strings.Join(ids, ","),
	}

	c.get("/me/audiobooks/contains").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return nil, err
	}

	return res, nil
}
