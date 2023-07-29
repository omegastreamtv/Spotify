package spotify

import "fmt"

type Playlist struct {
	// true if the owner allows other users to modify the playlist.
	Collaborative bool `json:"collaborative"`
	// The playlist description. Only returned for modified, verified playlists, otherwise null.
	Description string `json:"description"`
	// Known external URLs for this playlist.
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	// A link to the Web API endpoint providing full details of the playlist.
	Href string `json:"href"`
	// The Spotify ID for the playlist.
	ID string `json:"id"`
	// Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order.
	//
	// Note: If returned, the source URL for the image (url) is temporary and will expire in less than a day.
	Images []Image `json:"images"`
	// The name of the playlist.
	Name string `json:"name"`
	// The user who owns the playlist
	Owner Owner `json:"owner"`
	// The playlist's public/private status: true the playlist is public, false the playlist is private, null the playlist status is not relevant.
	Public bool `json:"public"`
	// The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version
	SnapshotID string `json:"snapshot_id"`
	// The object type: "playlist"
	Type string `json:"type"`
	// The Spotify URI for the playlist.
	URI string `json:"uri"`
}

type BasicPlaylist struct {
	Playlist
	// A collection containing a link ( href ) to the Web API endpoint where full details of the playlist's tracks can be retrieved, along with the total number of tracks in the playlist. Note, a track object may be null. This can happen if a track is no longer available.
	Tracks struct {
		// A link to the Web API endpoint where full details of the playlist's tracks can be retrieved.
		Href string
		// Number of tracks in the playlist.
		Total int
	}
}

type FullPlaylistState struct {
	Playlist
	// Information about the followers of the playlist.
	Followers Followers `json:"followers"`
	// The tracks of the playlist.
	Tracks struct {
		Pagination `json:"pagination"`
		Items      []PlaylistTrackState `json:"items"`
	} `json:"tracks"`
}

type FullPlaylist struct {
	Playlist
	// Information about the followers of the playlist.
	Followers Followers `json:"followers"`
	// The tracks of the playlist.
	Tracks struct {
		Pagination `json:"pagination"`
		Items      []PlaylistTrack `json:"items"`
	} `json:"tracks"`
}

type PlaylistTrackState struct {
	// The date and time the track or episode was added. Note: some very old playlists may return null in this field.
	AddedAt string `json:"added_at"`
	// The Spotify user who added the track or episode. Note: some very old playlists may return null in this field.
	AddedBy User `json:"added_by"`
	// Whether this track or episode is a local file or not.
	IsLocal bool `json:"is_local"`
	// Information about the track or episode.
	Track interface{} `json:"track"`
}

type PlaylistTrack struct {
	// The date and time the track or episode was added. Note: some very old playlists may return null in this field.
	AddedAt string `json:"added_at"`
	// The Spotify user who added the track or episode. Note: some very old playlists may return null in this field.
	AddedBy User `json:"added_by"`
	// Whether this track or episode is a local file or not.
	IsLocal bool `json:"is_local"`
	// Information about the track or episode.
	Track Either[Track, Episode] `json:"track"`
}

type Owner struct {
	// Known public external URLs for this user.
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	// Information about the followers of this user.
	Followers Followers `json:"followers"`
	// A link to the Web API endpoint for this user.
	Href string `json:"href"`
	// The Spotify user ID for this user.
	ID string `json:"id"`
	// The object type.
	Type string `json:"type"`
	// The Spotify URI for this user.
	URI string `json:"uri"`
	// The name displayed on the user's profile. null if not available.
	DisplayName string `json:"display_name"`
}

