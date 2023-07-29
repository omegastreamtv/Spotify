package spotify

type GetAvailableGenreSeedsResponse struct {
	Genres []string `json:"genres"`
}

// Retrieve a list of available genres seed parameter values for recommendations. (https://developer.spotify.com/documentation/web-api/reference/get-recommendations)
func (c *Client) GetAvailableGenreSeeds() (*GetAvailableGenreSeedsResponse, error) {
	genres := GetAvailableGenreSeedsResponse{}
	var spotifyErr *SpotifyError

	_, err := c.get("/recommendations/available-genre-seeds").Receive(&genres, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &genres, nil
}
