package spotify

// Get a playlist owned by a Spotify user.
func (c *Client) GetPlaylist() {}

// Change a playlist's name and public/private state. (The user must, of course, own the playlist.)
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) ChangePlaylistDetails() {}

// Get full details of the items of a playlist owned by a Spotify user.
//
// Required scope: playlist-read-private
func (c *Client) GetPlaylistItems() {}

// Either reorder or replace items in a playlist depending on the request's parameters. To reorder items, include range_start, insert_before, range_length and snapshot_id in the request's body. To replace items, include uris as either a query parameter or in the request's body. Replacing items in a playlist will overwrite its existing items. This operation can be used for replacing or clearing items in a playlist.
//
// Note: Replace and reorder are mutually exclusive operations which share the same endpoint, but have different parameters. These operations can't be applied together in a single request.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) UpdatePlaylistItems() {}

// Add one or more items to a user's playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) AddItemsToPlaylist() {}

// Remove one or more items from a user's playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) RemovePlaylistItems() {}

// Get a list of the playlists owned or followed by the current Spotify user.
//
// Required scope: playlist-read-private
func (c *Client) GetCurrentUsersPlaylists() {}

// Get a list of the playlists owned or followed by a Spotify user.
//
// Required scope: playlist-read-private, playlist-read-collaborative
func (c *Client) GetUsersPlaylists() {}

// Create a playlist for a Spotify user. (The playlist will be empty until you add tracks.)
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) CreatePlaylist() {}

// GetFeaturedPlaylists
func (c *Client) GetFeaturedPlaylists() {}

// Get a list of Spotify playlists tagged with a particular category.
func (c *Client) GetCategorysPlaylists() {}

// Get the current image associated with a specific playlist.
func (c *Client) GetPlaylistCoverImage() {}

// Replace the image used to represent a specific playlist.
//
// Required scope: ugc-image-upload, playlist-modify-public, playlist-modify-private
func (c *Client) AddCustomPlaylistCoverImage() {}