type GetPlaylistParams struct {
	Market Market `url:"market"`
	// Filters for the query: a comma-separated list of the fields to return. If omitted, all fields are returned. For example, to get just the playlist''s description and URI: fields=description,uri.
	//
	// A dot separator can be used to specify non-reoccurring fields, while parentheses can be used to specify reoccurring fields within objects. For example, to get just the added date and user ID of the adder: fields=tracks.items(added_at,added_by.id).
	//
	// Use multiple parentheses to drill down into nested objects, for example: fields=tracks.items(track(name,href,album(name,href))).
	//
	// Fields can be excluded by prefixing them with an exclamation mark, for example: fields=tracks.items(track(name,href,album(!name,href)))
	Fields string `url:"fields"`
	// A comma-separated list of item types that your client supports besides the default track type. Valid types are: track and episode.
	AdditionalTypes string `url:"additional_types"`
}

type GetPlaylistStateResponse struct {
	FullPlaylistState
}

type GetPlaylistResponse struct {
	FullPlaylist
}

// Get a playlist owned by a Spotify user.
func (c *Client) GetPlaylist(playlistId string, params *GetPlaylistParams) (*GetPlaylistResponse, error) {
	playlistState := GetPlaylistStateResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/playlists/%s", playlistId)).QueryStruct(params).Receive(&playlistState, &err)

	if err != nil {
		return nil, err
	}

	playlist := GetPlaylistResponse{
		FullPlaylist: FullPlaylist{
			Playlist:  playlistState.Playlist,
			Followers: playlistState.Followers,
			Tracks: struct {
				Pagination "json:\"pagination\""
				Items      []PlaylistTrack "json:\"items\""
			}{
				Pagination: playlistState.Tracks.Pagination,
				Items:      []PlaylistTrack{},
			},
		},
	}

	for _, it := range playlistState.Tracks.Items {
		// Convert Track of type interface{} to Either[Track, Episode]
		eitherTrackOrEpisode, err := EitherTrackOrEpisode(it.Track)
		if err != nil {
			return nil, err
		}
		playlistTrack := PlaylistTrack{
			AddedAt: it.AddedAt,
			AddedBy: it.AddedBy,
			IsLocal: it.IsLocal,
			Track:   eitherTrackOrEpisode,
		}
		playlist.FullPlaylist.Tracks.Items = append(playlist.FullPlaylist.Tracks.Items, playlistTrack)
	}

	return &playlist, nil
}

type ChangePlaylistDetailsBody struct {
	// The new name for the playlist.
	Name string `json:"name,omitempty"`
	// If true the playlist will be public, if false it will be private.
	Public bool `json:"public,omitempty"`
	// If true, the playlist will become collaborative and other users will be able to modify the playlist in their Spotify client.
	//
	// Note: You can only set collaborative to true on non-public playlists.
	Collaborative bool `json:"collaborative,omitempty"`
	// Value for playlist description as displayed in Spotify Clients and in the Web API.
	Description string `json:"description,omitempty"`
}

// Change a playlist's name and public/private state. (The user must, of course, own the playlist.)
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) ChangePlaylistDetails(playlistId string, body *ChangePlaylistDetailsBody) error {
	var res struct{}
	var err *SpotifyError

	c.put(fmt.Sprintf("/playlists/%s", playlistId)).BodyJSON(body).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}

