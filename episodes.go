package spotify

// Get Spotify catalog information for a single episode identified by its unique Spotify ID.
func (c *Client) GetEpisode() {}

// Get Spotify catalog information for several episodes based on their Spotify IDs.
//
// Required scope: user-read-playback-position
func (c *Client) GetSeveralEpisodes() {}

// Get a list of the episodes saved in the current Spotify user's library.
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) GetUsersSavedEpisodes() {}

// Save one or more episodes to the current user's library.
//
// Required scope: user-library-modify
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) SaveEpisodesForCurrentUser() {}

// Remove one or more episodes from the current user's library.
//
// Required scope: user-library-modify
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) RemoveUsersSavedEpisodes() {}

// Check if one or more episodes is already saved in the current Spotify user's 'Your Episodes' library.
//
// Required scope: user-library-read
//
// This API endpoint is in beta and could change without warning. Please share any feedback that you have, or issues that you discover, in our developer community forum. (https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer)
func (c *Client) CheckUsersSavedEpisodes() {}
