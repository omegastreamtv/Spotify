package spotify

type GetAvailableMarketsResponse struct {
	Markets []string `json:"markets"`
}

// Get the list of markets where Spotify is available.
func (c *Client) GetAvailableMarkets() (*GetAvailableMarketsResponse, error) {
	markets := GetAvailableMarketsResponse{}
	var err *SpotifyError

	c.get("/markets").Receive(&markets, &err)

	if err != nil {
		return nil, err
	}

	return &markets, nil
}
