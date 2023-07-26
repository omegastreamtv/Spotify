package spotify

// External URLs for this context.
type ExternalURLs struct {
	// The Spotify URL for the object.
	Spotify string `json:"spotify"`
}

// The cover art for the episode in various sizes, widest first.
type Image struct {
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

// Information about the currently active track or episode.
type Playback struct {
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
	ProgressMS int `json:"progress_ms,omitempty"`
	// If something is currently playing, return true.
	IsPlaying bool `json:"is_playing"`
	// The currently playing track or episode. Can be null.
	Item interface{} `json:"item,omitempty"`
	// The object type of the currently playing item. Can be one of track, episode, ad or unknown.
	CurrentlyPlayingType string `json:"currently_playing_type"`
	// Allows to update the user interface based on which playback actions are available within the current context.
	Actions Actions `json:"actions"`
}

type GetPlaybackStateParams struct {
	Market Market `url:"market"`
	// A comma-separated list of item types that your client supports besides the default track type. Valid types are: track and episode.
	AdditionalTypes string `url:"additional_types"`
}

type GetPlaybackStateResponse struct {
	Playback
}

// Get information about the user’s current playback state, including track or episode, progress, and active device.
//
// Required scopes: user-read-playback-state
func (c *Client) GetPlaybackState(params *GetPlaybackStateParams) (*GetPlaybackStateResponse, error) {
	state := GetPlaybackStateResponse{}
	var err *SpotifyError

	c.get("/me/player").QueryStruct(params).Receive(&state, &err)

	if err != nil {
		return nil, err
	}

	return &state, nil
}

type TransferPlaybackBody struct {
	// A JSON array containing the ID of the device on which playback should be started/transferred.
	//
	// Note: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return 400 Bad Request
	DeviceIDs []string `json:"device_ids"`
	// true: ensure playback happens on new device.
	//
	// false or not provided: keep the current playback state.
	Play bool `json:"play,omitempty"`
}

// Transfer playback to a new device and determine if it should start playing.
//
// Required scope: user-modify-playback-state
func (c *Client) TransferPlayback(body *TransferPlaybackBody) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player").BodyJSON(body).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type GetAvailableDevicesResponse struct {
	Devices []Device `json:"devices"`
}

// Get information about a user’s available devices.
func (c *Client) GetAvailableDevices() (*GetAvailableDevicesResponse, error) {
	devices := GetAvailableDevicesResponse{}
	var err *SpotifyError

	c.get("/me/player/devices").Receive(&devices, &err)

	if err != nil {
		return nil, err
	}

	return &devices, nil
}

type GetCurrentlyPlayingTrackParams struct {
	Market Market `url:"market"`
	// A comma-separated list of item types that your client supports besides the default track type. Valid types are: track and episode.
	AdditionalTypes string `url:"additional_types"`
}

type GetCurrentlyPlayingTrackResponse struct {
	Playback
}

// Get the object currently being played on the user's Spotify account.
//
// Required scope: user-read-currently-playing
func (c *Client) GetCurrentlyPlayingTrack(params *GetCurrentlyPlayingTrackParams) (*GetCurrentlyPlayingTrackResponse, error) {
	content := GetCurrentlyPlayingTrackResponse{}
	var err *SpotifyError

	c.get("/me/player/currently-playing").QueryStruct(params).Receive(&content, &err)

	if err != nil {
		return nil, err
	}

	return &content, nil
}

