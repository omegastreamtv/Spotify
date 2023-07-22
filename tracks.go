package spotify

func (c *Client) GetTrack() {}

// Get Spotify catalog information for multiple tracks based on their Spotify IDs.
func (c *Client) GetSeveralTracks() {}

// Get a list of the songs saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) GetUsersSavedTracks() {}

// Save one or more tracks to the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) SaveTracksForCurrentUser() {}

// Remove one or more tracks from the current user's 'Your Music' library.
//
// Required scope: user-library-modify
func (c *Client) RemoveUsersSavedTracks() {}

// Check if one or more tracks is already saved in the current Spotify user's 'Your Music' library.
//
// Required scope: user-library-read
func (c *Client) CheckUsersSavedTracks() {}

// Get audio features for multiple tracks based on their Spotify IDs.
func (c *Client) GetTracksAudioFeatures() {}

// Get Track's Audio Features
//
// Get audio feature information for a single track identified by its unique Spotify ID.
func (c *Client) GetATracksAudioFeatures() {}

// Get a low-level audio analysis for a track in the Spotify catalog. The audio analysis describes the track’s structure and musical content, including rhythm, pitch, and timbre.
func (c *Client) GetTracksAudioAnalysis() {}

// Recommendations are generated based on the available information for a given seed entity and matched against similar artists and tracks. If there is sufficient information about the provided seeds, a list of tracks will be returned together with pool size details.
//
// For artists and tracks that are very new or obscure there might not be enough data to generate a list of tracks.
func (c *Client) GetRecommendations() {}
