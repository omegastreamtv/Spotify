package spotify

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

// Get Spotify catalog information for a single show identified by its unique Spotify ID.
//
// Required scope: user-read-playback-position
func (c *Client) GetShow() {}

// Get Spotify catalog information for several shows based on their Spotify IDs.
func (c *Client) GetSeveralShows() {}

// Get Spotify catalog information about an show’s episodes. Optional parameters can be used to limit the number of episodes returned.
func (c *Client) GetShowEpisodes() {}

// Get a list of shows saved in the current Spotify user's library. Optional parameters can be used to limit the number of shows returned.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedShows() {}

// Save one or more shows to current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) SaveShowsForCurrentUser() {}

// Delete one or more shows from current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedShows() {}

// Check if one or more shows is already saved in the current Spotify user's library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedShows() {}
