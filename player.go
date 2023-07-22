package spotify

import (
	"fmt"
)

// External URLs for this context.
type ExternalURLs struct {
	// The Spotify URL for the object.
	Spotify string `json:"spotify"`
}

// The cover art for the episode in various sizes, widest first.
type Images struct {
	// The source URL of the image.
	URL string `json:"url"`
	// The image height in pixels.
	Height int `json:"height"`
	// The image width in pixels.
	Width int `json:"width"`
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

// The copyright statements of the show.
type Copyrights struct {
	// The copyright text for this content.
	Text string `json:"text"`
	// The type of copyright: C = the copyright, P = the sound recording (performance) copyright.
	Type string `json:"type"`
}

// Known external IDs for the track.
type ExternalIDs struct {
	// Known external IDs for the track.
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

type Artists struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Images     `json:"images,omitempty"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

// Allows to update the user interface based on which playback actions are available within the current context.
type Actions struct {
	// Interrupting playback. Optional field.
	InterruptingPlayback bool `json:"interrupting_playback"`
	// Pausing. Optional field.
	Pausing bool `json:"pausing"`
	// Resuming. Optional field.
	Resuming bool `json:"resuming"`
	// Seeking playback location. Optional field.
	Seeking bool `json:"seeking"`
	// Skipping to the next context. Optional field.
	SkippingNext bool `json:"skipping_next"`
	// Skipping to the previous context. Optional field.
	SkippingPrev bool `json:"skipping_prev"`
	// Toggling repeat context flag. Optional field.
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	// Toggling shuffle flag. Optional field.
	TogglingShuffle bool `json:"toggling_shuffle"`
	// Toggling repeat track flag. Optional field.
	TogglingRepeatTrack bool `json:"toggling_repeat_track"`
	// Transfering playback between devices. Optional field.
	TransferringPlayback bool `json:"transferring_playback"`
}

type Device struct {
	// The device ID.
	ID string `json:"id"`
	// If this device is the currently active device.
	IsActive bool `json:"is_active"`
	// If this device is currently in a private session.
	IsPrivateSession bool `json:"is_private_session"`
	// Whether controlling this device is restricted. At present if this is "true" then no Web API commands will be accepted by this device.
	IsRestricted bool `json:"is_restricted"`
	// A human-readable name for the device. Some devices have a name that the user can configure (e.g. "Loudest speaker") and some devices have a generic name associated with the manufacturer or device model.
	Name string `json:"name"`
	// Device type, such as "computer", "smartphone" or "speaker".
	Type string `json:"type"`
	// The current volume in percent.
	VolumePercent int `json:"volume_percent"`
}

// A Context Object. Can be null.
type Context struct {
	// The object type, e.g. "artist", "playlist", "album", "show".
	Type string `json:"type"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// External URLs for this context.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// The Spotify URI for the context.
	URI string `json:"uri"`
}

type Track struct {
	// The album on which the track appears. The album object includes a link in href to full information about the album.
	Album Album `json:"album"`
	// The artists who performed the track. Each artist object includes a link in href to more detailed information about the artist.
	Artists []Artists `json:"artists"`
	// A list of the countries in which the track can be played, identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`
	// The disc number (usually 1 unless the album consists of more than one disc).
	DiscNumber int `json:"disc_number"`
	// The track length in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Whether or not the track has explicit lyrics ( true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`
	// Known external IDs for the track.
	ExternalIDs ExternalIDs `json:"external_ids"`
	// Known external URLs for this track.
	ExternalURLs ExternalURLs `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the track.
	Href string `json:"href"`
	// The Spotify ID for the track.
	ID string `json:"id"`
	// Part of the response when Track Relinking is applied. If true, the track is playable in the given market. Otherwise false.
	IsPlayable bool `json:"is_playable"`
	// Part of the response when Track Relinking is applied, and the requested track has been replaced with different track.
	//
	// The track in the linked_from object contains information about the originally requested track.
	LinkedFrom struct{} `json:"linked_from"`
	// Included in the response when a content restriction is applied.
	Restrictions Restrictions `json:"restrictions"`
	// The name of the track.
	Name string `json:"name"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.
	//
	// The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.
	//
	// Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. Note: the popularity value may lag actual popularity by a few days: the value is not updated in real time.
	Popularity int `json:"popularity"`
	// A link to a 30 second preview (MP3 format) of the track. Can be null
	PreviewURL string `json:"preview_url,omitempty"`
	// The number of the track. If an album has several discs, the track number is the number on the specified disc.
	TrackNumber int `json:"track_number"`
	// The object type: "track".
	Type string `json:"type"`
	// The Spotify URI for the track.
	URI string `json:"uri"`
	// Whether or not the track is from a local file.
	IsLocal bool `json:"is_local"`
}

type GetPlaybackStateResponse struct {
	// The device that is currently active.
	Device Device `json:"device"`
	// off, track, context
	RepeatState string `json:"repeat_state"`
	// If shuffle is on or off.
	ShuffleState bool `json:"shuffle_state"`
	// A Context Object. Can be null.
	Context Context `json:"context,omitempty"`
	// Unix Millisecond Timestamp when data was fetched.
	Timestamp int `json:"timestamp"`
	// Progress into the currently playing track or episode. Can be null.
	ProgressMs int `json:"progress_ms,omitempty"`
	// If something is currently playing, return true.
	IsPlaying bool `json:"is_playing"`

	// The currently playing track or episode. Can be null.
	Item Track `json:"item,omitempty"`

	// The object type of the currently playing item. Can be one of track, episode, ad or unknown.
	CurrentlyPlayingType string `json:"currently_playing_type"`
	// Allows to update the user interface based on which playback actions are available within the current context.
	Actions Actions `json:"actions"`
}

// Get information about the user’s current playback state, including track or episode, progress, and active device.
//
// Required scopes: user-read-playback-state
func (c *Client) GetPlaybackState() (GetPlaybackStateResponse, *SpotifyError) {
	var response GetPlaybackStateResponse
	var err SpotifyError
	c.get(fmt.Sprintf("%s/me/player", URL)).Receive(&response, &err)

	return response, &err
}

// Transfer playback to a new device and determine if it should start playing.
//
// Required scope: user-modify-playback-state
func (c *Client) TransferPlayback() {}

// Get information about a user’s available devices.
func (c *Client) GetAvailableDevices() {}

// Get the object currently being played on the user's Spotify account.
//
// Required scope: user-read-currently-playing
func (c *Client) GetCurrentlyPlayingTrack() {}

// Start a new context or resume current playback on the user's active device.
//
// Required scope: user-modify-playback-state
func (c *Client) StartResumePlayback() {}

// Pause playback on the user's account.
//
// Required scope: user-modify-playback-state
func (c *Client) PausePlayback() {}

// Skips to next track in the user’s queue.
//
// Required scope: user-modify-playback-state
func (c *Client) SkipToNext() {}

// Skips to previous track in the user’s queue.
//
// Required scope: user-modify-playback-state
func (c *Client) SkipToPrevious() {}

// Seeks to the given position in the user’s currently playing track.
//
// Required scope: user-modify-playback-state
func (c *Client) SeekToPosition() {}

// Set the repeat mode for the user’s playback. Options are repeat-track, repeat-context, and off.
//
// Required scope: user-modify-playback-state
func (c *Client) SetRepeatMode() {}

// Set the volume for the user’s current playback device.
//
// Required scope: user-modify-playback-state
func (c *Client) SetPlaybackVolume() {}

// Toggle shuffle on or off for user’s playback.
//
// Required scope: user-modify-playback-state
func (c *Client) TogglePlaybackShuffle() {}

// Get tracks from the current user's recently played tracks. Note: Currently doesn't support podcast episodes.
//
// Required scope: user-read-recently-played
func (c *Client) GetRecentlyPlayedTracks() {}

// Get the list of objects that make up the user's queue.
//
// Required scope: user-read-playback-state
func (c *Client) GetTheUsersQueue() {}

// Add an item to the end of the user's current playback queue.
//
// Required scope: user-modify-playback-state
func (c *Client) AddItemToPlaybackQueue() {}
