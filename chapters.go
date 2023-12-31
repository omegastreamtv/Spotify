package spotify

import "fmt"

type Chapter struct {
	// A URL to a 30 second preview (MP3 format) of the episode. null if not available.
	AudioPreviewURL string `json:"audio_preview_url"`
	// A list of the countries in which the chapter can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The number of the chapter
	ChapterNumber int `json:"chapter_number"`
	// A description of the episode. HTML tags are stripped away from this field, use html_description field in case HTML tags are needed.
	Description string `json:"description"`
	// A description of the episode. This field may contain HTML tags.
	HTMLDescription string `json:"html_description"`
	// The episode length in milliseconds.
	DurationMS int `json:"duration_ms"`
	// Whether or not the episode has explicit content (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// External URLs for this episode.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the episode.
	Href string `json:"href"`
	// The Spotify ID for the episode.
	ID string `json:"id"`
	// The cover art for the episode in various sizes, widest first.
	Images []Image `json:"images"`
	// True if the episode is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable"`
	// A list of the languages used in the episode, identified by their ISO 639-1 code.
	Languages []string `json:"languages"`
	// The name of the episode.
	Name string `json:"name"`
	// The date the episode was first released, for example "1981-12-15". Depending on the precision, it might be shown as "1981" or "1981-12".
	ReleaseDate string `json:"release_date"`
	// The precision with which release_date value is known.
	ReleaseDatePrecision string `json:"release_date_precision"`
	// The user's most recent position in the episode. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
	ResumePoint ResumePoint `json:"resume_point"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for the episode.
	URI string `json:"uri"`
	// Included in the response when a content restriction is applied.
	Restrictions Restrictions `json:"restrictions"`
}

type FullChapter struct {
	Chapter
	// The audiobook for which the chapter belongs.
	Audiobook Audiobook `json:"audiobook"`
}

type GetAChapterParams struct {
	Market Market `url:"market,omitempty"`
}

type GetAChapterResponse struct {
	Chapter
}

// Get Spotify catalog information for a single chapter.
//
// Note: Chapters are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetAChapter(chatperId string, market Market) (*GetAChapterResponse, error) {
	chapter := GetAChapterResponse{}
	var spotifyErr *SpotifyError

	params := GetAChapterParams{
		Market: market,
	}

	_, err := c.get(fmt.Sprintf("/chapters/%s", chatperId)).QueryStruct(params).Receive(&chapter, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &chapter, nil
}

type GetSeveralChaptersParams struct {
	Market Market `url:"market,omitempty"`
}

type GetSeveralChaptersResponse struct {
	Chapters []Chapter `json:"chapters"`
}

// Get Spotify catalog information for several chapters identified by their Spotify IDs.
//
// Note: Chapters are only available for the US, UK, Ireland, New Zealand and Australia markets.
func (c *Client) GetSeveralChapters(chatperIds []string, market Market) (*GetSeveralChaptersResponse, error) {
	chapters := GetSeveralChaptersResponse{}
	var spotifyErr *SpotifyError

	params := GetSeveralChaptersParams{
		Market: market,
	}

	_, err := c.get("/chapters").QueryStruct(params).Receive(&chapters, &spotifyErr)
	if err != nil {
		return nil, err
	}

	if spotifyErr != nil {
		return nil, spotifyErr
	}

	return &chapters, nil
}
