package spotify

type SearchForItemParams struct {
	// Your search query.
	//
	// You can narrow down your search using field filters. The available filters are album, artist, track, year, upc, tag:hipster, tag:new, isrc, and genre. Each field filter only applies to certain result types.
	//
	// The artist and year filters can be used while searching albums, artists and tracks. You can filter on a single year or a range (e.g. 1955-1960).
	//
	// The album filter can be used while searching albums and tracks.
	//
	// The genre filter can be used while searching artists and tracks.
	//
	// The isrc and track filters can be used while searching tracks.
	//
	// The upc, tag:new and tag:hipster filters can only be used while searching albums. The tag:new filter will return albums released in the past two weeks and tag:hipster can be used to return only albums with the lowest 10% popularity.
	Q string `url:"q"`
	// A comma-separated list of item types to search across. Search results include hits from all the specified item types. For example: q=abacab&type=album,track returns both albums and tracks matching "abacab".
	Type   []string `url:"type"`
	Market Market   `url:"market,omitempty"`
	// The maximum number of results to return in each item type.
	Limit int `url:"limit,omitempty"`
	// The index of the first result to return. Use with limit to get the next page of search results.
	Offset int `url:"offset,omitempty"`
	// If include_external=audio is specified it signals that the client can play externally hosted audio content, and marks the content as playable in the response. By default externally hosted audio content is marked as unplayable in the response.
	IncludeExternal string `url:"include_external,omitempty"`
}

type SearchForItemResponse struct {
	Tracks struct {
		Pagination
		Items []Track `json:"items"`
	} `json:"tracks"`
	Artists struct {
		Pagination
		Items []Artist `json:"items"`
	} `json:"artists"`
	Albums struct {
		Pagination
		Items []Album `json:"items"`
	} `json:"albums"`
	Playlists struct {
		Pagination
		Items []Playlist `json:"items"`
	} `json:"playlists"`
	Shows struct {
		Pagination
		Items []Show `json:"items"`
	} `json:"shows"`
	Episodes struct {
		Pagination
		Items []Episode `json:"items"`
	} `json:"episodes"`
	Audiobooks struct {
		Pagination
		Items []Audiobook `json:"items"`
	} `json:"audiobooks"`
}

// Get Spotify catalog information about albums, artists, playlists, tracks, shows, episodes or audiobooks that match a keyword string.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) SearchForItem(params SearchForItemParams) (*SearchForItemResponse, error) {
	result := SearchForItemResponse{}
	var err *SpotifyError

	c.get("/search").Receive(&result, &err)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
