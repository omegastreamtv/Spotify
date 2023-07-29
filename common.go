package spotify

// The copyright statements of the content.
type Copyright struct {
	// The copyright text for this content.
	Text string `json:"text"`
	// The type of copyright: C = the copyright, P = the sound recording (performance) copyright.
	Type string `json:"type"`
}

// Known external IDs for the content.
type ExternalIDs struct {
	// Known external IDs for the content.
	//
	// https://en.wikipedia.org/wiki/International_Standard_Recording_Code
	ISRC string `json:"isrc"`
	// International Article Number
	//
	// https://en.wikipedia.org/wiki/International_Article_Number
	EAN string `json:"ean"`
	// Universal Product Code
	//
	// https://en.wikipedia.org/wiki/Universal_Product_Code
	UPC string `json:"upc"`
}

// External URLs for this context.
type ExternalURLs struct {
	// The Spotify URL for the object.
	Spotify string `json:"spotify"`
}

type Followers struct {
	// This will always be set to null, as the Web API does not support it at the moment.
	Href string `json:"href,omitempty"`
	// The total number of followers.
	Total int `json:"total"`
}

// The cover art for the content in various sizes, widest first.
type Image struct {
	// The source URL of the image.
	URL string `json:"url"`
	// The image height in pixels.
	Height int `json:"height"`
	// The image width in pixels.
	Width int `json:"width"`
}

type Pagination struct {
	// A link to the Web API endpoint returning the full result of the request
	Href string `json:"href"`
	// The maximum number of items in the response (as set in the query or by default).
	Limit int `json:"limit"`
	// URL to the next page of items. (null if none)
	Next string `json:"next"`
	// The offset of the items returned (as set in the query or by default)
	Offset int `json:"offset"`
	// URL to the previous page of items. (null if none)
	Previous string `json:"previous"`
	// The total number of items available to return.
	Total int `json:"total"`
}

// Included in the response when a content restriction is applied.
type Restrictions struct {
	// The reason for the restriction. Supported values:
	//
	// market - The content item is not available in the given market.
	// product - The content item is not available for the user's subscription type.
	// explicit - The content item is explicit and the user's account is set to not play explicit content.
	//
	// Additional reasons may be added in the future. Note: If you use this field, make sure that your application safely handles unknown values.
	Reason string `json:"reason"`
}

// The user's most recent position in the episode or chapter. Set if the supplied access token is a user token and has the scope 'user-read-playback-position'.
type ResumePoint struct {
	// Whether or not the episode has been fully played by the user.
	FullyPlayed bool `json:"fully_played"`
	// The user's most recent position in the episode in milliseconds.
	ResumePositionMS int `json:"resume_position_ms"`
}
