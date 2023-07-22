package spotify

// Get Spotify catalog information for a single artist identified by their unique Spotify ID.
func (c *Client) GetArtist() {}

// Get Spotify catalog information for several artists based on their Spotify IDs.
func (c *Client) GetSeveralArtists() {}

// Get Spotify catalog information about an artist's albums.
func (c *Client) GetArtistsAlbums() {}

// Get Spotify catalog information about an artist's top tracks by country.
func (c *Client) GetArtistsTopTracks() {}

// Get Spotify catalog information about artists similar to a given artist. Similarity is based on analysis of the Spotify community's listening history.
func (c *Client) GetArtistsRelatedArtists() {}