type GetPlaylistItemsParams struct {
	Market Market `url:"market"`
	// Filters for the query: a comma-separated list of the fields to return. If omitted, all fields are returned. For example, to get just the playlist''s description and URI: fields=description,uri.
	//
	// A dot separator can be used to specify non-reoccurring fields, while parentheses can be used to specify reoccurring fields within objects. For example, to get just the added date and user ID of the adder: fields=tracks.items(added_at,added_by.id).
	//
	// Use multiple parentheses to drill down into nested objects, for example: fields=tracks.items(track(name,href,album(name,href))).
	//
	// Fields can be excluded by prefixing them with an exclamation mark, for example: fields=tracks.items(track(name,href,album(!name,href)))
	Fields string `url:"fields"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
	// A comma-separated list of item types that your client supports besides the default track type. Valid types are: track and episode.
	AdditionalTypes string `url:"additional_types"`
}

type GetPlaylistStateItemsResponse struct {
	Pagination `json:"pagination"`
	Items      []PlaylistTrackState `json:"items"`
}

type GetPlaylistItemsResponse struct {
	Pagination `json:"pagination"`
	Items      []PlaylistTrack `json:"items"`
}

// Get full details of the items of a playlist owned by a Spotify user.
//
// Required scope: playlist-read-private
func (c *Client) GetPlaylistItems(playlistId string, params *GetPlaylistItemsParams) (*GetPlaylistItemsResponse, error) {
	itemsState := GetPlaylistStateItemsResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/playlists/%s/tracks", playlistId)).QueryStruct(params).Receive(&itemsState, &err)

	if err != nil {
		return nil, err
	}

	items := GetPlaylistItemsResponse{
		Pagination: itemsState.Pagination,
	}

	for _, it := range itemsState.Items {
		// Convert Track of type interface{} to Either[Track, Episode]
		eitherTrackOrEpisode, err := EitherTrackOrEpisode(it.Track)
		if err != nil {
			return nil, err
		}

		pl := PlaylistTrack{
			AddedAt: it.AddedAt,
			AddedBy: it.AddedBy,
			IsLocal: it.IsLocal,
			Track:   eitherTrackOrEpisode,
		}

		items.Items = append(items.Items, pl)
	}

	return &items, nil
}

type UpdatePlaylistItemsBody struct {
	URIs []string `json:"ur_is,omitempty"`
	// The position of the first item to be reordered.
	RangeStart int `json:"range_start,omitempty"`
	// The position where the items should be inserted.
	//
	// To reorder the items to the end of the playlist, simply set insert_before to the position after the last item.
	InsertBefore int `json:"insert_before,omitempty"`
	// The amount of items to be reordered. Defaults to 1 if not set.
	//
	// The range of items to be reordered begins from the range_start position, and includes the range_length subsequent items.
	RangeLength int `json:"range_length,omitempty"`
	// The playlist's snapshot ID against which you want to make the changes.
	SnapshotID string `json:"snapshot_id,omitempty"`
}

type UpdatePlaylistItemsResponse struct {
	// A snapshot ID for the playlist
	SnapshotID string `json:"snapshot_id"`
}

// Either reorder or replace items in a playlist depending on the request's parameters. To reorder items, include range_start, insert_before, range_length and snapshot_id in the request's body. To replace items, include uris as either a query parameter or in the request's body. Replacing items in a playlist will overwrite its existing items. This operation can be used for replacing or clearing items in a playlist.
//
// Note: Replace and reorder are mutually exclusive operations which share the same endpoint, but have different parameters. These operations can't be applied together in a single request.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) UpdatePlaylistItems(playlistId string, body *UpdatePlaylistItemsBody) (*UpdatePlaylistItemsResponse, error) {
	playlist := UpdatePlaylistItemsResponse{}
	var err *SpotifyError

	c.put(fmt.Sprintf("/playlists/%s/tracks", playlistId)).BodyJSON(body).Receive(&playlist, &err)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

type AddItemsToPlaylistBody struct {
	// A JSON array of the Spotify URIs to add. A maximum of 100 items can be added in one request.
	URIs []string `json:"ur_is,omitempty"`
	// The position to insert the items, a zero-based index. For example, to insert the items in the first position: position=0; to insert the items in the third position: position=2. If omitted, the items will be appended to the playlist. Items are added in the order they are listed in the query string or request body.
	Position int `json:"position"`
}

type AddItemsToPlaylistResponse struct {
	// A snapshot ID for the playlist
	SnapshotID string `json:"snapshot_id"`
}

// Add one or more items to a user's playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) AddItemsToPlaylist(playlistId string, body *AddItemsToPlaylistBody) (*AddItemsToPlaylistResponse, error) {
	playlist := AddItemsToPlaylistResponse{}
	var err *SpotifyError

	c.post(fmt.Sprintf("/playlists/%s/tracks", playlistId)).BodyJSON(body).Receive(&playlist, &err)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

type RemovePlaylistItemsBody struct {
	Tracks []struct {
		URI string `json:"uri"`
	} `json:"tracks"`
	// The playlist's snapshot ID against which you want to make the changes. The API will validate that the specified items exist and in the specified positions and make the changes, even if more recent changes have been made to the playlist.
	SnapshotID string `json:"snapshot_id"`
}

type RemovePlaylistItemsResponse struct {
	// A snapshot ID for the playlist
	SnapshotID string `json:"snapshot_id"`
}

// Remove one or more items from a user's playlist.
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) RemovePlaylistItems(playlistId string, body *RemovePlaylistItemsBody) (*RemovePlaylistItemsResponse, error) {
	playlist := RemovePlaylistItemsResponse{}
	var err *SpotifyError

	c.delete(fmt.Sprintf("/playlists/%s/tracks", playlistId)).BodyJSON(body).Receive(&playlist, &err)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

type GetCurrentUsersPlaylistsParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first playlist to return. Default: 0 (the first object). Maximum offset: 100.000. Use with limit to get the next set of playlists.
	Offset int `url:"offset"`
}

type GetCurrentUsersPlaylistsResponse struct {
	Pagination
	Items []BasicPlaylist `json:"items"`
}

// Get a list of the playlists owned or followed by the current Spotify user.
//
// Required scope: playlist-read-private
func (c *Client) GetCurrentUsersPlaylists(params *GetCurrentUsersPlaylistsParams) (*GetCurrentUsersPlaylistsResponse, error) {
	playlists := GetCurrentUsersPlaylistsResponse{}
	var err *SpotifyError

	c.get("/me/playlists").QueryStruct(params).Receive(&playlists, &err)

	if err != nil {
		return nil, err
	}

	return &playlists, nil
}

type GetUsersPlaylistsParams struct {
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first playlist to return. Default: 0 (the first object). Maximum offset: 100.000. Use with limit to get the next set of playlists.
	Offset int `url:"offset"`
}

type GetUsersPlaylistsResponse struct {
	Pagination
	Items []BasicPlaylist `json:"items"`
}

// Get a list of the playlists owned or followed by a Spotify user.
//
// Required scope: playlist-read-private, playlist-read-collaborative
func (c *Client) GetUsersPlaylists(userId string, params *GetUsersPlaylistsParams) (*GetUsersPlaylistsResponse, error) {
	playlists := GetUsersPlaylistsResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/users/%s/playlists", userId)).QueryStruct(params).Receive(&playlists, &err)

	if err != nil {
		return nil, err
	}

	return &playlists, nil
}

type CreatePlaylistBody struct {
	// The name for the new playlist, for example "Your Coolest Playlist". This name does not need to be unique; a user may have several playlists with the same name.
	Name string `json:"name"`
	// Defaults to true. If true the playlist will be public, if false it will be private. To be able to create private playlists, the user must have granted the playlist-modify-private scope
	Public bool `json:"public,omitempty"`
	// Defaults to false. If true the playlist will be collaborative. Note: to create a collaborative playlist you must also set public to false. To create collaborative playlists you must have granted playlist-modify-private and playlist-modify-public scopes.
	Collaborative bool `json:"collaborative,omitempty"`
	// Playlist description as displayed in Spotify Clients and in the Web API.
	Description string `json:"description,omitempty"`
}

type CreatePlaylistResponse struct {
	FullPlaylist
}

// Create a playlist for a Spotify user. (The playlist will be empty until you add tracks.)
//
// Required scope: playlist-modify-public, playlist-modify-private
func (c *Client) CreatePlaylist(userId string, body *CreatePlaylistBody) (*CreatePlaylistResponse, error) {
	playlist := CreatePlaylistResponse{}
	var err *SpotifyError

	c.post(fmt.Sprintf("/users/%s/playlists", userId)).BodyJSON(body).Receive(&playlist, &err)

	if err != nil {
		return nil, err
	}

	return &playlist, nil
}

type GetFeaturedPlaylistsParams struct {
	// A country: an ISO 3166-1 alpha-2 country code. Provide this parameter if you want the list of returned items to be relevant to a particular country. If omitted, the returned items will be relevant to all countries.
	Country string `url:"country"`
	// The desired language, consisting of a lowercase ISO 639-1 language code and an uppercase ISO 3166-1 alpha-2 country code, joined by an underscore. For example: es_MX, meaning "Spanish (Mexico)". Provide this parameter if you want the results returned in a particular language (where available).
	//
	// Note: if locale is not supplied, or if the specified language is not available, all strings will be returned in the Spotify default language (American English). The locale parameter, combined with the country parameter, may give odd results if not carefully matched. For example country=SE&locale=de_DE will return a list of categories relevant to Sweden but as German language strings.
	Locale string `url:"locale"`
	// A timestamp in ISO 8601 format: yyyy-MM-ddTHH:mm:ss. Use this parameter to specify the user's local time to get results tailored for that specific date and time in the day. If not provided, the response defaults to the current UTC time. Example: "2014-10-23T09:00:00" for a user whose local time is 9AM. If there were no featured playlists (or there is no data) at the specified time, the response will revert to the current UTC time.
	Timestamp string `url:"timestamp"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetFeaturedPlaylistsResponse struct {
	Message   string `json:"message"`
	Playlists struct {
		Pagination
		Items []BasicPlaylist `json:"items"`
	} `json:"playlists"`
}

// GetFeaturedPlaylists
func (c *Client) GetFeaturedPlaylists(params *GetFeaturedPlaylistsParams) (*GetFeaturedPlaylistsResponse, error) {
	playlists := GetFeaturedPlaylistsResponse{}
	var err *SpotifyError

	c.get("/browse/featured-playlists").QueryStruct(params).Receive(&playlists, &err)

	if err != nil {
		return nil, err
	}

	return &playlists, nil
}

type GetCategorysPlaylistsParams struct {
	// A country: an ISO 3166-1 alpha-2 country code. Provide this parameter to ensure that the category exists for a particular country.
	Country string `url:"country"`
	// The maximum number of items to return. Default: 20. Minimum: 1. Maximum: 50.
	Limit int `url:"limit"`
	// The index of the first item to return. Default: 0 (the first item). Use with limit to get the next set of items.
	Offset int `url:"offset"`
}

type GetCategorysPlaylistsResponse struct {
	Message   string `json:"message"`
	Playlists struct {
		Pagination
		Items []BasicPlaylist `json:"items"`
	} `json:"playlists"`
}

// Get a list of Spotify playlists tagged with a particular category.
func (c *Client) GetCategorysPlaylists(catId string, params *GetCategorysPlaylistsParams) (*GetCategorysPlaylistsResponse, error) {
	playlists := GetCategorysPlaylistsResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/browse/categories/%s/playlists", catId)).QueryStruct(params).Receive(&playlists, &err)

	if err != nil {
		return nil, err
	}

	return &playlists, nil
}

// A set of images
type GetPlaylistCoverImageResponse []Image

// Get the current image associated with a specific playlist.
func (c *Client) GetPlaylistCoverImage(playlistId string) (*GetPlaylistCoverImageResponse, error) {
	images := GetPlaylistCoverImageResponse{}
	var err *SpotifyError

	c.get(fmt.Sprintf("/playlists/%s/images", playlistId)).Receive(&images, &err)

	if err != nil {
		return nil, err
	}

	return &images, nil
}

// Base64 encoded JPEG image data, maximum payload size is 256 KB.
type AddCustomPlaylistCoverImageBody []byte

// Replace the image used to represent a specific playlist.
//
// Required scope: ugc-image-upload, playlist-modify-public, playlist-modify-private
func (c *Client) AddCustomPlaylistCoverImage(playlistId string) error {
	var res struct{}
	var err *SpotifyError

	c.put(fmt.Sprintf("/playlists/%s/images", playlistId)).Receive(&res, &err)

	if err != nil {
		return err
	}

	return nil
}
