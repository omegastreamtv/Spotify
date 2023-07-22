package spotify

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
