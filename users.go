package spotify

// Get detailed profile information about the current user (including the current user's username).
//
// Required scope: user-read-private, user-read-email
func (c *Client) GetCurrentUsersProfile() {}

// Get the current user's top artists or tracks based on calculated affinity.
//
// Required scope: user-top-read
func (c *Client) GetUsersTopItems() {}

// Get public profile information about a Spotify user.
func (c *Client) GetUsersProfile() {}

// Add the current user as a follower of a playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) FollowPlaylist() {}

// Remove the current user as a follower of a playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) UnfollowPlaylist() {}

// Get the current user's followed artists.
//
// Required scope: user-follow-read
func (c *Client) GetFollowedArtists() {}

// Add the current user as a follower of one or more artists or other Spotify users.
//
// Required scope: user-follow-modify
func (c *Client) FollowArtistsOrUsers() {}

// Remove the current user as a follower of one or more artists or other Spotify users.
//
// Required scope: user-follow-modify
func (c *Client) UnfollowArtistsOrUsers() {}

// Check to see if the current user is following one or more artists or other Spotify users.
//
//	Required scope: user-follow-read
func (c *Client) CheckIfUserFollowsArtistsOrUsers() {}

// Check to see if one or more Spotify users are following a specified playlist.
func (c *Client) CheckIfUsersFollowPlaylist() {}
