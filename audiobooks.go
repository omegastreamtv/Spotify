package spotify

type Audiobook struct{}

// Get Spotify catalog information for a single audiobook.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetAnAudiobook() {}

// Get Spotify catalog information for several audiobooks identified by their Spotify IDs.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetSeveralAudiobooks() {}

// Get Spotify catalog information about an audiobook's chapters.
//
// Note: Audiobooks are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetAudiobookChapters() {}

// Get a list of the audiobooks saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedAudioBooks() {}

// Save one or more audiobooks to the current Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) SaveAudiobooksForCurrentUser() {}

// Remove one or more audiobooks from the Spotify user's library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedAudiobooks() {}

// Check if one or more audiobooks are already saved in the current Spotify user's library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedAudiobooks() {}