type StartResumePlaybackParams struct {
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

type StartResumePlaybackBody struct {
	// Optional. Spotify URI of the context to play. Valid contexts are albums, artists & playlists.
	ContextURI string `json:"context_uri,omitempty"`
	// Optional. A JSON array of the Spotify track URIs to play.
	URLs []string `json:"ur_ls,omitempty"`
	// Optional. Indicates from where in the context playback should start. Only available when context_uri corresponds to an album or playlist object "position" is zero based and can’t be negative.
	Offset     interface{} `json:"offset,omitempty"`
	PositionMS int         `json:"position_ms,omitempty"`
}

// Start a new context or resume current playback on the user's active device.
//
// Required scope: user-modify-playback-state
func (c *Client) StartResumePlayback(params *StartResumePlaybackParams, body *StartResumePlaybackBody) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/play").QueryStruct(params).BodyJSON(body).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type PausePlaybackParams struct {
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Pause playback on the user's account.
//
// Required scope: user-modify-playback-state
func (c *Client) PausePlayback(params *PausePlaybackParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/pause").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type SkipToNextParams struct {
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Skips to next track in the user’s queue.
//
// Required scope: user-modify-playback-state
func (c *Client) SkipToNext(params *SkipToNextParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/next").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type SkipToPreviousParams struct {
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Skips to previous track in the user’s queue.
//
// Required scope: user-modify-playback-state
func (c *Client) SkipToPrevious(params *SkipToPreviousParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/previous").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type SeekToPositionParams struct {
	// The position in milliseconds to seek to. Must be a positive number. Passing in a position that is greater than the length of the track will cause the player to start playing the next song.
	PositionMS int `url:"position_ms"`
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Seeks to the given position in the user’s currently playing track.
//
// Required scope: user-modify-playback-state
func (c *Client) SeekToPosition(params *SeekToPositionParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/seek").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type SetRepeatModeParams struct {
	// track, context or off.
	//
	// "track" will repeat the current track.
	//
	// "context" will repeat the current context.
	//
	// "off" will turn repeat off.
	State string `url:"state"`
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Set the repeat mode for the user’s playback. Options are repeat-track, repeat-context, and off.
//
// Required scope: user-modify-playback-state
func (c *Client) SetRepeatMode(params *SetRepeatModeParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/repeat").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type SetPlaybackVolumeParams struct {
	// The volume to set. Must be a value from 0 to 100 inclusive.
	VolumePercent int `url:"volume_percent"`
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Set the volume for the user’s current playback device.
//
// Required scope: user-modify-playback-state
func (c *Client) SetPlaybackVolume(params *SetPlaybackVolumeParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/volume").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type TogglePlaybackShuffleParams struct {
	// true: Shuffle user's playback.
	//
	// false: Do not shuffle user's playback.
	State bool `url:"state"`
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Toggle shuffle on or off for user’s playback.
//
// Required scope: user-modify-playback-state
func (c *Client) TogglePlaybackShuffle(params *TogglePlaybackShuffleParams) error {
	var res struct{}
	var err *SpotifyError

	c.put("/me/player/shuffle").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type GetRecentlyPlayedTracksParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit,omitempty"`
	// A Unix timestamp in milliseconds. Returns all items after (but not including) this cursor position. If after is specified, before must not be specified.
	After int `url:"after,omitempty"`
	// A Unix timestamp in milliseconds. Returns all items before (but not including) this cursor position. If before is specified, after must not be specified.
	Before int `url:"before,omitempty"`
}

type GetRecentlyPlayedTracksResponse struct {
	Pagination
	Items []struct {
		// The track the user listened to.
		Track Track `json:"track"`
		// The date and time the track was played.
		PlayedAt string `json:"played_at"`
		// The context the track was played from.
		Context Context `json:"context"`
	} `json:"items"`
}

// Get tracks from the current user's recently played tracks. Note: Currently doesn't support podcast episodes.
//
// Required scope: user-read-recently-played
func (c *Client) GetRecentlyPlayedTracks(params *GetRecentlyPlayedTracksParams) (*GetRecentlyPlayedTracksResponse, error) {
	tracks := GetRecentlyPlayedTracksResponse{}
	var err *SpotifyError

	c.get("/me/player/recently-played").QueryStruct(params).Receive(&tracks, &err)

	if err != nil {
		return nil, err
	}

	return &tracks, nil
}

type GetTheUsersQueueResponse struct {
	// The currently playing track or episode. Can be null.
	CurrentlyPlaying interface{} `json:"currently_playing"`
	// The tracks or episodes in the queue. Can be empty.
	Queue []interface{} `json:"queue"`
}

// Get the list of objects that make up the user's queue.
//
// Required scope: user-read-playback-state
func (c *Client) GetTheUsersQueue() (*GetTheUsersQueueResponse, error) {
	queue := GetTheUsersQueueResponse{}
	var err *SpotifyError

	c.get("/me/player/queue").Receive(&queue, &err)

	if err != nil {
		return nil, err
	}

	return &queue, nil
}

type AddItemToPlaybackQueueParams struct {
	// The uri of the item to add to the queue. Must be a track or an episode uri.
	URI string `url:"uri"`
	// The id of the device this command is targeting. If not supplied, the user's currently active device is the target.
	DeviceID string `url:"device_id"`
}

// Add an item to the end of the user's current playback queue.
//
// Required scope: user-modify-playback-state
func (c *Client) AddItemToPlaybackQueue(params *AddItemToPlaybackQueueParams) error {
	var res struct{}
	var err *SpotifyError

	c.post("/me/player/queue").QueryStruct(params).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}
