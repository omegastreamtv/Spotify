package spotify

import "fmt"

type GetSeveralBrowseCategoriesParams struct {
	// A country: an ISO 3166-1 alpha-2 country code. Provide this parameter if you want to narrow the list of returned categories to those relevant to a particular country. If omitted, the returned items will be globally relevant.
	Country string `url:"country"`
	// The desired language, consisting of an ISO 639-1 language code and an ISO 3166-1 alpha-2 country code, joined by an underscore. For example: es_MX, meaning "Spanish (Mexico)". Provide this parameter if you want the category metadata returned in a particular language.
	//
	// Note: if locale is not supplied, or if the specified language is not available, all strings will be returned in the Spotify default language (American English). The locale parameter, combined with the country parameter, may give odd results if not carefully matched. For example country=SE&locale=de_DE will return a list of categories relevant to Sweden but as German language strings.
	Locale string `url:"locale"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type Category struct {
	Href  string  `json:"href"`
	Icons []Image `json:"icons"`
	ID    string  `json:"id"`
	Name  string  `json:"name"`
}

type GetSeveralBrowseCategoriesResponse struct {
	Categories struct {
		Pagination
		Items []Category `json:"items"`
	} `json:"categories"`
}

// Get a list of categories used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
func (c *Client) GetSeveralBrowseCategories(params *GetSeveralBrowseCategoriesParams) (*GetSeveralBrowseCategoriesResponse, error) {
	categories := GetSeveralBrowseCategoriesResponse{}
	var err *SpotifyError

	c.get("/browse/categories").QueryStruct(params).Receive(&categories, &err)

	if err != nil {
		return nil, err
	}

	return &categories, nil
}

type GetSingleBrowseCategoryParams struct {
	Country string `url:"country"`
	Locale  string `url:"locale"`
}

type GetSingleBrowseCategoryResponse struct {
	Href  string  `json:"href"`
	Icons []Image `json:"icons"`
	ID    string  `json:"id"`
	Name  string  `json:"name"`
}

// Get a single category used to tag items in Spotify (on, for example, the Spotify player’s “Browse” tab).
func (c *Client) GetSingleBrowseCategory(id string, params *GetSingleBrowseCategoryParams) (*GetSingleBrowseCategoryResponse, error) {
	category := GetSingleBrowseCategoryResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/browse/categories/%s", id)).QueryStruct(params).Receive(&category, &err)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
